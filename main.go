package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type FileLocation struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

type ReportingPlan struct {
	PlanName       string `json:"plan_name"`
	PlanIdType     string `json:"plan_id_type"`
	PlanId         string `json:"plan_id"`
	PlanMarketType string `json:"plan_market_type"`
}

type ReportingStructure struct {
	ReportingPlans    []ReportingPlan `json:"reporting_plans"`
	InNetworkFiles    []FileLocation  `json:"in_network_files"`
	AllowedAmountFile FileLocation    `json:"allowed_amount_file"`
}

type Record map[string]interface{}

func main() {
	fileName := "data/2024-04-01_anthem_index.json"
	outFile, err := os.Create("data/solution.txt")
	planNyPattern := regexp.MustCompile("\\sNY\\s")
	planPpoPattern := regexp.MustCompile("\\sPPO\\s")
	fileNyPattern := regexp.MustCompile("NY")
	s3Pattern := regexp.MustCompile("s3\\.amazonaws\\.com")
	// Keep track of the files that have already been written to the output file
	fileTracker := map[string]int{}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		var record ReportingStructure
		// Each line is a separate JSON object within the larger JSON object
		line, err := reader.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Remove the trailing newline and comma from the line
		err = json.Unmarshal(line[:len(line)-2], &record)
		if err != nil {
			continue
		}
		for _, plan := range record.ReportingPlans {
			planHasNy := planNyPattern.MatchString(plan.PlanName)
			planHasPpo := planPpoPattern.MatchString(plan.PlanName)
			if planHasNy && planHasPpo {
				for _, file := range record.InNetworkFiles {
					fileHasS3 := s3Pattern.MatchString(file.Location)
					fileHasNy := fileNyPattern.MatchString(file.Location)
					if fileHasS3 && fileHasNy {
						if _, ok := fileTracker[file.Description]; !ok {
							_, err := outFile.WriteString(file.Location + "\n")
							if err != nil {
								fmt.Println("Error writing to file:", err)
							}
						} else {
							continue
						}
					}
				}
			}
		}
	}
}
