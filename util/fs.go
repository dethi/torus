package util

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// SplitFilename splits a filename in two parts: the name and the
// file extension.
func SplitFilename(filename string) (name string, ext string) {
	filename = filepath.Base(filename)
	if filename == "." {
		return "", ""
	}

	ext = strings.ToLower(filepath.Ext(filename))
	if len(ext) == len(filename) {
		return filename, ""
	}

	name = filename[:len(filename)-len(ext)]
	return
}

// AddPathPrefix adds the prefix to all pathname.
func AddPathPrefix(prefix string, files ...string) []string {
	for i, _ := range files {
		files[i] = filepath.Join(prefix, files[i])
	}
	return files
}

// CreateTarball creates a new tarball at pathname location with all
// the given filepaths.
func CreateTarball(pathname string, files ...string) error {
	tarball, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer tarball.Close()

	tw := tar.NewWriter(tarball)
	defer tw.Close()

	for _, path := range files {
		err := func() error {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			stat, err := f.Stat()
			if err != nil {
				return err
			}

			hdr := &tar.Header{
				Name:    filepath.Base(path),
				Mode:    0644,
				Size:    stat.Size(),
				ModTime: stat.ModTime(),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return errors.Wrapf(err, "write header for %v", path)
			}

			if _, err := io.Copy(tw, f); err != nil {
				return errors.Wrapf(err, "write file %v", path)
			}
			return nil
		}()
		if err != nil {
			return errors.Wrapf(err, "create tarball %v", pathname)
		}
	}
	return nil
}
