package util

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// SplitFilename splits a filename in two parts: the name and the
// file extension.
func SplitFilename(filename string) (string, string) {
	ext := strings.ToLower(filepath.Ext(filename))
	name := filename[:len(filename)-len(ext)]
	return name, ext
}

const cleanName = `(?i)((\[ *)?[a-z]+.cpasbien.[a-z]+( *\])?)|(web(-?dl)?)|(xvid)`

var regexClean = regexp.MustCompile(cleanName)

func CleanName(filename string) string {
	s, ext := SplitFilename(filename)
	s = regexClean.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ".", "-", -1)
	s = strings.Replace(s, "[", "-", -1)
	s = strings.Replace(s, "]", "-", -1)
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Trim(s, "-")
	s = strings.Title(s)

	var last rune = -1
	s = strings.Map(func(c rune) rune {
		if c == '-' && last == '-' {
			return -1
		}
		last = c
		return c
	}, s)

	return s + ext
}

func AddPathPrefix(prefix string, files ...string) []string {
	for i, _ := range files {
		files[i] = filepath.Join(prefix, files[i])
	}
	return files
}

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
				Name:    CleanName(filepath.Base(path)),
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
