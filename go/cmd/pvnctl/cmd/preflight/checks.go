package preflight

import (
	"context"
	"fmt"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type ClusterContext struct {
	Context    string
	Config     clientcmdapi.Config
	eksDetails *EKSClusterDetails
}

type CheckResult struct {
	Success bool
	Warning bool
	Message string
	// educational information for the user to resolve the issue
	Followup string
}

type KubernetesCheckFunc func(context.Context, ClusterContext, *kubernetes.Clientset) CheckResult

type KubernetesCheck struct {
	RunningText string
	Check       KubernetesCheckFunc
}

// mostly a sanity check on very basic access permissions
func ListNamespaceCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	_, err := cs.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to list namespaces: %v", err),
		}
	}
	return CheckResult{
		Success: true,
		Message: "Successfully listed namespaces",
	}
}

// check that the cluster has nodes attached to it
func ComputeCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	nodes, err := cs.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to list nodes: %v", err),
		}
	}
	if len(nodes.Items) == 0 && cc.eksDetails == nil {
		return CheckResult{
			Success: false,
			Message: "No nodes found",
		}
	}
	if cc.eksDetails != nil && len(cc.eksDetails.FargateProfiles) > 0 {
		hasNonFargateNodes := cc.eksDetails.NodeGroups != nil && len(cc.eksDetails.NodeGroups) > 0
		if hasNonFargateNodes {
			return CheckResult{
				Success: true,
				Message: "Found Fargate profile(s) and non-Fargate node group(s)",
			}
		}

		// check for a wildcard namespace with no labels required
		// check for an explicit prodvana namespace with no labels required
		// check for a wildcard namespace with labels required
		hasGeneralWildcardProfile := false
		hasProdvanaWildcardProfile := false
		hasGeneralLabelledProfile := false
		for _, profile := range cc.eksDetails.FargateProfiles {
			if profile.PodExecutionRoleArn == nil {
				continue
			}
			if profile.Selectors == nil || len(profile.Selectors) == 0 {
				continue
			}
			for _, selector := range profile.Selectors {
				hasLabels := selector.Labels != nil && len(selector.Labels) > 0
				if selector.Namespace != nil && *selector.Namespace == "*" {
					if hasLabels {
						hasGeneralLabelledProfile = true
					} else {
						hasGeneralWildcardProfile = true
					}
				}
				if selector.Namespace != nil && *selector.Namespace == "prodvana" {
					if hasLabels {
						hasProdvanaWildcardProfile = true
					}
				}
			}
		}
		if hasGeneralWildcardProfile {
			return CheckResult{
				Success: true,
				Message: "Found Fargate profile with wildcard namespace and no labels required",
			}
		} else if hasProdvanaWildcardProfile {
			return CheckResult{
				Success: true,
				Warning: true,
				Message: "Found Fargate profile targeting prodvana namespace with no labels required.",
				// TODO(mike) update this / point to prodvana docs that can better explain what to do about this
				Followup: "The Prodvana Agent will be able to run in this cluster, but it may not be able to deploy pods to other namespaces.",
			}
		} else if hasGeneralLabelledProfile {
			return CheckResult{
				Success: true,
				Warning: true,
				Message: "Found Fargate profile with wildcard namespace and labels required.",
				// TODO (mike) update this / point to prodvana docs that can better explain what to do about this
				Followup: "Deploying the Prodvana Agent to this cluster will require additional configuration -- the agent Deployment will need to be updated to include the required labels in the pod template spec.\n" +
					"Pods deployed by Prodvana must also have these labels in order to be scheduled on Fargate.",
			}
		} else {
			return CheckResult{
				Success: false,
				Message: "No nodes found; this appears to be an EKS cluster with Fargate Profiles enabled, but none of the selectors will match the prodvana namespace used by the agent.",
			}
		}
	}

	if len(nodes.Items) == 0 {
		return CheckResult{
			Success: false,
			Message: "No nodes found",
		}
	}

	return CheckResult{
		Success: true,
		Message: fmt.Sprintf("Found %d node(s)", len(nodes.Items)),
	}
}

func AuthCanCreateAndUpdateNamespaceCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "create",
				Group:    "", // Empty string for core API group
				Resource: "namespaces",
			},
		},
	}
	response, err := cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to create namespaces",
		}
	}
	sar = &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "update",
				Group:    "", // Empty string for core API group
				Resource: "namespaces",
			},
		},
	}
	response, err = cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to update namespaces",
		}
	}

	return CheckResult{
		Success: true,
		Message: "Current user has permission to update namespaces",
	}
}

func AuthCanCreateAndUpdateDeploymentCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "create",
				Group:    "apps",
				Resource: "deployments",
			},
		},
	}
	response, err := cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to create deployments",
		}
	}

	sar = &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "update",
				Group:    "apps",
				Resource: "deployments",
			},
		},
	}
	response, err = cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to update deployments",
		}
	}

	return CheckResult{
		Success: true,
		Message: "Current user has permission to create and update deployments",
	}
}

func AuthCanCreateAndUpdateServiceCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "create",
				Group:    "",
				Resource: "services",
			},
		},
	}
	response, err := cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to create services",
		}
	}

	sar = &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "update",
				Group:    "",
				Resource: "services",
			},
		},
	}
	response, err = cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to update services",
		}
	}

	return CheckResult{
		Success: true,
		Message: "Current user has permission to create and update services",
	}
}

func AuthCanCreateAndUpdateServiceAccountCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "create",
				Group:    "",
				Resource: "serviceaccounts",
			},
		},
	}
	response, err := cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to create serviceaccounts",
		}
	}

	sar = &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "update",
				Group:    "",
				Resource: "serviceaccounts",
			},
		},
	}
	response, err = cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to update serviceaccounts",
		}
	}

	return CheckResult{
		Success: true,
		Message: "Current user has permission to create and update serviceaccounts",
	}
}

func AuthCanGrantClusterRoleBindingCheck(ctx context.Context, cc ClusterContext, cs *kubernetes.Clientset) CheckResult {
	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "create",
				Group:    "rbac.authorization.k8s.io",
				Resource: "clusterrolebindings",
				Name:     "cluster-admin",
			},
		},
	}
	response, err := cs.AuthorizationV1().SelfSubjectAccessReviews().Create(ctx, sar, metav1.CreateOptions{})
	if err != nil {
		return CheckResult{
			Success: false,
			Message: fmt.Sprintf("Failed to create SelfSubjectAccessReview: %v", err),
		}
	}

	if !response.Status.Allowed {
		return CheckResult{
			Success: false,
			Message: "Current user does not have permission to create clusterrolebindings",
		}
	}

	return CheckResult{
		Success: true,
		Message: "Current user has permission to create clusterrolebindings",
	}
}

var k8sChecks = []KubernetesCheck{
	{
		RunningText: "Validating cluster compute...",
		Check:       ComputeCheck,
	},
	{
		RunningText: "Validating access to Kubernetes API by listing namespaces...",
		Check:       ListNamespaceCheck,
	},
	{
		RunningText: "Validating user Namespace authorization...",
		Check:       AuthCanCreateAndUpdateNamespaceCheck,
	},
	{
		RunningText: "Validating user Deployment authorization...",
		Check:       AuthCanCreateAndUpdateDeploymentCheck,
	},
	{
		RunningText: "Validating user Service authorization...",
		Check:       AuthCanCreateAndUpdateServiceCheck,
	},
	{
		RunningText: "Validating user ServiceAccount authorization...",
		Check:       AuthCanCreateAndUpdateServiceAccountCheck,
	},
	{
		RunningText: "Validating user ClusterRoleBinding authorization...",
		Check:       AuthCanGrantClusterRoleBindingCheck,
	},
}
