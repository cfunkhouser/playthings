package main

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
)

// Package metadata for indexing.
type Package struct {
	Name string `json:"name"`
}

func (p Package) String() string {
	return p.Name
}

type dependencyMap map[string]map[string]*Package

// PackageIndex is a directed graph of Package dependencies.
type PackageIndex struct {
	sync.RWMutex

	// Map of packages stored internally
	packages     map[string]*Package
	dependsOn    map[string]map[string]*Package
	dependedOnBy map[string]map[string]*Package
}

// initPackage initializes an empty node in the index if the package did not
// already exist. The lock must be held by the caller.
func (i *PackageIndex) initPackage(ps ...Package) {
	for _, p := range ps {
		if _, ok := i.packages[p.Name]; !ok {
			i.packages[p.Name] = &p
			i.dependsOn[p.Name] = make(map[string]*Package)
			i.dependedOnBy[p.Name] = make(map[string]*Package)
		}
	}
}

// AddPackage metadata to the index.
func (i *PackageIndex) AddPackage(p Package) {
	i.Lock()
	defer i.Unlock()
	i.initPackage(p)
}

func (i *PackageIndex) HasPackage(p Package) bool {
	i.RLock()
	defer i.RUnlock()
	_, has := i.packages[p.Name]
	return has
}

// AddDependency mapping to the index. Both pkg and dependsOn are inserted into
// the index if they do not exist.
func (i *PackageIndex) AddDependency(pkg, dep Package) {
	i.Lock()
	defer i.Unlock()
	i.initPackage(pkg, dep)
	i.dependsOn[pkg.Name][dep.Name] = i.packages[dep.Name]
	i.dependedOnBy[dep.Name][pkg.Name] = i.packages[dep]
}

// Dependencies for a package.
func (i *PackageIndex) Dependencies(pkg Package) []Package {
	i.RLock()
	defer i.RUnlock()
	if deps, ok := i.dependencies[pkg]; ok {
		return deps
	}
	return []Package{}
}

// ErrDependedUpon is returned by Remove when the Package to be removed is still
// a dependency of one or more packages still in the index.
var ErrDependedUpon = errors.New("package is still depended on")

// Remove a Package from the index.
func (i *PackageIndex) Remove(pkg Package) error {

	return nil
}

func (i *PackageIndex) String() string {
	i.RLock()
	defer i.RUnlock()

	var keys []Package
	for pkg := range i.dependencies {
		keys = append(keys, pkg)
	}

	var buf bytes.Buffer
	for _, pkg := range keys {
		for _, dep := range i.dependencies[pkg] {
			fmt.Fprintf(&buf, "%q -> %q\n", pkg, dep)
		}
	}
	return buf.String()
}

func main() {
	index := PackageIndex{
		dependencies: make(map[Package][]Package),
	}

	index.AddDependency(Package{"X"}, Package{"Y"})
	index.AddDependency(Package{"Y"}, Package{"Z"})
	index.AddDependency(Package{"X"}, Package{"Z"})

	fmt.Printf("Index:\n%v", &index)
}
