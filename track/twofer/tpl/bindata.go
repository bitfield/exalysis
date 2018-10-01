// Code generated by go-bindata. DO NOT EDIT.
// sources:
// comment-format.md (1.155kB)
// missing-comment.md (101B)
// plus-used.md (246B)
// stub.md (362B)
// wrong-comment-format.md (108B)

package tpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _commentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\xaa\xec\x36\x0c\x5d\xd7\x5f\x21\xe8\xa2\x6f\x20\x89\x6f\x5b\xba\xe8\xdb\x3e\xda\x4b\x17\x85\x47\xb9\xbb\xa1\xdc\x71\x6c\x25\x16\xe3\x58\xc1\x52\x26\x9d\x96\xfe\x7b\x71\x26\x33\xbc\xa1\x50\xe8\x26\x10\x4b\x96\xce\x39\x3a\x72\xdb\xb6\xe6\x95\x21\x3a\x81\xe3\x58\xd0\x29\x8c\x0b\x05\x4c\x94\x51\x7e\xff\x10\x55\x67\xf9\x68\xed\xc8\xc9\xe5\xb1\xe3\x32\xda\xc0\xde\xe2\x30\xa0\x57\xba\xe0\xfb\xc8\x5d\xd4\x29\x1d\x8c\xeb\x79\x51\x88\xbc\x82\x32\xac\x85\x14\xc1\xf3\x34\x61\x56\xe9\x8c\x79\x8b\x08\xb3\xf3\x67\x37\x3e\x8e\x41\x22\x2f\x29\x80\xa8\x2b\x0a\x2b\x69\x04\x8d\x08\x2b\x97\x00\xef\x9f\x6f\xb9\xef\x30\x70\x4a\xbc\x62\x80\xfe\xba\x85\xef\x45\xb2\x9b\xb0\x33\xe6\x74\x3a\x19\x6b\xe1\xf3\xa3\x34\x9f\x61\x2e\x7c\xa1\x80\x02\xd1\xe5\x70\x05\xcf\xf9\x82\x45\x88\x33\x4c\xa8\x91\x83\xc0\xc0\x05\x96\x4c\x2a\xa0\xd7\x99\xbc\x4b\xe9\x0a\x8b\x60\x00\xca\x50\xd0\xd3\x8c\xd2\x99\xf9\x8b\x9a\x5b\x1f\xf3\x33\x17\xc0\x3f\x66\x2e\x8a\x01\x86\x25\x7b\x25\xce\xd2\xdc\xcb\x36\xb5\x95\xa8\xcb\x2a\x0d\xb8\x1c\x1e\x58\x2f\xae\x90\xeb\x13\x4a\xb3\x51\xf8\x6f\xfe\xcf\xc4\xde\x7a\x79\xe3\x5f\xd3\x4e\xa2\x02\xde\x0a\xcd\xcc\x59\xaa\xd0\x13\xa5\x44\x89\x14\x8b\x74\xa6\x42\xba\xdf\xf8\xa0\xbd\x00\x65\x3d\xd4\x0f\xfc\x65\xbe\xb2\x16\xba\xae\x33\x7f\xdf\xa8\xd4\x79\x1c\x3f\x71\xc0\xdf\xf0\x42\xb8\x7e\xda\x27\xf5\xc5\xc0\x49\xe3\xd2\x77\x9e\xa7\x7d\xf6\x76\x64\xbb\xd2\x99\xec\xbf\x6f\x7d\xbd\x33\x6a\x05\xb3\x62\xf6\x28\x07\xa8\xa9\x30\x2c\x45\x23\x16\x90\x19\x3d\x0d\x84\x02\x1a\x9d\x42\x60\xff\xf0\xc6\x5d\x85\x1e\x61\x58\x52\x82\x47\x89\x06\x30\x07\xca\xe3\x4d\x19\x07\x33\x16\xe2\xd0\x19\xf3\x4b\x06\x17\x02\x55\xed\xab\x02\x3c\x2b\x4d\xf4\x67\xcd\xac\x63\x2d\xe8\x82\xeb\xab\x24\xd7\x66\x37\x4f\x0d\x55\x65\x0b\xde\x9a\x06\xb7\xcd\x0d\x68\xaa\x4e\xa9\x3e\xe1\x15\xae\xbc\x14\xf0\x1c\xd0\x24\xe6\x73\x95\x0e\x8e\xaf\xfc\x8d\x54\xb0\x4b\x45\xea\xf6\x86\x9c\x76\x95\x3e\x5a\xbb\x46\x9a\x67\x2c\xa2\xce\x9f\xb1\x6c\x62\x7d\xf7\xf2\xed\x0f\xf6\xe5\x47\xfb\xfd\x8b\x1d\xb9\x7d\xba\xdc\x8e\x1c\xd8\x3f\x7d\xb9\x8c\xad\xcb\xa1\xbd\xa5\xda\x83\x91\xc5\x47\xd8\x56\xb1\xc6\xeb\xbe\x3d\x9a\x3d\x4e\x0e\x0d\x9c\xb6\x9f\xd3\xcd\x66\xa7\x91\x2b\xca\x53\x5d\x33\x77\x46\x70\x30\xf1\x66\x30\xe5\x4d\x8e\x8d\xbc\xe0\xe6\x56\xe0\xfc\x2c\x3f\xe5\x2d\xcc\xc3\x40\x9e\x5c\x02\xd1\x6b\xc2\xdb\x13\xd0\x98\xe3\x4f\xf7\x3d\x87\x57\xfe\x1f\x8f\x41\xed\x2c\x88\xbb\xea\x4e\x38\x53\x1e\x4d\x8f\x91\xf2\x86\x46\x10\x7c\x64\xf2\x75\xcd\xda\xb6\xfd\x27\x00\x00\xff\xff\xca\xe2\xd9\xbf\x83\x04\x00\x00")

func commentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_commentFormatMd,
		"comment-format.md",
	)
}

func commentFormatMd() (*asset, error) {
	bytes, err := commentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "comment-format.md", size: 1155, mode: os.FileMode(420), modTime: time.Unix(1538426888, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbb, 0x6f, 0x43, 0xf7, 0x78, 0x9a, 0xe0, 0xa3, 0xf0, 0xb3, 0xae, 0x47, 0x1, 0x4c, 0x3f, 0x80, 0xd2, 0x90, 0xf1, 0x91, 0x81, 0x59, 0x27, 0xdb, 0x77, 0x7a, 0x48, 0x26, 0xef, 0x61, 0x4b, 0x27}}
	return a, nil
}

var _missingCommentMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\xc1\x0d\xc2\x30\x0c\x86\xd1\x55\xbe\x0b\x47\x3a\x06\x7b\x98\xd4\x22\x11\x89\x7f\xa9\x76\x41\x6c\xcf\xa9\xf7\xf7\xee\xdc\x92\x11\xbc\x44\x76\x9d\x73\xa7\xdb\xc7\x31\x9a\xd6\xf2\xa8\x8d\x87\x0e\x96\x0e\x67\xf7\xb2\x31\x93\x9f\x4e\x9a\xc5\x05\xa7\xf4\xc6\x8a\xea\x7e\xa5\x24\xbd\xd5\x50\xf0\xf4\xa9\xef\xf6\x0f\x00\x00\xff\xff\x04\x3d\xae\x58\x65\x00\x00\x00")

func missingCommentMdBytes() ([]byte, error) {
	return bindataRead(
		_missingCommentMd,
		"missing-comment.md",
	)
}

func missingCommentMd() (*asset, error) {
	bytes, err := missingCommentMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "missing-comment.md", size: 101, mode: os.FileMode(420), modTime: time.Unix(1538428246, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x1e, 0x1f, 0x6f, 0xb5, 0xfc, 0x5c, 0x55, 0x7c, 0x5f, 0x99, 0x68, 0x72, 0xf8, 0x35, 0x9b, 0x61, 0x26, 0x20, 0x27, 0xdf, 0x91, 0x9, 0x7e, 0x9b, 0x17, 0x87, 0x7, 0x5c, 0x97, 0x8, 0x83}}
	return a, nil
}

var _plusUsedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcf\x4d\x6e\xc4\x30\x08\xc5\xf1\xfd\x9c\xe2\x75\x5d\x35\xf7\xc8\xbe\x07\x08\x65\x48\x8c\xc6\x86\xc8\x10\x55\xbe\x7d\xe5\x7e\xa8\xb3\xfb\x6f\xde\x0f\xf1\x86\x15\x21\x82\xe1\x17\xa8\x0b\xae\x50\x3b\xb0\xbd\x6e\x48\x07\xbb\x31\xa5\x18\xa5\x20\x8b\x20\xb2\xab\x1d\x0b\xde\x0b\x25\x34\x70\x4a\xdf\x85\xb3\x0e\x54\x39\x34\x5f\xb0\x1a\x98\xe2\x5f\x33\x4f\xd0\xe7\x2c\xdf\x6f\x53\xd8\xf6\x96\x1b\x4e\xe2\x07\x1d\x02\x2e\xc2\x0f\x68\xc2\xaf\x5c\xb0\x7e\x9b\xcd\xbb\x80\xbd\x35\xb7\x3a\x70\x85\xdc\xb1\x7b\xff\x3d\x3d\xb3\x51\xe6\x4c\xb2\x3b\x0a\xfd\x0d\xe8\xa4\x0f\xad\x9a\x2a\x71\xcb\x42\x86\xd0\x76\xd6\xf1\xf4\xc2\xdc\xfc\x28\xb1\x7c\x05\x00\x00\xff\xff\x80\x22\x7a\x88\xf6\x00\x00\x00")

func plusUsedMdBytes() ([]byte, error) {
	return bindataRead(
		_plusUsedMd,
		"plus-used.md",
	)
}

func plusUsedMd() (*asset, error) {
	bytes, err := plusUsedMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plus-used.md", size: 246, mode: os.FileMode(420), modTime: time.Unix(1538310351, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xb5, 0xf6, 0xb4, 0x29, 0x7f, 0x12, 0x9, 0x5d, 0x70, 0xef, 0x87, 0x17, 0x61, 0xbe, 0x89, 0x47, 0xcd, 0x83, 0x7c, 0x35, 0x4b, 0x79, 0xad, 0xe, 0x83, 0x6e, 0xaa, 0xcd, 0x17, 0x81, 0x15}}
	return a, nil
}

var _stubMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xb1\x6a\xec\x30\x10\x45\x7b\x7f\xc5\x85\xd7\xbc\x2d\x62\xf7\x69\x43\x48\x9f\x36\x84\x30\x2b\x5d\x5b\x62\x65\xcd\x32\x1a\x7b\xf1\xdf\x07\x3b\x0e\xa4\x13\x17\x74\xce\x99\x27\xbc\x73\xe6\x7c\xa5\xc1\x15\x91\x85\x4e\x48\x29\xf0\x44\x34\x5f\xae\x08\x3a\xcf\xac\xde\xfa\xee\xe5\x7c\xc1\x93\x38\xc4\x88\x5c\x83\x9a\x31\x38\xd4\x90\xcd\x58\xb8\x4a\x75\xcc\x72\x23\xb2\x23\x89\xc5\x1f\xb2\x51\xe2\xc1\x0c\x1a\xd9\x77\x1d\xf0\xb6\xeb\x56\x16\xbd\xd3\x1a\x7c\xff\x11\x35\x2c\xbb\x41\x3c\x6b\xc5\xc7\x4a\xdb\xd0\x68\x59\x97\x56\xb6\xcf\xff\xc9\xfd\xde\x9e\x87\x61\xd2\x22\x75\xea\xd5\xa6\x21\x6a\x18\x38\x8e\x0c\x9e\x57\x7e\x4d\xda\x27\x9f\xcb\xbf\x33\x59\x6c\xbb\xf4\xdd\xeb\x81\x39\x27\xe4\x7a\x54\x8c\xb9\x10\x2d\xe9\x52\x22\xae\x44\xd3\x99\x9e\x72\x9d\xb0\xe9\x82\xc7\x31\x3f\xf6\x43\x7e\xcb\xf7\x9c\x49\xa3\x86\xdd\xfa\xb7\xe4\x9c\x2e\xc8\x23\x3c\xe5\x86\x07\x8d\x90\x0a\x1d\xc7\x1c\xb2\x14\xdc\x25\xdc\x64\x62\xff\x1d\x00\x00\xff\xff\x24\x8a\xad\xe6\x6a\x01\x00\x00")

func stubMdBytes() ([]byte, error) {
	return bindataRead(
		_stubMd,
		"stub.md",
	)
}

func stubMd() (*asset, error) {
	bytes, err := stubMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stub.md", size: 362, mode: os.FileMode(420), modTime: time.Unix(1538428444, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xf4, 0x84, 0x85, 0x96, 0x8f, 0x78, 0x89, 0x7c, 0x8c, 0x43, 0xae, 0x2b, 0x28, 0x5f, 0x4d, 0xf3, 0xef, 0xb6, 0xb4, 0xeb, 0x42, 0xa3, 0xce, 0x72, 0xee, 0xf6, 0xeb, 0xc1, 0xdf, 0x22, 0x3c}}
	return a, nil
}

var _wrongCommentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xc1\x0d\xc2\x30\x0c\x05\xd0\x55\xfe\x85\x23\x1d\x83\x09\x58\x20\x75\x7e\x71\x44\x62\x4b\xb6\x0b\x62\x7b\xa4\x5e\x9f\xde\x1d\x4f\x25\xc4\xd7\xa2\x15\x3a\x53\x62\xec\xc3\x5e\x28\x25\x6e\x89\x54\x3f\x67\x87\xb6\x0f\x2f\x12\x8f\xa0\x14\x0e\x8f\xd5\x6a\xc3\xc3\x03\x9d\x47\x1b\x33\xf1\xf3\x13\xd2\x0c\xa2\x94\xf7\xb5\x93\x52\xc3\x0d\x3b\xa7\x7f\xb7\x7f\x00\x00\x00\xff\xff\x82\x89\x6a\xf2\x6c\x00\x00\x00")

func wrongCommentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_wrongCommentFormatMd,
		"wrong-comment-format.md",
	)
}

func wrongCommentFormatMd() (*asset, error) {
	bytes, err := wrongCommentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wrong-comment-format.md", size: 108, mode: os.FileMode(420), modTime: time.Unix(1538428444, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe7, 0xfd, 0x73, 0x49, 0x20, 0x9e, 0x44, 0x25, 0x47, 0x46, 0xe4, 0xe4, 0x61, 0x38, 0x71, 0x5a, 0x9c, 0x19, 0x27, 0x24, 0x3d, 0x4, 0xdd, 0x45, 0xf7, 0x2e, 0x24, 0x9, 0xd9, 0xb8, 0xe8, 0x3c}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"comment-format.md": commentFormatMd,

	"missing-comment.md": missingCommentMd,

	"plus-used.md": plusUsedMd,

	"stub.md": stubMd,

	"wrong-comment-format.md": wrongCommentFormatMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"comment-format.md":       &bintree{commentFormatMd, map[string]*bintree{}},
	"missing-comment.md":      &bintree{missingCommentMd, map[string]*bintree{}},
	"plus-used.md":            &bintree{plusUsedMd, map[string]*bintree{}},
	"stub.md":                 &bintree{stubMd, map[string]*bintree{}},
	"wrong-comment-format.md": &bintree{wrongCommentFormatMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
