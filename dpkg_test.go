package main

import (
	"os"
	"reflect"
	"testing"
)

// Test buildDpkgStructure
func TestBuildDpkgStructure(t *testing.T) {
	input := `Package: foo
Status: install ok installed
Source: foo
Package: bar
Status: install ok installed
Source: bar 1.2.3
Package: baz
Status: deinstall ok config-files
Source: baz
`
	want := []dpkgPackage{
		{Package: "foo", Status: "install ok installed", Source: "foo"},
		{Package: "bar", Status: "install ok installed", Source: "bar"},
		{Package: "baz", Status: "deinstall ok config-files", Source: "baz"},
	}
	got := buildDpkgStructure(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildDpkgStructure() = %v, want %v", got, want)
	}
}

// Test removePotentialVersionSuffix
func TestRemovePotentialVersionSuffix(t *testing.T) {
	input := "bar 1.2.3"
	want := "bar"
	got := removePotentialVersionSuffix(input)
	if got != want {
		t.Errorf("removePotentialVersionSuffix(%q) = %q, want %q", input, got, want)
	}
}

// Test getDpkgSourcePackages
func TestGetDpkgSourcePackages(t *testing.T) {
	// Create a temporary file with dpkg status content
	content := `Package: foo
Status: install ok installed
Source: foo
Package: bar
Status: install ok installed
Source: bar 1.2.3
Package: baz
Status: deinstall ok config-files
Source: baz
`
	tmpfile, err := os.CreateTemp("", "dpkg-status")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	got := getDpkgSourcePackages(tmpfile.Name())
	want := []string{"bar", "foo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getDpkgSourcePackages() = %v, want %v", got, want)
	}
}

// Test readGardenLinuxVersion
func TestReadGardenLinuxVersion(t *testing.T) {
	content := `NAME="Garden Linux"
GARDENLINUX_VERSION=123.45
ID=gardenlinux
`
	tmpfile, err := os.CreateTemp("", "os-release")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	got := readGardenLinuxVersion(tmpfile.Name())
	want := "123.45"
	if got != want {
		t.Errorf("readGardenLinuxVersion() = %q, want %q", got, want)
	}
}
