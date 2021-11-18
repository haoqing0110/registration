// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/hub/clusterrole/manifests/managedcluster-registration-clusterrole.yaml
// pkg/hub/clusterrole/manifests/managedcluster-work-clusterrole.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: open-cluster-management:managedcluster:registration
rules:
# Allow spoke registration agent to get/update coordination.k8s.io/lease
- apiGroups: ["coordination.k8s.io"]
  resources: ["leases"]
  #TODO: for backward compatible, we do not limit the resource name in release 2.3.
  #After release 2.3, we will limit the resource name.
  #resourceNames: ["managed-cluster-lease"]
  verbs: ["get", "update"]
# Allow agent to get/list/watch managed cluster addons
- apiGroups: ["addon.open-cluster-management.io"]
  resources: ["managedclusteraddons"]
  verbs: ["get", "list", "watch"]
# Allow agent to update the status of managed cluster addons
- apiGroups: ["addon.open-cluster-management.io"]
  resources: ["managedclusteraddons/status"]
  verbs: ["patch", "update"]
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["managedclusterscalars"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["managedclusterscalars/status"]
  verbs: ["patch", "update"]
`)

func pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYamlBytes() ([]byte, error) {
	return _pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYaml, nil
}

func pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYaml() (*asset, error) {
	bytes, err := pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/clusterrole/manifests/managedcluster-registration-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgHubClusterroleManifestsManagedclusterWorkClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: open-cluster-management:managedcluster:work
rules:
# Allow work agent to send event to hub
- apiGroups: ["", "events.k8s.io"]
  resources: ["events"]
  verbs: ["create", "patch", "update"]
# Allow work agent to get/list/watch/update manifestworks
- apiGroups: ["work.open-cluster-management.io"]
  resources: ["manifestworks"]
  verbs: ["get", "list", "watch", "update"]
# Allow work agent to update the status of manifestwork
- apiGroups: ["work.open-cluster-management.io"]
  resources: ["manifestworks/status"]
  verbs: ["patch", "update"]
`)

func pkgHubClusterroleManifestsManagedclusterWorkClusterroleYamlBytes() ([]byte, error) {
	return _pkgHubClusterroleManifestsManagedclusterWorkClusterroleYaml, nil
}

func pkgHubClusterroleManifestsManagedclusterWorkClusterroleYaml() (*asset, error) {
	bytes, err := pkgHubClusterroleManifestsManagedclusterWorkClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/clusterrole/manifests/managedcluster-work-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"pkg/hub/clusterrole/manifests/managedcluster-registration-clusterrole.yaml": pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYaml,
	"pkg/hub/clusterrole/manifests/managedcluster-work-clusterrole.yaml":         pkgHubClusterroleManifestsManagedclusterWorkClusterroleYaml,
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
	"pkg": {nil, map[string]*bintree{
		"hub": {nil, map[string]*bintree{
			"clusterrole": {nil, map[string]*bintree{
				"manifests": {nil, map[string]*bintree{
					"managedcluster-registration-clusterrole.yaml": {pkgHubClusterroleManifestsManagedclusterRegistrationClusterroleYaml, map[string]*bintree{}},
					"managedcluster-work-clusterrole.yaml":         {pkgHubClusterroleManifestsManagedclusterWorkClusterroleYaml, map[string]*bintree{}},
				}},
			}},
		}},
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
