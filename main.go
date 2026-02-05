package main

import (
    "fmt"
    "os"
    "strings"
)

func runCommand(command string, args []string, dpkgStatusFilePath, etcOsReleaseFilePath string, jsonOutput bool, programName string) {
    gardenLinuxVersion := readGardenLinuxVersion(etcOsReleaseFilePath)
    var cves []sourcePackageCve

    switch command {
    case "what-if":
        packagesToCheck := args
        cves = getCvesForPackageList(packagesToCheck, gardenLinuxVersion)
        printCves(cves, jsonOutput)
    case "check":
        dpkgSourcePackages := getDpkgSourcePackages(dpkgStatusFilePath)
        cves = getCvesForPackageList(dpkgSourcePackages, gardenLinuxVersion)
        printCves(cves, jsonOutput)
    case "executive-summary":
        dpkgSourcePackages := getDpkgSourcePackages(dpkgStatusFilePath)
        cves = getCvesForPackageList(dpkgSourcePackages, gardenLinuxVersion)
        fmt.Printf("This machine has %d potential security issues\nRun `%s check` to get the full list\n", len(cves), programName)
    default:
        fmt.Printf("Unknown command: %s\n", command)
    }
}

func main() {
    jsonOutput := strings.ToLower(os.Getenv("GLVD_CLIENT_JSON_OUTPUT")) == "true"

    devMode := os.Getenv("GLVD_CLIENT_DEV_MODE")
    var dpkgStatusFilePath string
    var etcOsReleaseFilePath string

    if len(devMode) > 0 {
        println("Running in dev mode")
        dpkgStatusFilePath = "test-data/var-lib-dpkg-status.txt"
        etcOsReleaseFilePath = "test-data/etc-os-release.txt"
    } else {
        dpkgStatusFilePath = "/var/lib/dpkg/status"
        etcOsReleaseFilePath = "/etc/os-release"
    }

    args := os.Args[1:]
    programName := os.Args[0]

    if len(args) == 0 {
        fmt.Printf("Usage: %s <command> <args>\nCommands: what-if, check, executive-summary\nArgs: List of source packages for command what-if\n", programName)
        os.Exit(0)
    }

    command := args[0]
    runCommand(command, args[1:], dpkgStatusFilePath, etcOsReleaseFilePath, jsonOutput, programName)
}
