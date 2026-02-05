package main

type sourcePackageCve struct {
	CveId                string  `json:"cveId"`
	BaseScore            float32 `json:"baseScore"`
	VectorString         string  `json:"vectorString"`
	SourcePackageName    string  `json:"sourcePackageName"`
	SourcePackageVersion string  `json:"sourcePackageVersion"`
	GardenlinuxVersion   string  `json:"gardenlinuxVersion"`
	IsVulnerable         bool    `json:"isVulnerable"`
	CvePublishedDate     string  `json:"cvePublishedDate"`
}

type dpkgPackage struct {
	Package string
	Status  string
	Source  string
}

type payload struct {
	PackageNames []string `json:"packageNames"`
}
