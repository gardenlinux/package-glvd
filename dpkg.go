package main

import (
    "log"
    "os"
    "slices"
    "strings"
)

func buildDpkgStructure(dpkgStatusFileContents string) []dpkgPackage {
    var packages []dpkgPackage

    lines := strings.Split(dpkgStatusFileContents, "\n")

    for _, line := range lines {
        if strings.HasPrefix(line, "Package: ") {
            pkg := strings.Replace(line, "Package: ", "", 1)
            packages = append(packages, dpkgPackage{Package: pkg})
        }

        if strings.HasPrefix(line, "Status: ") {
            packages[len(packages)-1].Status = strings.Replace(line, "Status: ", "", 1)
        }

        if strings.HasPrefix(line, "Source: ") {
            sourcePackageNameWithPotentialVersion := strings.Replace(line, "Source: ", "", 1)
            sourcePackageName := removePotentialVersionSuffix(sourcePackageNameWithPotentialVersion)
            packages[len(packages)-1].Source = sourcePackageName
        }
    }

    return packages
}

func removePotentialVersionSuffix(input string) string {
    return strings.Split(input, " ")[0]
}

func getDpkgSourcePackages(dpkgStatusFilePath string) []string {
    dat, err := os.ReadFile(dpkgStatusFilePath)
    if err != nil {
        log.Fatal(err)
    }

    packages := buildDpkgStructure(string(dat))

    var pkgs []string

    for _, pkg := range packages {
        if pkg.Status == "install ok installed" {
            if len(pkg.Source) > 0 {
                pkgs = append(pkgs, pkg.Source)
            } else {
                pkgs = append(pkgs, pkg.Package)
            }
        }
    }

    // De-duplicate entries
    slices.Sort(pkgs)
    return slices.Compact(pkgs)
}

func readGardenLinuxVersion(osReleaseFilePath string) string {
    dat, err := os.ReadFile(osReleaseFilePath)
    if err != nil {
        log.Fatal(err)
    }

    lines := strings.Split(string(dat), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "GARDENLINUX_VERSION=") {
            return strings.Replace(line, "GARDENLINUX_VERSION=", "", 1)
        }
    }
    log.Fatal("Could not parse os-release, failed to identify Garden Linux version.")
    return ""
}