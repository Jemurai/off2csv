package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/jemurai/fkit/finding"
)

// ReportFindings - report a set of findings to JIRA
func ReportFindings(findings []finding.Finding) {
	CSVHeaderRow()
	for i := 0; i < len(findings); i++ {
		finding := findings[i]
		FindingToCSVRow(finding)
	}
}

// Print a Header Row
func CSVHeaderRow() {
	fmt.Printf("Name,Description,Detail,Severity,Confidence,Fingerprint,Timestamp,Source,Location,CVSS,References,Cwes,Tags\n")
}

// Print a Finding to a CSV Row
func FindingToCSVRow(finding finding.Finding) {
	var refs string
	var cwes string
	var tags string
	fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", csvFriendly(finding.Name), csvFriendly(finding.Description),
		csvFriendly(finding.Detail), csvFriendly(finding.Severity), csvFriendly(finding.Confidence),
		csvFriendly(finding.Fingerprint), finding.Timestamp.Format(time.RFC3339Nano), csvFriendly(finding.Source),
		csvFriendly(finding.Location), fmt.Sprint(finding.Cvss), refs, cwes, tags)
}

func csvFriendly(instring string) string {
	return strings.Replace(instring, ",", "", -1)
}
