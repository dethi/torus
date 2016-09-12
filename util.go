package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/zeebo/bencode"
)

func MustTemplate(name string) *template.Template {
	t := template.New(name).Delims("[[", "]]")
	return template.Must(t.Parse(string(MustAsset(name))))
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func InfoHash(data []byte) string {
	var torrent struct {
		Info interface{} `bencode:"info"`
	}

	if err := bencode.DecodeBytes(data, &torrent); err != nil {
		return ""
	}

	binfo, err := bencode.EncodeBytes(torrent.Info)
	if err != nil {
		return ""
	}
	hash := sha1.Sum(binfo)
	return fmt.Sprintf("%x", hash)
}

func CleanName(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	s := filename[:len(filename)-len(ext)]
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
