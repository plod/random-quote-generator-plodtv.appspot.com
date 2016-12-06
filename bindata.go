// Code generated by go-bindata.
// sources:
// static/random.html
// DO NOT EDIT!

package random

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

var _randomHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x57\x69\x93\xd3\x38\x10\xfd\xbc\xf3\x2b\xb4\x1e\xb6\xec\x6c\xc6\x8e\x93\x39\x18\x32\x4e\x58\x8e\x85\x1a\x8a\x6b\x87\xa3\x8a\x65\xf9\xa0\xd8\xed\x58\x83\x6d\x19\x49\x9e\x10\xa8\xf9\xef\xdb\xb2\xe2\x23\x17\xa0\x54\xe2\x43\xad\xa7\xd7\xaf\x5b\x2d\x25\xf8\x3d\xe2\xa1\x5a\x16\x40\x12\x95\xa5\xd3\x83\x83\x40\x5f\x49\x4a\xf3\xf9\xc4\x82\xdc\x9a\xe2\x0b\xa0\xd1\xf4\x80\x90\x20\x03\x45\x49\x98\x50\x21\x41\x4d\xac\x52\xc5\xee\x39\xf6\xeb\x1e\xc5\x54\x0a\xd3\x2b\x9a\x47\x3c\x23\xff\x94\x5c\x01\x79\x41\xc3\x84\xe5\x10\x0c\x4c\x1f\x5a\xa1\x9d\x0c\x05\x2b\x14\x91\x22\x9c\x58\x89\x52\x85\x1c\x0f\x06\x21\x8f\xc0\xbb\xfe\x52\x82\x58\x7a\x21\xcf\x06\xe6\xd6\x3d\xf6\x86\xf8\xc9\x58\xee\x5d\x4b\x8b\xb0\x5c\xc1\x5c\x30\xb5\x9c\x58\x32\xa1\xa3\xd3\x33\x37\x79\xff\x3e\xff\x40\xd9\x83\xc7\x57\x6f\x5f\x8d\x5e\x7f\x7b\xf7\x34\x2b\x9f\x3f\x13\xe7\x0f\x9f\xbf\x7b\x73\xfd\xf4\xf2\x5f\xf9\xf8\xc3\xd3\xec\xf2\xd9\xf3\x9b\xd1\xec\x7c\x62\x91\x50\x70\x29\xb9\x60\x73\x96\x4f\x2c\x9a\xf3\x7c\x99\xf1\x52\x5a\xd3\x60\x60\x28\xad\xf8\xa5\x2c\xff\x4c\x12\x01\x71\x4b\x2f\xe6\xb9\x92\xde\x9c\xf3\x79\x0a\xb4\x60\xb2\xe2\x18\x4a\x79\x3f\xa6\x19\x4b\x97\x93\x57\x72\x41\xd3\xc8\x22\x02\x52\xe4\xa6\x96\x29\xc8\x04\x40\x59\x3f\xf0\xb8\x94\xe0\x69\x58\xba\x00\xc9\x33\xa8\x10\x47\x74\x76\xe2\x1f\xc3\xdd\xd3\x48\xfb\xdb\xe1\xa5\xe5\xad\x60\x0d\xe0\x8c\x47\xcb\xa3\x2a\x56\xe4\x7b\xf5\x22\xa3\x02\x9d\x1a\x13\xff\xa2\x7a\x4c\x80\xcd\x13\x35\x26\x43\xdf\xff\xc3\xbc\xb9\x6d\xc6\x99\x01\xba\xe9\xd9\x5d\xe3\xc0\x98\x18\x0f\x2e\x9a\x4e\x25\x68\x2e\x99\x62\x1c\x51\x67\x34\xfc\x3c\x17\xbc\xcc\x23\x32\x94\x35\x5e\x75\x91\x10\x6a\x93\x15\x8b\x88\xc9\x22\xa5\x08\x16\xa7\xf0\xd5\xd8\xe9\x3b\x37\x62\xc2\xd8\x8d\x89\xe0\x8b\x35\x8a\x5d\x86\xd5\x85\x0a\xc5\xc2\x14\x36\xfc\xa2\xa5\xe2\xc6\x6c\xc1\x22\x95\x8c\xc9\x89\xef\x17\xab\x29\x5a\x76\x6e\xc8\x53\x2e\xc6\x64\x91\x30\x05\xab\x4e\x2e\x22\x10\xae\xa0\x11\x2b\xe5\x98\x8c\x70\x14\xf1\xcd\x65\xd4\x20\x14\x34\x8a\x58\x3e\xd7\x7a\x35\xa0\xfc\xab\x8b\x39\x16\xf1\x05\x8a\x8a\x9f\x63\x3d\xc2\xd5\xfd\xe4\xd0\xf7\xfd\xae\xa8\xc9\x70\x43\x52\xc9\xbe\xc1\x98\x1c\x8f\x6a\x2c\xdd\x30\xa7\xc0\xad\xc3\x72\x7c\x56\x77\xd5\x08\xe3\x19\xc4\x5c\xc0\x11\xa9\x9f\x69\xac\x40\xec\x09\xd5\x13\x7c\x78\x60\xb2\xe6\x62\xd7\xcc\xc3\x51\xad\xa9\x6e\xbc\xa0\x21\x2e\x99\x31\xb9\xb7\x9e\x0b\xcd\xa4\xed\x2c\x21\x62\x40\x8e\x04\xad\xff\xe2\xa1\x1f\x59\x2d\x88\x89\x83\x2b\xea\xb4\xda\xe2\xbf\xc1\x77\x1d\x09\xb6\x91\x52\x88\x77\x01\x8d\x5a\x08\x05\x5f\x95\x4b\x53\x36\xd7\x59\xa3\xe7\xed\x1a\xe6\xf4\xa6\xb5\xdc\x91\x76\xba\x5d\x97\x52\xb1\x78\xe9\x36\x5c\x24\x2a\x01\xee\x0c\xd4\x02\x20\xef\xa2\x79\x19\x9d\xb3\xf0\x61\xa9\x54\x93\xca\x55\x2a\x16\x05\x50\x5c\x06\x21\x6a\x8a\xc5\xa2\x23\x76\x37\x3b\xd6\x7b\x78\xa9\x74\xa8\x37\x5f\xd7\x79\xdc\x3a\xac\x5b\x93\x77\xa7\x98\x56\xa7\xdd\x1e\x93\xb6\xdb\xd3\xae\x25\xf3\x49\x77\xc4\x2a\xf5\x0f\xe3\x38\xbe\xf8\xf9\x0a\xd7\x7e\x63\x75\x31\x15\x25\x18\x98\xe2\x1e\xe8\xf2\x50\x57\x2c\xb3\x60\xa7\x0d\x54\xb0\x5a\x98\xd3\x20\x19\x62\x61\xd2\x3f\xc9\x48\xdf\xe0\x0f\x46\x63\x1a\xcc\x8c\x7c\x61\x4a\xa5\x9c\x58\x1d\x45\xb1\x6e\x47\x13\x4b\x8b\x8e\xf5\x30\x60\xb5\x45\x4c\x49\x4c\x5d\xb5\x60\x0a\x33\x47\xd7\x3a\x86\x5f\x03\xf2\x53\xb0\x1c\x16\xd5\xfe\x62\x4d\x5f\xc2\xc2\x6c\x35\xed\xd8\x41\x45\x67\x50\xf3\x35\xfe\x0c\x1a\x87\xba\x15\xb9\x75\x2f\x2e\x73\x53\xc9\xe6\xa0\xcc\x0e\xf6\x48\x0b\xea\xf4\x3a\xf9\x50\xb7\x1b\x2a\x48\x0a\x9a\xb6\x24\x13\x62\xfb\xc3\xd1\xf1\xc9\xe9\xd9\xdd\xf3\x7b\x0f\x1e\x3e\x7a\xfc\xf7\x13\xfb\x62\xe7\x88\x2a\x40\xda\xfe\x70\x87\x01\xae\x43\xe2\x68\x2b\x86\x16\xfe\x05\x5e\x02\x72\x86\x97\x7e\x9f\xec\x62\xd0\x04\x9c\xf4\x27\x35\x95\x8f\x2f\xa8\x4a\xbc\x38\xe5\xc8\xba\xba\x15\x95\x1b\xe8\xc1\x9f\x64\x78\xd6\xfb\xb4\x3d\xe9\xed\xd6\x1b\x01\xaa\x14\xb9\x81\x6e\xed\x6f\x77\xaa\x54\xa9\xee\xf4\x36\xc8\xdd\xf1\xb0\xeb\xd9\x9b\x57\x2f\x1d\x6b\xf0\x45\x5b\x48\xeb\xa8\x19\xe6\xd0\x9d\xce\xb0\x98\x38\xd4\x4b\x21\x9f\xab\x04\xfd\x1e\x6e\x62\xd6\xd0\x8e\x9d\x0c\xed\x9e\xa7\xb7\x3d\xc7\x0a\x20\x9b\xbe\x16\x7c\x96\x42\x46\xae\x40\x09\x06\x37\xb8\x94\x48\x2c\xf0\xec\x81\x3b\x74\x30\xc0\x7e\xab\xb7\xed\x74\x8d\x34\x5a\x43\x7a\xc4\xb3\xa2\x44\x15\x89\xa4\x4b\x89\xab\x6e\xff\xf0\x5b\x48\x25\xec\x26\xa8\xe3\x27\xda\xdc\xc1\x48\x6e\x26\xd3\x5e\x3a\x7a\xe5\x21\x21\x3c\x51\x38\xdf\x89\xdd\xee\x68\xf6\xb8\x8b\x78\x44\xec\x2a\x34\x36\x59\x7b\x4d\x6e\xf7\x03\x77\x4b\xdb\x6a\x02\x7b\x73\xc3\xb4\x8f\xba\x68\x7b\xb0\xb0\x8a\x4a\x9e\x82\x97\xf2\xb9\x43\x3f\xfa\x9f\x7e\x20\x6d\x13\xa4\x3b\x95\xa5\xb7\x2a\xc0\x3d\x4f\x97\x74\xa7\xf7\x2b\x41\xa9\xc6\x55\x67\xc6\x3d\xd6\x5a\x6b\xac\xc0\x02\xde\x22\x26\x2a\x0d\xb9\x3e\x41\xbe\xbb\xba\xd4\x91\xc4\xa2\x99\x2b\xc7\xb6\xec\xfe\x6e\x06\x7d\xdb\x22\x76\xbf\x9d\xa3\x6f\xd7\x8f\xfa\xe8\x87\x4f\xce\x0d\xa3\xe4\xaf\x22\xe5\x91\xba\xe9\xd9\x7b\x09\x5b\x87\xa6\xaa\xa1\xae\x29\x0b\x3f\x3b\x4d\x9e\xef\x5b\xb3\xba\x2d\x18\x4a\xbd\xf0\x78\x01\xb9\xd3\x1c\x06\x57\x65\xb0\x3a\x04\x56\x6e\xdd\xd7\x4c\x27\x56\xbf\xf1\x71\x0f\x89\x9d\xa1\xdf\x5e\xd8\x28\xee\xaa\x00\xa2\xc2\x31\x8d\xe0\x32\xdf\xca\xc6\x2e\x54\x8b\x70\xc7\x21\xf8\xdf\xa0\xcc\x50\x3e\xd2\xf3\x04\x6e\x14\xcb\xfd\x7e\xae\x4d\x93\xb0\x08\x36\x27\x69\xeb\xc6\xc5\xc6\x40\x62\x1d\x36\x55\x9d\xfc\x92\xa0\xfb\xb0\xb4\x1f\x07\xe4\xb7\xae\x4b\xed\x41\x1a\xb7\x89\x6a\x93\xc3\x8d\xab\xfa\xa7\xf3\x7f\x00\x00\x00\xff\xff\x95\x3f\xfb\xb1\xfa\x0c\x00\x00")

func randomHtmlBytes() ([]byte, error) {
	return bindataRead(
		_randomHtml,
		"random.html",
	)
}

func randomHtml() (*asset, error) {
	bytes, err := randomHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "random.html", size: 3322, mode: os.FileMode(436), modTime: time.Unix(1481025099, 0)}
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
	"random.html": randomHtml,
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
	"random.html": &bintree{randomHtml, map[string]*bintree{}},
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

