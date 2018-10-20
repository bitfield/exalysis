// Code generated by go-bindata. DO NOT EDIT.
// sources:
// go_routines.md (129B)
// ifs_to_switch.md (75B)
// loop_rune_not_byte.md (152B)
// maprune.md (195B)
// move_map.md (148B)
// try_switch.md (92B)
// type_conversion.md (149B)
// unicode.md (123B)
// unicode_loop.md (99B)

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

var _go_routinesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xbd\x0d\xc2\x30\x14\x45\xe1\x9e\x29\xee\x02\x64\x05\x1a\x1a\x46\xa0\x7c\x71\x2e\xb1\x15\xec\x17\xbd\x1f\x45\xd9\x1e\x91\xfe\x9c\xef\x8e\x17\x9c\xc4\xa9\x09\x31\x22\xbd\x8d\x15\xab\xc2\x34\xa3\x0d\xfa\x84\xb7\x26\x76\xd3\x59\xe6\xef\x89\x43\x46\x70\x41\x28\xba\x6c\x44\x0b\x78\xee\x34\x7c\xc4\x63\xc2\xb3\x2d\x97\x54\x2a\xcb\x86\xa8\xbc\xcd\x1c\xa5\x76\xb1\xcd\x51\xf5\x40\xcf\x52\xaf\x96\xf6\x7f\x9b\x3f\x7e\x01\x00\x00\xff\xff\x3a\x05\x57\x27\x81\x00\x00\x00")

func go_routinesMdBytes() ([]byte, error) {
	return bindataRead(
		_go_routinesMd,
		"go_routines.md",
	)
}

func go_routinesMd() (*asset, error) {
	bytes, err := go_routinesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go_routines.md", size: 129, mode: os.FileMode(420), modTime: time.Unix(1540032141, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xda, 0x30, 0x9a, 0x4e, 0x53, 0x1c, 0xb2, 0x9, 0xa5, 0x94, 0x2b, 0xb5, 0x49, 0x3b, 0x85, 0x6b, 0x97, 0x91, 0x2c, 0x8f, 0xca, 0xba, 0x39, 0x6c, 0xe9, 0x76, 0xe9, 0x59, 0x11, 0x82, 0x60, 0x48}}
	return a, nil
}

var _ifs_to_switchMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xc1\x09\x85\x30\x10\x04\xd0\x56\xa6\x81\xff\x9b\xb0\x91\x2c\x71\x64\x17\xc9\xe6\xb0\x13\xc4\xee\x05\x4f\xde\xdf\x0f\x9b\xb3\x9f\x98\x4b\x90\x13\xad\xae\x50\xf7\x86\x92\x89\x83\x29\x68\x62\x15\x11\x59\xa2\xed\xaf\x1a\x96\x37\x5a\x1c\x1f\x56\xff\x27\x00\x00\xff\xff\x20\x27\x3c\x25\x4b\x00\x00\x00")

func ifs_to_switchMdBytes() ([]byte, error) {
	return bindataRead(
		_ifs_to_switchMd,
		"ifs_to_switch.md",
	)
}

func ifs_to_switchMd() (*asset, error) {
	bytes, err := ifs_to_switchMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ifs_to_switch.md", size: 75, mode: os.FileMode(420), modTime: time.Unix(1538428401, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa4, 0x86, 0x72, 0xff, 0x80, 0xc8, 0xab, 0x94, 0x30, 0x68, 0x2b, 0xeb, 0x8a, 0x7c, 0x94, 0xc0, 0xd3, 0x8, 0xe6, 0xb8, 0x75, 0x6e, 0xb3, 0x39, 0x8c, 0xd7, 0x4d, 0x2f, 0x2, 0x9d, 0xc4, 0xc9}}
	return a, nil
}

var _loop_rune_not_byteMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcd\x41\xae\x83\x30\x0c\x84\xe1\xfd\x3b\xc5\xac\x23\xbd\xde\xa3\xfb\x1e\x20\xae\x71\xc1\x52\x70\x90\xed\x80\xb8\x7d\x15\x75\xff\xcd\xfc\xff\x78\xa6\x38\xa5\xda\x8a\x7e\x8a\x83\x50\x23\x5d\x6d\xad\xb8\xb4\x35\x1c\xde\x4f\x65\x41\xf5\x61\x52\x03\xd7\xa6\xbc\x41\x03\x04\xee\xfb\xd1\x24\x05\xbc\x91\x13\xe7\x5c\xdb\x02\x26\xc3\x1f\x77\x0b\x8d\x44\xff\xa0\x94\x7d\xb4\xd4\xa3\x49\x29\x78\xdf\x29\xf1\xc0\xcb\x6f\x8c\x98\xd1\x79\x1b\x50\x8b\x14\x5a\x26\xff\x89\x6f\x00\x00\x00\xff\xff\x9c\x11\x9d\x85\x98\x00\x00\x00")

func loop_rune_not_byteMdBytes() ([]byte, error) {
	return bindataRead(
		_loop_rune_not_byteMd,
		"loop_rune_not_byte.md",
	)
}

func loop_rune_not_byteMd() (*asset, error) {
	bytes, err := loop_rune_not_byteMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "loop_rune_not_byte.md", size: 152, mode: os.FileMode(420), modTime: time.Unix(1538236752, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x86, 0x33, 0xbb, 0x93, 0xb2, 0x6a, 0x28, 0x25, 0x18, 0x75, 0x64, 0x3a, 0xe2, 0x56, 0x8b, 0x2, 0xc, 0xda, 0x34, 0x4d, 0x5d, 0xc8, 0x66, 0xe2, 0x7c, 0xbc, 0xe6, 0xeb, 0x34, 0xb6, 0x23}}
	return a, nil
}

var _mapruneMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\x41\x4e\x43\x31\x0c\x84\xe1\x3d\xa7\x98\x0d\xea\x8a\x1c\x80\x5d\x8f\x81\x10\x52\x42\x32\x6d\x4d\x83\xfd\x88\xed\xa2\xde\x1e\xf5\x89\xf5\x7c\x1a\xfd\x2f\x78\xb3\x44\xb7\x9c\x03\xe9\x44\x43\xfd\x6e\xdb\xfb\x4a\xe5\x87\x68\x54\x88\x7a\xb0\x0d\xd8\x09\xf5\xd9\x2b\x4e\xb6\x30\x64\xb1\x07\xa6\xd9\x35\xb7\xc7\x12\x17\xc2\xbb\x2d\xe2\xd6\x66\x72\x47\x0d\xf5\xf1\x52\x5f\xf1\xa4\x86\xb8\x6f\x44\x37\xbd\x71\xb9\x98\x42\xc9\xc1\x51\x70\xfc\x57\x10\x47\x5f\x6c\xc1\x81\x5f\x89\x0b\x5c\xf4\x3c\x89\x9f\xb4\xa0\x83\xe5\x5c\x70\x38\x1e\xf0\x95\x1e\x98\x72\xdd\x4b\x3f\xef\xc1\x5a\xfe\x02\x00\x00\xff\xff\x89\x51\xc0\xa3\xc3\x00\x00\x00")

func mapruneMdBytes() ([]byte, error) {
	return bindataRead(
		_mapruneMd,
		"maprune.md",
	)
}

func mapruneMd() (*asset, error) {
	bytes, err := mapruneMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "maprune.md", size: 195, mode: os.FileMode(420), modTime: time.Unix(1538428428, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2d, 0x96, 0x55, 0x1b, 0xe7, 0x98, 0x73, 0xc9, 0x82, 0x75, 0xcb, 0xc, 0x94, 0xfd, 0xcf, 0xdc, 0xdd, 0x74, 0x85, 0x1d, 0xab, 0x26, 0x72, 0xf, 0x7d, 0x24, 0xbf, 0xa4, 0xbc, 0xd4, 0x10, 0x9f}}
	return a, nil
}

var _move_mapMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcd\x41\x0e\xc2\x30\x0c\x44\xd1\x3d\xa7\x98\x0b\xd0\xcb\xc0\x86\x5d\x2d\xd7\x34\x16\xd4\xae\x92\x09\xa8\xb7\x47\x41\x82\xfd\xfb\xfa\x67\xdc\xb2\x43\xaa\x41\xab\x09\x3d\x56\x08\x1a\x85\xae\xd8\x64\x87\x47\xf3\xc5\xc0\x62\x98\x2f\x9a\xd5\x66\xdc\x7b\x28\x3d\x63\xc2\xb5\x1e\xd8\xf2\x35\x22\x16\x6f\xd0\x5c\x0c\xd9\xf9\x4f\x7e\x14\xa7\x96\x70\x62\x98\xb1\xb1\x05\x19\xcf\x03\x19\x6a\x78\x3b\xcb\x57\xef\xa2\x0f\x59\x6d\xfa\x04\x00\x00\xff\xff\xdb\xe5\x4a\xc7\x94\x00\x00\x00")

func move_mapMdBytes() ([]byte, error) {
	return bindataRead(
		_move_mapMd,
		"move_map.md",
	)
}

func move_mapMd() (*asset, error) {
	bytes, err := move_mapMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "move_map.md", size: 148, mode: os.FileMode(420), modTime: time.Unix(1538248324, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0x1d, 0x8f, 0xd7, 0x76, 0xf7, 0x90, 0xb3, 0x24, 0x4f, 0x51, 0xd9, 0xdc, 0xf, 0x48, 0xb9, 0xaf, 0xb2, 0x32, 0xb5, 0xfd, 0x6e, 0x73, 0xda, 0xb0, 0xa6, 0x9d, 0xc1, 0x2, 0x18, 0xd6, 0x27}}
	return a, nil
}

var _try_switchMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc8\xd1\x09\x83\x21\x0c\x04\xe0\x55\x6e\x81\x0e\xd2\x2d\x0c\xfe\x89\x3d\xb0\x51\x4c\x44\xdc\xbe\xf4\xf5\x7b\xe1\x6d\xb8\x63\x43\x96\x62\x4f\xd8\x58\x60\x62\x07\xbd\x41\x50\xe2\x30\xeb\xa7\x80\x1e\xa9\xf2\x60\xd8\x5f\xbf\x32\x0b\x0e\x7b\x07\xbd\x2e\x95\x50\xc4\x54\x7d\x10\x6c\x4e\x63\x15\xcf\x7e\x7f\x01\x00\x00\xff\xff\x0e\x2b\xc4\xd7\x5c\x00\x00\x00")

func try_switchMdBytes() ([]byte, error) {
	return bindataRead(
		_try_switchMd,
		"try_switch.md",
	)
}

func try_switchMd() (*asset, error) {
	bytes, err := try_switchMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "try_switch.md", size: 92, mode: os.FileMode(420), modTime: time.Unix(1538428428, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8e, 0x85, 0x1, 0xc4, 0xba, 0x97, 0x80, 0x68, 0xeb, 0x66, 0x93, 0x9, 0xec, 0xcd, 0x71, 0x81, 0xed, 0x27, 0x77, 0x59, 0x88, 0x5a, 0xf0, 0x2f, 0x64, 0xe6, 0x20, 0x98, 0x62, 0x54, 0xea, 0xbb}}
	return a, nil
}

var _type_conversionMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x31\x0e\x83\x30\x0c\x46\xe1\xbd\xa7\xf8\x37\xa6\xe6\x0e\xdc\xa1\x4b\xb7\x44\x89\x01\xab\xc8\x46\xb6\x03\xe2\xf6\x55\xc4\xfa\xf4\xe9\xbd\xf1\xb1\x1b\xa1\x28\xa7\x72\x43\x17\xa1\x4a\xee\x65\xc4\xfb\x20\x54\x95\x93\xcc\x59\xc5\x13\xbe\xda\x51\xb5\xef\x0d\x97\xda\x0f\x17\xc7\xf6\xa8\x6c\x5d\x28\x83\xc5\x83\x4a\x83\x2e\xc8\x1e\xc6\xb2\x8e\x86\x57\x6c\x84\xbc\xa8\x65\xec\xaa\x47\xc2\x8c\xe1\xc1\x8e\x6a\x54\x82\xda\xb3\xa2\xb4\x26\x4c\xf3\x94\xfe\x01\x00\x00\xff\xff\xc2\x76\x2e\x41\x95\x00\x00\x00")

func type_conversionMdBytes() ([]byte, error) {
	return bindataRead(
		_type_conversionMd,
		"type_conversion.md",
	)
}

func type_conversionMd() (*asset, error) {
	bytes, err := type_conversionMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type_conversion.md", size: 149, mode: os.FileMode(420), modTime: time.Unix(1538428428, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7c, 0xe, 0xb2, 0x9, 0xcf, 0xb0, 0x4b, 0x6d, 0x24, 0x44, 0xae, 0x6e, 0xb8, 0xae, 0xe9, 0x6e, 0x32, 0x98, 0x1a, 0x86, 0x8f, 0x3, 0xfb, 0x58, 0xdd, 0x0, 0xc4, 0xa, 0xf, 0x96, 0xad, 0xfa}}
	return a, nil
}

var _unicodeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xca\x41\x0a\x02\x31\x0c\x05\xd0\xab\xfc\x8d\x4b\x05\x8f\x23\x22\x34\x36\xbf\x4e\x70\x68\x86\x26\xf5\xfc\x32\xca\x6c\x1f\xef\x8c\x9b\x4f\x54\x9f\xab\x62\x91\x0f\x21\x58\xdd\xdf\x90\x44\x99\xdd\xaa\x2b\x2f\xa7\xfb\xf5\x11\x05\xd6\xc3\x94\xc8\x85\x68\x3e\xf6\xb6\xed\x96\x14\x85\x37\x94\xc8\x61\xfd\x15\x47\x7f\xb2\xf9\xf8\xf7\x5f\x4d\x87\xf5\x3a\x28\x41\xc4\x46\xea\x37\x00\x00\xff\xff\x71\xe7\x46\x8e\x7b\x00\x00\x00")

func unicodeMdBytes() ([]byte, error) {
	return bindataRead(
		_unicodeMd,
		"unicode.md",
	)
}

func unicodeMd() (*asset, error) {
	bytes, err := unicodeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unicode.md", size: 123, mode: os.FileMode(420), modTime: time.Unix(1538428428, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0x28, 0xdd, 0x43, 0xd4, 0x2, 0x2f, 0x2a, 0x5, 0xa6, 0x2e, 0xe8, 0x80, 0x72, 0xc6, 0x76, 0xed, 0x37, 0xc4, 0xc, 0x22, 0x92, 0xec, 0x40, 0x29, 0xc3, 0xbb, 0xc7, 0x35, 0x6e, 0xef, 0x2b}}
	return a, nil
}

var _unicode_loopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x41\x0a\x42\x31\x0c\x45\xd1\xad\xbc\x89\x43\x05\x97\x23\x22\x34\xa4\x0f\x7f\xb0\x24\xa5\x49\x5d\xbf\xfc\x81\xe3\x7b\xcf\x15\x8f\xd8\xd0\xd8\xa3\xe3\x90\x2f\x21\x18\x11\x1f\x48\xa1\x6d\x37\x8d\xce\xdb\xe5\x79\x7f\x65\x43\x05\x16\xe7\x10\x25\x5a\xd6\x32\x7f\xe7\x3f\x99\xa3\x0e\x9e\x72\x9e\x9b\xb9\x2e\x4a\x12\x39\xc9\xfe\x0b\x00\x00\xff\xff\x94\x62\xa1\xb6\x63\x00\x00\x00")

func unicode_loopMdBytes() ([]byte, error) {
	return bindataRead(
		_unicode_loopMd,
		"unicode_loop.md",
	)
}

func unicode_loopMd() (*asset, error) {
	bytes, err := unicode_loopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unicode_loop.md", size: 99, mode: os.FileMode(420), modTime: time.Unix(1538428428, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6, 0xcb, 0xa2, 0x9, 0xd4, 0x29, 0x1d, 0xf2, 0x29, 0x96, 0xf6, 0x79, 0x74, 0x43, 0xc3, 0x92, 0xab, 0x87, 0x97, 0x57, 0xa4, 0x70, 0xcf, 0x1f, 0x53, 0x5, 0x2a, 0xac, 0x8f, 0x54, 0xd8, 0x38}}
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
	"go_routines.md": go_routinesMd,

	"ifs_to_switch.md": ifs_to_switchMd,

	"loop_rune_not_byte.md": loop_rune_not_byteMd,

	"maprune.md": mapruneMd,

	"move_map.md": move_mapMd,

	"try_switch.md": try_switchMd,

	"type_conversion.md": type_conversionMd,

	"unicode.md": unicodeMd,

	"unicode_loop.md": unicode_loopMd,
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
	"go_routines.md":        &bintree{go_routinesMd, map[string]*bintree{}},
	"ifs_to_switch.md":      &bintree{ifs_to_switchMd, map[string]*bintree{}},
	"loop_rune_not_byte.md": &bintree{loop_rune_not_byteMd, map[string]*bintree{}},
	"maprune.md":            &bintree{mapruneMd, map[string]*bintree{}},
	"move_map.md":           &bintree{move_mapMd, map[string]*bintree{}},
	"try_switch.md":         &bintree{try_switchMd, map[string]*bintree{}},
	"type_conversion.md":    &bintree{type_conversionMd, map[string]*bintree{}},
	"unicode.md":            &bintree{unicodeMd, map[string]*bintree{}},
	"unicode_loop.md":       &bintree{unicode_loopMd, map[string]*bintree{}},
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
