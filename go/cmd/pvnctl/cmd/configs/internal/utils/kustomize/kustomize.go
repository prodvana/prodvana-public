package kustomize

import (
	"context"
	go_errors "errors"
	"os/exec"

	"github.com/pkg/errors"
)

const (
	kubectl   = "kubectl"
	kustomize = "kustomize"
)

func kubectlCmd(ctx context.Context, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, kubectl, args...)
}

func GetRawConfig(ctx context.Context, dir string) ([]byte, error) {
	cmd := kubectlCmd(ctx, kustomize, dir)
	output, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if go_errors.As(err, &exitErr) {
			return nil, errors.Wrapf(err, "failed to run kustomize on %s, stderr:\n%s", dir, string(exitErr.Stderr))
		}
		return nil, errors.Wrapf(err, "failed to run kustomize on %s", dir)
	}
	return output, nil
}
