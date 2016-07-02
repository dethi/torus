// Code generated by go-bindata.
// sources:
// tmpl/list.html.tmpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tmplListHtmlTmpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xdf\x6b\xdb\x30\x10\x7e\x6e\xff\x8a\x5b\xf6\x90\x16\x5a\x6b\x83\x0d\x46\xe6\x04\x1a\x0a\x5d\x18\x94\xb2\x0c\xfa\x50\xc2\x50\xac\x4b\xac\xd6\x91\x8c\x74\x49\x13\x82\xff\xf7\x9d\x2d\x37\xb1\x4b\x7f\xb1\xf5\xc9\xd2\xdd\xf7\xdd\x9d\xbe\xd3\xc9\xf1\x07\x65\x13\xda\xe4\x08\x29\x2d\xb2\xc1\x61\x1c\x3e\x00\x71\x8a\x52\x95\x0b\x5e\x66\xda\xdc\x81\xc3\xac\xdf\xf1\xb4\xc9\xd0\xa7\x88\xd4\x81\x92\xd5\xef\x10\xae\x49\x24\xde\x77\x20\x75\x38\xeb\x77\x84\x48\x94\x89\x6e\xbd\xc2\x4c\xaf\x5c\x64\x90\x84\x36\x9e\xa4\x21\x8f\xd2\x25\x29\xbb\xc4\xe7\x47\xa6\x85\x36\x51\x15\x42\xd4\x09\x49\x53\x86\x83\xdf\xd6\x39\x34\x04\x63\x74\x2b\x9d\x60\x2c\x82\xb9\x2c\x4e\x3c\x54\x17\x4f\xad\xda\xd4\xac\x7c\x10\x7b\x72\xd6\xcc\x07\xd7\x67\xbf\x2e\x47\x97\x17\xbd\x58\xd4\x06\xa0\x14\x81\x42\x3c\x0f\xd2\x21\x70\x7d\x48\xa8\x40\xce\x08\x1d\x7c\xf9\x06\xa9\x5d\x3a\x0f\x6a\x59\xe2\xc0\x73\x4a\x36\x2f\xf4\xdc\x49\xd2\xd6\xc4\x22\xdf\x25\x39\x5b\x49\x9d\xc9\x69\xc6\x31\xb4\xbf\x03\x9f\xcb\x04\x7b\x70\x73\x13\xed\x1c\xe7\x6c\x1f\x97\xe6\xc9\x04\x2e\x86\x15\x37\x90\xb5\xc9\x97\xd4\x10\xae\x03\x5a\xb1\xa8\x95\x0a\xa7\x53\xbb\xde\x2b\xa0\xf4\x2a\xf8\xac\xa3\xd3\xe9\xe6\x34\xb1\x86\xa4\x36\xe8\x3a\x83\x58\xb0\xf3\x31\x8c\x24\xf9\xd7\x40\xa9\x7e\x1d\x93\xcb\xb9\x36\xd5\x91\x9f\x40\x06\xa8\x4f\x9c\xce\x09\xbc\x4b\xfe\xbd\xdb\xb7\xbe\x0c\x1a\x22\x0d\x9a\x61\xc3\x06\x60\x25\x1d\x04\x3c\xf4\xa1\xc5\x3f\xda\xd6\x10\x00\x99\xe7\x23\xd5\x83\x6e\xa9\x7d\xb9\x9c\x4c\xba\x27\x0d\xa7\xfe\x89\x9b\xe0\x1d\x57\x4c\xde\xb6\x10\xda\x28\x5c\x5f\xca\x05\x06\xd0\xe8\x61\xcb\xa0\x1a\x53\x1c\x7f\x3f\xac\x97\x75\xf1\x52\xa9\x6b\xad\xe6\x48\x47\x8d\x30\xcd\xd3\xdd\x57\x5e\x1f\x85\xed\xd0\xae\x1b\xf5\x02\xec\x34\xe5\x94\x1f\xf7\x7d\x6f\x14\x05\x90\x67\x7c\x73\x52\x9b\xa9\x0a\x15\x4a\x8f\xa2\x88\x31\x07\x07\x7c\xac\x25\xd9\x99\x4d\x96\xbe\x07\xe4\x96\x18\x8c\xb9\xbd\x47\x87\x6a\xb8\x79\x30\xee\xc2\x15\xc7\xf5\xf2\x3f\x4e\xc2\x57\x70\xb8\x19\xf3\xbc\x24\x3c\x41\x2f\x1c\xe7\xf1\x55\x6d\x9d\x8a\xc5\xe6\x19\xe6\xaa\x6f\x1a\x46\x80\xad\x79\x52\xff\x13\xe0\x41\xc2\x8c\x1d\xe7\xe8\x13\x64\xae\x99\x77\x8b\x93\x37\x50\xff\x48\x9f\x34\xe8\x67\x7b\x76\x83\x3c\x79\x57\x7d\xca\xd9\x7b\x41\x96\xf6\x68\x76\xdf\x33\x75\x39\xd1\xcf\x67\x6e\xcf\x7b\xab\x1b\x84\x0b\xbe\x66\x54\xf6\x63\xdb\x12\x55\xb3\x87\xb9\xb1\xac\x9f\xf3\x94\x28\xf7\x3d\x21\xc8\x45\x0a\x29\xd5\xd1\xcc\x09\x25\x49\x8a\xed\x76\x64\x66\xf6\x87\xf4\x69\x51\x44\x24\xf9\x91\xd8\x6e\xcb\x0e\x14\x45\x2c\xe4\xa0\xdb\x08\x5a\x84\x3b\x5a\x56\x73\x85\xee\x4a\xce\xb9\x69\x5f\x3f\xbd\xa7\x0c\xfb\x47\xeb\x79\x31\x9e\x7a\xd8\xde\xd2\x0b\xce\xe7\xe8\x88\xcd\xa5\xad\xf9\x64\xc5\x22\xfc\x7a\xf8\x5f\x54\xfd\x32\xff\x06\x00\x00\xff\xff\xc8\x5a\xd9\xb3\x4a\x07\x00\x00"

func tmplListHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplListHtmlTmpl,
		"tmpl/list.html.tmpl",
	)
}

func tmplListHtmlTmpl() (*asset, error) {
	bytes, err := tmplListHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/list.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tmpl/list.html.tmpl": tmplListHtmlTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"tmpl": &bintree{nil, map[string]*bintree{
		"list.html.tmpl": &bintree{tmplListHtmlTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

