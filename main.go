package main

import (
	//"encoding/json"
	"fmt"
	//"io"
	"log"
	"runtime"
	//"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Starting Go Routines")

	//parse CSV and read lines
	properties, lines, err := readCSV("./data1.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	//Loop through lines in CSV file and construct json objects with the data
	for i, line := range lines {
		if i == 0 {
			continue //skip first line
		}
		wg.Add(1)

		convert(properties, line)
	}

	fmt.Println("Waiting To Finish")
	wg.Wait()

	//just testing
	/*for _, i := range json_objs {
		dec := json.NewDecoder(strings.NewReader(i))
		for {
			var m CSVRow
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal("err")
			}
			fmt.Println("\n", m)
		}
	}*/

	fmt.Println("\nTerminating Program")

}
