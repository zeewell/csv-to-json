package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"strings"
)

func main() {
	csvFile, err := os.Open("./data.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}

		//this is a temp csv row add the details then append it to the csv_rows slice
		var temp CSVRow
		temp.PolicyID = line[0]
		temp.Statecode = line[1]
		temp.County = line[2]
		temp.Eq_site_limit = line[3]
		temp.Hu_site_limit = line[4]
		temp.Fl_site_limit = line[5]
		temp.Fr_site_limit = line[6]
		temp.Tiv_2011 = line[7]
		temp.Tiv_2012 = line[8]
		temp.Eq_site_deductible = line[9]
		temp.Hu_site_deductible = line[10]
		temp.Fl_site_deductible = line[11]
		temp.Fr_site_deductible = line[12]
		temp.Point_latitude = line[13]
		temp.Point_longitude = line[14]
		temp.Line = line[15]
		temp.Construction = line[16]
		temp.Point_granularity = line[17]
		csv_rows = append(csv_rows, temp)
	}

	_json, err := json.Marshal(csv_rows)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(_json)
	fmt.Println("\n\n")

}
