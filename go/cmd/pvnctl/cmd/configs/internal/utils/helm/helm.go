package helm

import (
	"context"
	go_errors "errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	helm = "helm"
)

func helmCmd(ctx context.Context, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, helm, args...)
}

func Package(ctx context.Context, dir string) (string, func(), error) {
	tmpDir, err := os.MkdirTemp("", "pvnctl-helm-package")
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to create temp dir")
	}
	cmd := helmCmd(ctx, "package", "--destination", tmpDir, dir)
	_, err = cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if go_errors.As(err, &exitErr) {
			return "", nil, errors.Wrapf(err, "failed to run helm package on %s, stderr:\n%s", dir, string(exitErr.Stderr))
		}
		return "", nil, errors.Wrapf(err, "failed to run helm package on %s", dir)
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to read temp dir")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".tgz") {
			return filepath.Join(tmpDir, name), func() {
				os.RemoveAll(tmpDir)
			}, nil
		}
	}
	return "", nil, errors.Errorf("failed to find .tgz file in helm package output")
}
