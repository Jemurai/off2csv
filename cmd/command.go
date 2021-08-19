package cmd

import (
	utils "github.com/jemurai/fkit/utils"
	csvutils "github.com/jemurai/off2csv/utils"
)

func Report(file string) {
	findings := utils.BuildFindingsFromFile(file)
	csvutils.ReportFindings(findings)
}
