package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func printCves(cves []sourcePackageCve, jsonOutput bool) {
	if jsonOutput {
		output, err := json.MarshalIndent(cves, " ", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	} else {
		for _, cve := range cves {
			fmt.Printf("%-18s %4.1f %-46s %-20s %-20s\n", cve.CveId, cve.BaseScore, cve.VectorString, cve.SourcePackageName, cve.SourcePackageVersion)
		}
	}
}
