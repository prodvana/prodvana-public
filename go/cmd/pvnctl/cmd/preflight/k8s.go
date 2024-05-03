package preflight

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil/tui"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	failureStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
)

type checksFinishedMsg struct{}

type setCheckRunningMsg struct {
	index int
}

type checkUpdateMsg struct {
	index  int
	result CheckResult
}

type model struct {
	checklist *tui.Checklist
}

func (m model) Init() tea.Cmd {
	return m.checklist.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case checkUpdateMsg:
		if msg.result.Success {
			m.checklist.SetSuccess(msg.index, msg.result.Message)
		} else if msg.result.Warning {
			m.checklist.SetInconclusive(msg.index, msg.result.Message)
		} else {
			m.checklist.SetFailure(msg.index, msg.result.Message)
		}
	case setCheckRunningMsg:
		m.checklist.SetRunning(msg.index)
	case checksFinishedMsg:
		return m, tea.Quit
	}
	var cmd tea.Cmd
	m.checklist, cmd = m.checklist.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.checklist.View()
}

func getClientsetForContext(contextName string, config clientcmdapi.Config) (*kubernetes.Clientset, error) {
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: contextName}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(config, contextName, configOverrides, nil)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	// pass along the current env to the exec provider if it exists
	// this ensures that any aws env credentials are properly passed
	// to the provider
	if restConfig.ExecProvider != nil && restConfig.ExecProvider.Command == "aws" {
		restConfig.ExecProvider.Env = append(restConfig.ExecProvider.Env,
			clientcmdapi.ExecEnvVar{Name: "AWS_ACCESS_KEY_ID", Value: os.Getenv("AWS_ACCESS_KEY_ID")},
			clientcmdapi.ExecEnvVar{Name: "AWS_SECRET_ACCESS_KEY", Value: os.Getenv("AWS_SECRET_ACCESS_KEY")},
			clientcmdapi.ExecEnvVar{Name: "AWS_SESSION_TOKEN", Value: os.Getenv("AWS_SESSION_TOKEN")},
		)
	}
	return kubernetes.NewForConfig(restConfig)
}

var testName = "pvn-preflight-test"

func testK8sDeployment(ctx context.Context, cs *kubernetes.Clientset) (bool, error) {
	// create namespace
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: testName,
		},
	}
	_, err := cs.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	if err != nil {
		return false, err
	}

	// create service account
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testName,
			Namespace: ns.Name,
		},
	}

	_, err = cs.CoreV1().ServiceAccounts(ns.Name).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return false, err
	}

	replica := int32(1)
	// create deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testName,
			Namespace: ns.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": testName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testName,
					Namespace: ns.Name,
					Labels: map[string]string{
						"app": testName,
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccount.Name,
					Containers: []corev1.Container{
						{
							Name:  testName,
							Image: "us-docker.pkg.dev/pvn-infra/pvn-public/protections:latest",
							Command: []string{
								"sleep", "infinity",
							},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	_, err = cs.AppsV1().Deployments(ns.Name).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return false, err
	}

	timeoutChan := time.After(2 * time.Minute)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	schedulingErrMsg := func() {
		fmt.Println("This can be caused by a few things:")
		fmt.Println("  - The cluster is not configured with enough resources to schedule the test deployment")
		fmt.Println("  - The cluster uses Fargate and there are no Fargate Profiles configured that allow the test deployment to run")
	}

	fmt.Println("Waiting for test deployment to be ready...")
LoopBreak:
	for {
		select {
		case <-timeoutChan:
			labelSelector := metav1.FormatLabelSelector(deployment.Spec.Selector)
			pods, err := cs.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{
				LabelSelector: labelSelector,
			})
			if err == nil && len(pods.Items) > 0 {
				pod := pods.Items[0]
				if pod.Status.Phase == "Running" {
					fmt.Println(successStyle.Render("[ ✓ ] Test deployment was successful"))
					break LoopBreak
				} else {
					fmt.Println(failureStyle.Render("[ x ] The test deployment pod did not become healthy"))
					// Check for  ImagePullBackoff for specific details
					foundImagePullBackoff := false
					for _, containerStatus := range pod.Status.ContainerStatuses {
						if containerStatus.State.Waiting != nil {
							reason := containerStatus.State.Waiting.Reason
							if reason == "ImagePullBackOff" || reason == "ErrImagePull" {
								foundImagePullBackoff = true
							}
						}
					}
					if foundImagePullBackoff {
						fmt.Println("The pod was unable to pull the docker image.")
						fmt.Println("This can indicate an issue with the cluster's networking configuration")
						fmt.Println("Ensure that the Fargate Profile's subnets have a route to a NAT Gateway in a public subnet")
					} else {
						schedulingErrMsg()
						printEvents(cs, ns.Name, deployment.Name)
					}
				}
			} else {
				fmt.Println(failureStyle.Render("[ x ] Scheduling the test deployment timed out"))
				schedulingErrMsg()
				printEvents(cs, ns.Name, deployment.Name)
			}
			break LoopBreak
		case <-ticker.C:
			deployment, err := cs.AppsV1().Deployments(ns.Name).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
			if err != nil {
				fmt.Printf("Error getting deployment: %v\n", err)
				continue
			}
			if deployment.Status.ReadyReplicas == *deployment.Spec.Replicas {
				fmt.Println(successStyle.Render("[ ✓ ] Test deployment was successful"))
				break LoopBreak
			}
		}
	}
	return true, nil
}

func printEvents(clientset *kubernetes.Clientset, namespace, deploymentName string) {
	fieldSelector := clientset.CoreV1().Events(namespace).GetFieldSelector(&deploymentName, &namespace, nil, nil)
	events, err := clientset.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
	if err != nil {
		fmt.Printf("Error getting events for the deployment: %v\n", err)
		return
	}

	for _, event := range events.Items {
		// Filter for events related to scheduling issues
		if event.InvolvedObject.Kind == "Pod" { // && event.Reason == "FailedScheduling" {
			fmt.Printf("Event: %s - %s\n", event.Reason, event.Message)
		}
	}
}

func cleanupTestK8sDeployment(ctx context.Context, cs *kubernetes.Clientset) {
	// check if the namespace exists before trying to delete it
	_, err := cs.CoreV1().Namespaces().Get(ctx, testName, metav1.GetOptions{})
	if err != nil {
		if k8serr.IsNotFound(err) {
			return
		}
		fmt.Printf("Warning: failed to check for existing test namespace: %v\n", err)
		return
	}

	// check if the deployment exists before trying to delete it
	_, err = cs.AppsV1().Deployments(testName).Get(ctx, testName, metav1.GetOptions{})
	if err != nil {
		if k8serr.IsNotFound(err) {
			return
		}
		fmt.Printf("Warning: failed to check for existing test deployment: %v\n", err)
		return
	}

	err = cs.AppsV1().Deployments(testName).Delete(ctx, testName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Warning failed to delete test deployment: %v\n", err)
	}

	err = cs.CoreV1().Namespaces().Delete(ctx, testName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Warning failed to delete test namespace: %v\n", err)
	}
}

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Run preflight checks to validate your Kubernetes cluster is configured for Prodvana.",
	Long: `Run preflight checks to validate your Kubernetes cluster is configured for Prodvana.

pvnctl preflight k8s
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// look at k8s contexts and prompt user to accept checks against one
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, nil)
		config, err := kubeConfig.RawConfig()
		if err != nil {
			cmd.PrintErrf("Error loading kubeconfig: %v\n", err)
			os.Exit(1)
		}
		if len(config.Contexts) == 0 {
			caller := detectAWSCredentials()
			if caller != nil {
				accepted, err := tui.PromptInteractive("No Kubernetes contexts found in kubeconfig. We detected AWS credentials, would you like to use them to list EKS clusters to check? (Y/n)", true)
				if err != nil {
					cmd.PrintErrf("Error getting selection: %v\n", err)
					os.Exit(1)
				}
				if accepted {
					correctCreds, err := tui.PromptInteractive(fmt.Sprintf("We found the following AWS identity: %s. Is this the identity you would like to use to run the verification? (Y/n)", *caller.Arn), true)
					if err != nil {
						cmd.PrintErrf("Error getting selection: %v\n", err)
						os.Exit(1)
					}
					var sess *session.Session
					var profile string
					var region string
					if correctCreds {
						var err error
						region, err = tui.QuestionInteractive("Which AWS region would you like to list?", "us-east-1")
						if err != nil {
							cmd.PrintErrf("Error getting selection: %v\n", err)
							os.Exit(1)
						}
						if region == "" {
							region = "us-east-1"
						}
						sess, err = session.NewSessionWithOptions(session.Options{
							SharedConfigState: session.SharedConfigEnable,
							Config: aws.Config{
								Region: aws.String(region),
							},
						})
						if err != nil {
							cmd.PrintErrf("Error creating AWS session: %v\n", err)
							os.Exit(1)
						}
					} else {
						sess, profile, region, err = InteractiveGetAWSSession("")
						if err != nil {
							cmd.PrintErrf("Error creating AWS session: %v\n", err)
							os.Exit(1)
						}
					}

					eksClient := eks.New(sess)
					listOutput, err := eksClient.ListClusters(&eks.ListClustersInput{})
					if err != nil {
						if awsErr, ok := err.(awserr.Error); ok {
							// check for IAM Errors
							switch awsErr.Code() {
							case "AccessDenied":
								cmd.PrintErrf("Access Denied Error: %s\n", awsErr.Message())
								cmd.PrintErrln("Please ensure the AWS credentials you selected have the required permissions to list EKS clusters.")
							case "UnauthorizedOperation":
								cmd.PrintErrf("Unauthorized Operation Error: %s\n", awsErr.Message())
								cmd.PrintErrln("Please ensure the AWS credentials you selected have the required permissions to list EKS cluster.")
							default:
								cmd.PrintErrf("Error listing EKS clusters: %v\n", err)
							}
						} else {
							cmd.PrintErrf("Error listing EKS clusters: %v\n", err)
						}
						os.Exit(1)
					}

					eksClusters := make([]string, len(listOutput.Clusters))
					for idx, cluster := range listOutput.Clusters {
						eksClusters[idx] = *cluster
					}

					selection, err := tui.GetInteractiveSelection("Select EKS Cluster", eksClusters)
					if err != nil {
						cmd.PrintErrf("Error getting selection: %v\n", err)
						os.Exit(1)
					}
					if selection == "" {
						cmd.PrintErrln("No cluster selected, aborting preflight checks")
						os.Exit(1)
					}

					cmd.Printf("Please run the following command to add the cluster to your kubeconfig, then rerun pvnctl preflight:\n")
					if profile == "" {
						cmd.Printf("aws eks update-kubeconfig --name %s --region %s\n", selection, region)
					} else {
						cmd.Printf("aws eks update-kubeconfig --name %s --region %s --profile %s\n", selection, region, profile)
					}
					os.Exit(1)
				}
			} else {
				// TODO: provid instructions on what to do now
				cmd.PrintErrln("No Kubernetes contexts found in kubeconfig")
				os.Exit(1)
			}
		}

		items := maps.Keys(config.Contexts)
		sort.Strings(items)

		selected := items[0]
		if len(items) == 1 {
			cmd.Printf("Using Kubernetes context: %s\n", selected)
		}
		if len(items) > 1 {
			selection, err := tui.GetInteractiveSelection("Select Kubernetes Context", items)
			if err != nil {
				cmd.PrintErrf("Error getting selection: %v\n", err)
				os.Exit(1)
			}
			if selection == "" {
				cmd.PrintErrln("No context selected, aborting preflight checks")
				os.Exit(1)
			}
			selected = selection
			cmd.Printf("Selected context: %s\n", selected)
		}

		var eksDetails *EKSClusterDetails
		isEKS, clusterName, clusterRegion := DetectEKSCluster(selected)
		if isEKS {
			caller := detectAWSCredentials()
			if caller != nil {
				// ask user if we should fetch additional details about the cluster using their aws credentials
				accepted, err := tui.PromptInteractive("We detected AWS credentials, would you like to use them to fetch additional details about this EKS cluster?\nThis will help further validate the cluster setup. (Y/n)", true)
				if err != nil {
					cmd.PrintErrf("Error getting selection: %v\n", err)
					os.Exit(1)
				}
				if accepted {
					correctCreds, err := tui.PromptInteractive(fmt.Sprintf("We found the following AWS identity: %s. Is this the identity you would like to use to run the verification? (Y/n)", *caller.Arn), true)
					if err != nil {
						cmd.PrintErrf("Error getting selection: %v\n", err)
						os.Exit(1)
					}
					var sess *session.Session
					if correctCreds {
						var err error
						if clusterRegion == "" {
							region, err := tui.QuestionInteractive("Which AWS Region is this EKS cluster in?", "us-east-1")
							if err != nil {
								cmd.PrintErrf("Error getting selection: %v\n", err)
								os.Exit(1)
							}
							if region == "" {
								region = "us-east-1"
							}
							clusterRegion = region
						}
						sess, err = session.NewSessionWithOptions(session.Options{
							SharedConfigState: session.SharedConfigEnable,
							Config: aws.Config{
								Region: aws.String(clusterRegion),
							},
						})
						if err != nil {
							sess = nil
							msg := err.Error()
							if awsErr, ok := err.(awserr.Error); ok {
								msg = awsErr.Message()
							}
							cmd.PrintErrln(warningStyle.Render(fmt.Sprintf("We could not create an AWS session, we'll continue the checks without the additional information: %v", msg)))
						}
					} else {
						sess, _, _, err = InteractiveGetAWSSession(clusterRegion)
						if err != nil {
							sess = nil
							msg := err.Error()
							if awsErr, ok := err.(awserr.Error); ok {
								msg = awsErr.Message()
							}
							cmd.PrintErrln(warningStyle.Render(fmt.Sprintf("We could not create an AWS session, we'll continue the checks without the additional information: %v", msg)))
						}
					}
					if sess != nil {
						eksClient := eks.New(sess)
						eksDetails, err = fetchEKSClusterDetails(eksClient, clusterName)
						if err != nil {
							eksDetails = nil
							msg := err.Error()
							if awsErr, ok := err.(awserr.Error); ok {
								msg = awsErr.Message()
							}
							cmd.PrintErrln(warningStyle.Render(fmt.Sprintf("Fetching cluster details from AWS failed, we'll continue the checks without the additional information: %v", msg)))
						}
					}
				}
			}
		}

		accepted, err := tui.PromptInteractive("Prodvana will now run a series of *read-only* preflight checks on your Kubernetes cluster.\nThis information is used locally and disgarded once validated. Continue? (Y/n)", true)
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if !accepted {
			cmd.PrintErrln("Aborting preflight checks")
			os.Exit(1)
		}

		cs, err := getClientsetForContext(selected, config)
		if err != nil {
			cmd.PrintErrf("failed to connect to the Kubernetes cluster: %v\n", err)
			os.Exit(1)
		}

		checklistItems := make([]tui.ChecklistItem, len(k8sChecks))
		for idx, check := range k8sChecks {
			checklistItems[idx] = tui.ChecklistItem{
				Message: check.RunningText,
				State:   tui.StateWaiting,
			}
		}

		checklist := tui.NewChecklist(checklistItems)
		p := tea.NewProgram(model{
			checklist: checklist,
		})

		go func() {
			ctx := context.Background()
			for idx, check := range k8sChecks {
				p.Send(setCheckRunningMsg{
					index: idx,
				})
				result := check.Check(ctx, ClusterContext{
					Context:    selected,
					Config:     config,
					eksDetails: eksDetails,
				}, cs)
				p.Send(checkUpdateMsg{
					index:  idx,
					result: result,
				})
			}
			p.Send(checksFinishedMsg{})
		}()

		if _, err := p.Run(); err != nil {
			cmd.PrintErrf("Error running preflight checks: %v\n", err)
			os.Exit(1)
		}

		accepted, err = tui.PromptInteractive("To further validate your Kubernetes cluster, we will now create a test deployment.\nThis will be cleaned up after the checks are complete. Continue? (Y/n)", true)
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if !accepted {
			cmd.PrintErrln("Aborting preflight test deployment checks")
			os.Exit(1)
		}

		// run test deployment
		ctx := context.Background()
		cleanupTestK8sDeployment(ctx, cs)
		_, err = testK8sDeployment(ctx, cs)
		if err != nil {
			cmd.PrintErrf("Error running test deployment: %v\n", err)
			cleanupTestK8sDeployment(ctx, cs)
			os.Exit(1)
		}
		cleanupTestK8sDeployment(ctx, cs)
	},
}

func init() {
	RootCmd.AddCommand(k8sCmd)
}
