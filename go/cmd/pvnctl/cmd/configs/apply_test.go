package configs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeKustomizationDir(t *testing.T, path string) {
	err := os.MkdirAll(path, 0755)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(path, "kustomization.yaml"), []byte(`
resources:
- deployment.yaml
`), 0755)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(path, "deployment.yaml"), []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api
`), 0755)
	require.NoError(t, err)
}

func makeTestDirectory(t *testing.T) string {
	base, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	makeKustomizationDir(t, filepath.Join(base, "compiled/production"))
	makeKustomizationDir(t, filepath.Join(base, "compiled/production/components"))
	makeKustomizationDir(t, filepath.Join(base, "compiled/staging"))
	makeKustomizationDir(t, filepath.Join(base, "components/increase_storage"))
	return base
}

func TestPreProcessKustomizeConfigsNoExcludes(t *testing.T) {
	dir := makeTestDirectory(t)
	defer os.RemoveAll(dir)
	ctx := context.Background()
	processed, err := preprocessKustomizeConfigs(ctx, dir, ".", nil)
	defer os.RemoveAll(processed)
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(processed, "compiled/production/compiled.yaml"))
	require.FileExists(t, filepath.Join(processed, "compiled/production/components/compiled.yaml"))
	require.FileExists(t, filepath.Join(processed, "compiled/staging/compiled.yaml"))
	require.FileExists(t, filepath.Join(processed, "components/increase_storage/compiled.yaml"))
}

func TestPreProcessKustomizeConfigsWithExcludes(t *testing.T) {
	dir := makeTestDirectory(t)
	defer os.RemoveAll(dir)
	ctx := context.Background()
	for _, patterns := range [][]string{
		{"components"},
		{"**/components"},
		{"components/"},
		{"components/", "foo"},
	} {
		t.Run(fmt.Sprintf("%+v", patterns), func(t *testing.T) {
			processed, err := preprocessKustomizeConfigs(ctx, dir, ".", patterns)
			defer os.RemoveAll(processed)
			require.NoError(t, err)
			require.FileExists(t, filepath.Join(processed, "compiled/production/compiled.yaml"))
			require.NoFileExists(t, filepath.Join(processed, "compiled/production/components/compiled.yaml"))
			require.FileExists(t, filepath.Join(processed, "compiled/staging/compiled.yaml"))
			require.NoFileExists(t, filepath.Join(processed, "components/increase_storage/compiled.yaml"))
		})
	}
}

func TestIncludeExcludePatterns(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	for _, path := range []string{
		"secrets.yaml",
		"bigtenant/secret.yaml",
		"compiled/production/secret.yaml",
		"compiled/production/values.yaml",
		"compiled/staging/secret.yaml",
		"compiled/staging/values.yaml",
		"values.yaml",
	} {
		err := os.MkdirAll(filepath.Join(dir, filepath.Dir(path)), 0755)
		require.NoError(t, err)

		require.NoError(t, os.WriteFile(filepath.Join(dir, path), []byte{}, 0755))
	}

	for i, tc := range []struct {
		include       []string
		exclude       []string
		filesInOutput []string
	}{
		{nil, nil, []string{
			"secrets.yaml",
			"bigtenant/secret.yaml",
			"compiled/production/secret.yaml",
			"compiled/production/values.yaml",
			"compiled/staging/secret.yaml",
			"compiled/staging/values.yaml",
			"values.yaml",
		}},
		{nil, []string{"*secret*"}, []string{
			"compiled/production/values.yaml",
			"compiled/staging/values.yaml",
			"values.yaml",
		}},
		{[]string{"compiled/production"}, nil, []string{
			"compiled/production/secret.yaml",
			"compiled/production/values.yaml",
		}},
		{[]string{"compiled"}, []string{"secret.yaml"}, []string{
			"compiled/production/values.yaml",
			"compiled/staging/values.yaml",
		}},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			pruned, err := constructParallelDirectoryWithPatterns(dir, tc.include, tc.exclude)
			defer os.RemoveAll(pruned)
			require.NoError(t, err)

			var foundFiles []string
			require.NoError(
				t,
				filepath.WalkDir(pruned, func(path string, d os.DirEntry, err error) error {
					if !d.IsDir() {
						relativePath, err := filepath.Rel(pruned, path)
						require.NoError(t, err)
						foundFiles = append(foundFiles, relativePath)
					}
					return nil
				}))

			require.ElementsMatch(t, tc.filesInOutput, foundFiles)
		})
	}
}
