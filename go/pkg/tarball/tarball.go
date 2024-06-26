// adapted from https://github.com/mimoo/eureka/blob/33d5d6e7e90a07ef82773917e543fc10f6f76330/folders.go
package tarball

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// archive a directory, ignores symlinks
func ArchiveTgz(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	err := filepath.Walk(src, func(absFile string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		file, err := filepath.Rel(src, absFile)
		if err != nil {
			return errors.Wrapf(err, "Rel %s %s failed", src, absFile)
		}
		if file == "." {
			return nil
		}
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if regular file, write content
		if fi.Mode().IsRegular() {
			data, err := os.Open(absFile)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}

	return nil
}

// check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}

func UnarchiveTgz(src io.Reader, dst string, prefixFilters ...string) error {
	// ungzip
	zr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	// untar
	tr := tar.NewReader(zr)

	filtersUsed := make([]bool, len(prefixFilters))

	// uncompress each element
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		target := header.Name

		// validate name against path traversal
		if !validRelPath(header.Name) {
			return fmt.Errorf("tar contained invalid name error %q", target)
		}

		if len(prefixFilters) > 0 {
			match := false
			for idx, filter := range prefixFilters {
				if strings.HasPrefix(header.Name, filter) || (strings.HasPrefix(filter, header.Name) && header.Typeflag == tar.TypeDir) {
					match = true
					filtersUsed[idx] = true
					break
				}
			}
			if !match {
				continue
			}
		}

		// add dst + re-format slashes according to system
		target = filepath.Join(dst, header.Name)
		// if no join is needed, replace with ToSlash:
		// target = filepath.ToSlash(header.Name)

		// check the type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it (with 0755 permission)
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		// if it's a file create it (with same permission)
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			// copy over contents
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			fileToWrite.Close()
		}
	}

	for idx, filter := range prefixFilters {
		if !filtersUsed[idx] {
			return errors.Errorf("Did not encounter any files matching filter %s", filter)
		}
	}
	return nil
}
