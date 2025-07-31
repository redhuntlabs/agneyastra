package main

import (
	"fmt"
	"log"

	flags "github.com/JA3G3R/agneyastra/flag"
	"github.com/JA3G3R/agneyastra/pkg/config"
	"github.com/JA3G3R/agneyastra/pkg/correlation"
	"github.com/JA3G3R/agneyastra/pkg/report"
	"github.com/JA3G3R/agneyastra/pkg/secrets"
	"github.com/JA3G3R/agneyastra/utils"
)

func main() {
	
	// fmt.Println("executing flags")
	utils.Init()
	flags.Execute()

	if config.Correlate {
		fmt.Println("Correlating data...")
		correlation.AddCorelationScore()
	}
	if config.SecretsExtract {
		fmt.Println("Extracting secrets...")
		secrets.ExtractSecrets()
	}
	finalReport, err := report.GlobalReport.ReportToJSON()
	if err != nil {
		log.Println("Error converting report to JSON: ", err)
		return
	}
	fmt.Printf("%v\n", finalReport)
	// fmt.Println("Generating HTML Report")
	err = report.GlobalReport.GenerateHTMLReport(config.ReportPath, config.TemplateFile)
	if err != nil {
		log.Println("Error generating HTML report: ", err)
		return
	}

}	


