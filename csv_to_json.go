package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

//use getJson method to convert current CSV line to json object Append json object to slice
func convert(properties []string, data []string) {
	defer wg.Done()

	c := make(chan string)
	go func() { c <- getJson(properties, data) }()
	obj := <-c
	json_objs = append(json_objs, obj)
}

//construct json object based on properties
func getJson(properties []string, data []string) string {
	fields := map[string]string{}
	for i, field := range properties {
		fields[field] = data[i]
	}
	jsonBytes, err := json.MarshalIndent(fields, "", " ")
	if err != nil {
		return ""
	}

	_json := (string(jsonBytes))
	return _json
}

//read lines in CSV return properties from first CSV line, lines, and error
func readCSV(file_path string) ([]string, [][]string, error) {
	csvFile, err := os.Open(file_path)
	if err != nil {
		return nil, nil, err
	}

	r := csv.NewReader(csvFile)

	lines, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	//first line of CSV file that defines our json object properties
	properties := lines[0]

	return properties, lines, nil
}
