package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func getApiBaseUrl() string {
	url := os.Getenv("GLVD_API_BASE_URL")
	if url == "" {
		url = "https://security.gardenlinux.org"
	}
	return url
}

func getCvesForPackageList(dpkgSourcePackages []string, gardenLinuxVersion string) []sourcePackageCve {
	apiBaseUrl := getApiBaseUrl()
	client := &http.Client{}
	requestPayload, _ := json.Marshal(payload{PackageNames: dpkgSourcePackages})
	req, err := http.NewRequest("PUT", apiBaseUrl+"/v1/cves/"+gardenLinuxVersion+"/packages?sortBy=cveId&sortOrder=ASC", bytes.NewBuffer(requestPayload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results []sourcePackageCve
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}
