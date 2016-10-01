package template

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"dashboard.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xacVM\x8f#'\x10\xbdϯ }ٙC\xc3\xced\x15E\t\xb6\xb4\xd1HI.9dVɹ\fe\x83\x97\x86\x0eT{ֲ\xfc\xdf#\xfa\xc3v\xb7\xed\xf5\xac\xb4\xbe\xb8\xa1\x1e\xf0\xea\xd5\a\xc8\x1ftP\xb4\xad\x91\x19\xaa\xdc\xfcNv\u007f\x8cI\x83\xa0\xf3\ac\x92,9\x9c\u007f\n\xb1IRt\x83\xbb\xce\xe2\xac\xff\xcc\"\xbaY\x91h\xeb0\x19D*\x98\x89\xb8\x9c\x15\x86\xa8N\xbf\b\xa1\xb4_'\xae\\h\xf4\xd2AD\xaeB%`\r_\x84\xb3\x8b$\x16\x8d\xab@\xbc\xe7\x8f\xfcI\xa8ԏye=W)\x15L̿\xdbI\xcb\u0a44WL\xa1B\xf1\x81\xff\xc4\u007fl\x0f<\x9d\x9e\x9c+Š\x82\\\x04\xbd\xed\xa9$Td\x83g\xcaAJ\xb3\xc2`\f̦\xb2\x8e\xb6\x82\xb8͟\x8b\xe0tѡ\x19\x93\xdanN\xb1e\xde\xea`\x1d\xdbU\xf0\x04\xd6c<\xb1\xe7X<\x0e\x80V\xfdb\x88\x85y\x1cÞ\x06Xj\x16=\xf2\xc5V\xb5C\x96\fD\x84\x85CF!F\xf4\xc4\x12\xc6\rF)\xcc\xd3\t\x15\xa1\xed\xe6\xc0\xfb8\x90\xa2\xf7y\x88\xfbD\x82~x\xd1e\x15\\S\xf9t\xd5\xe1l͚\x19p˱\xd7\x1e\x0e0\x87\x1bt#\xebx\x9b\xd6^Z\u008a\x19H%\xe1\x17*\x15z\u0088z\xb2\x8c1Y\x1f\xc3\x01\xda\xfaU\xabg\x16%IQ_G\xf7\x8a\xeev\x0e=\xe3\u007f\xa3\nQ\xa7\xfd\xfel\xcdH\xc5\xef\xcf\xf5٦\xcf\xec\xa5\x06\x85oc\xcb?n\xc0\xba\x1c\xfb\xbc\xb2]\xb8߳\xdf\u007f\xbb\xcd[\n\x0f\x9b\xb3\xec\xf8\xe68B$\xab\x1c\x0e\xd0\nS\x82\x15f\xec+D\xdf\xfatU\xac\x1e\\f\xefsU\xfcۭ\xb8!\xf1\xb0jRi\xdd\xef\x939\x14Ab\x10\x91itH\xa8\x19,\t#\xfb\xf033\xa1\x89\x89\xe9&\xe3\xfa:a\x95]E\xc8I\xceoiֻ\xfb\xf5\xaa\xba{[\xedKjkv\bi;\xb0\xa9L\x14m=\xc9\x16Iǎ}\x9c\x8bg\tBf\xfe\x17T(\x05\x99K\xb6\x17\xeb\xd5U\xe3sx\xf5.\x80>\xb7K1>*#&t$\x1d{\xe8\xf0\xdb\xed\"\xf8\x15\x9eT\xd3m\xfe:\xe7tv!\x97\x1e\xe9\v\x80c\tT\b\xabP0\r\x04e?\x9a\x15\xbb\x1d\u007f\x06\xc2\xfd\xbe\x98\xdf\xd8\xc0\xa6Ҫ\x93\xbev\x9a\xd4\xfd\xdd#\xf2\xdeb\xb7\xe3\u007f\xfae\xf8\x03\x92\xd9\xef9A\xbc\xb0\x841i\x87\x8d\x97\xc0\x96P\xae\\X`&a/\x1c \xe0\x8c\xd8\x19٩\xeaYP\xf4z\xa4\xa2\x14\x13ݥh\xf3\xe8-m~\x19B.\x89\x81t;\xba\xd2\xe4/\xe4\xeeĞ/\x9d\xaf7<Y\xcf%V9\xbc\xff`L6\xf8\x1ca\xac\xe6\x93N%\xa7}\v\x0e\xf1\xca\xc1\x9a<\nV\x96L\xb3h\x9f\x02\x1a\xc9XA\xf9\xee<o\xb4\xd3ش\xcb.\x04g\x12\x98\x11\xb7\xeb\xf7g\xa7\xdd\xe1\xfaT\xd1\xd6\xc4\xf2\xbbkVd=\xc4\x1a6\xd0\xcd\x16,E\xf5\xe67M\x9f\xd4|\x9d\xc4#\u007f\xe4\xef\x0f\x13\xf9!\xb3N\x99\u007f\xb7\xed\xf8\xe8\x81c\x8f\xbe\u007f\xe0\x11\xbd\xc6x\xaf\x83j*\xf4\xc4\xffk0n_С\xa2\x10?:w\xff\x8e\xf7\xe0w\x0f\x0f\xbf\x0e\xe9r\xdcZ\x8a.ˤ\xe8^\x90\xff\a\x00\x00\xff\xff\x944\\\xfaY\n\x00\x00",
		hash:  "38b6160673f75cc19f725cf6fe25cd5e3a2bece0c6b2864fc552ea99f40c4c3d",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1475285822, 0),
		size:  2649,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	}
	return gzip.NewReader(strings.NewReader(f.data))
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}
