package tarball

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeFile(t *testing.T, path string, content []byte, perm os.FileMode) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	require.NoError(t, err)
	err = os.WriteFile(path, content, perm)
	require.NoError(t, err)
}

func makeTestDirectory(t *testing.T) string {
	base, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	makeFile(t, filepath.Join(base, "foo/otherdir/baz"), []byte("hi"), 0644)
	makeFile(t, filepath.Join(base, "foo/otherdir/baz2"), []byte("hi3"), 0644)
	makeFile(t, filepath.Join(base, "foo/bar/baz"), []byte("hi"), 0644)
	makeFile(t, filepath.Join(base, "foo/bar/restricted"), []byte("hi2"), 0600)
	makeFile(t, filepath.Join(base, "foo/bar/baz2"), []byte("hi3"), 0644)
	makeFile(t, filepath.Join(base, "foo/bar1"), []byte("bar1"), 0644)
	makeFile(t, filepath.Join(base, "foo/zbar1"), []byte("zbar1"), 0644)
	makeFile(t, filepath.Join(base, "empty"), nil, 0400)
	require.NoError(t, os.Symlink("does-not-matter", filepath.Join(base, "link")))
	return base
}

func validateTestDirectory(t *testing.T, base string, source bool) {
	count := 0
	err := filepath.Walk(base, func(path string, info fs.FileInfo, err error) error {
		count += 1
		require.NoError(t, err)
		path, err = filepath.Rel(base, path)
		require.NoError(t, err)
		switch path {
		case ".", "foo", "foo/bar", "foo/otherdir":
			require.True(t, info.IsDir())
		case "foo/bar1":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0644), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Equal(t, []byte("bar1"), content)
		case "foo/zbar1":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0644), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Equal(t, []byte("zbar1"), content)
		case "foo/bar/baz", "foo/otherdir/baz":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0644), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Equal(t, []byte("hi"), content)
		case "foo/bar/baz2", "foo/otherdir/baz2":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0644), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Equal(t, []byte("hi3"), content)
		case "foo/bar/restricted":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0600), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Equal(t, []byte("hi2"), content)
		case "empty":
			require.False(t, info.IsDir())
			require.Equal(t, os.FileMode(0400), info.Mode().Perm())
			content, err := os.ReadFile(filepath.Join(base, path))
			require.NoError(t, err)
			require.Empty(t, content)
		case "link":
			require.True(t, source, "should not try to tarball symlink")
		default:
			require.Failf(t, "unrecognized file", "unrecognized file %s", path)
		}
		return nil
	})
	require.NoError(t, err)
	if source {
		require.Equal(t, 13, count) // source has symlink
	} else {
		require.Equal(t, 12, count)
	}
}

func TestArchiveDirectory(t *testing.T) {
	dir := makeTestDirectory(t)
	validateTestDirectory(t, dir, true)
	var buffer bytes.Buffer
	require.NoError(t, ArchiveTgz(dir, &buffer))
	extractBase, err := os.MkdirTemp("", "")
	defer os.RemoveAll(extractBase)
	require.NoError(t, err)
	require.NoError(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase))
	validateTestDirectory(t, extractBase, false)
}

func TestTarballStableSameSource(t *testing.T) {
	dir := makeTestDirectory(t)
	var buffer, buffer2 bytes.Buffer
	require.NoError(t, ArchiveTgz(dir, &buffer))
	require.NoError(t, ArchiveTgz(dir, &buffer2))
	require.Equal(t, buffer2.Bytes(), buffer.Bytes())
}

func TestTarballStableDifferentSource(t *testing.T) {
	dir := makeTestDirectory(t)
	dir2 := makeTestDirectory(t)
	var buffer, buffer2 bytes.Buffer
	require.NoError(t, ArchiveTgz(dir, &buffer))
	require.NoError(t, ArchiveTgz(dir2, &buffer2))
	require.Equal(t, buffer2.Bytes(), buffer.Bytes())
}

func TestFilters(t *testing.T) {
	dir := makeTestDirectory(t)
	validateTestDirectory(t, dir, true)
	var buffer bytes.Buffer
	require.NoError(t, ArchiveTgz(dir, &buffer))
	extractBase, err := os.MkdirTemp("", "")
	defer os.RemoveAll(extractBase)
	require.NoError(t, err)
	require.NoError(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase, "foo/bar"))
	require.FileExists(t, filepath.Join(extractBase, "foo/bar/baz"))
	require.FileExists(t, filepath.Join(extractBase, "foo/bar/restricted"))
	require.NoFileExists(t, filepath.Join(extractBase, "empty"))

	extractBase, err = os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(extractBase)
	require.NoError(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase, "foo/bar/baz", "empty"))
	require.FileExists(t, filepath.Join(extractBase, "foo/bar/baz"))
	require.NoFileExists(t, filepath.Join(extractBase, "foo/bar/restricted"))
	require.FileExists(t, filepath.Join(extractBase, "empty"))

	extractBase, err = os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(extractBase)
	require.Error(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase, "doesnotexist"))

	// validate trailing slashes is ok if it is a directory
	extractBase, err = os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(extractBase)
	require.NoError(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase, "foo/bar/"))
	require.FileExists(t, filepath.Join(extractBase, "foo/bar/baz"))
	require.FileExists(t, filepath.Join(extractBase, "foo/bar/restricted"))

	// validate trailing slashes does not work on files
	extractBase, err = os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(extractBase)
	require.NoError(t, UnarchiveTgz(bytes.NewReader(buffer.Bytes()), extractBase, "foo/bar/baz/"))
	require.NoFileExists(t, filepath.Join(extractBase, "foo/bar/baz"))
	require.NoFileExists(t, filepath.Join(extractBase, "foo/bar/restricted"))
}
