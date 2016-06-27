# CSV-to-JSON
Code reads in a CSV file and outputs the data into an array of json objects.

## Usage

These are initial commits and code will be upgraded to make functionality more robust and intuitive.

configs.json is where you would map your CSV file structure to a GO struct type

To run:
```linux
    go build
    ./CSV-to-JSON
```

To see benchmarks on your data file:
```linux
    go test -bench=.
```

The sample data file contains 36,634 records in Florida for 2012 from a sample company that implemented an agressive growth plan in 2012. 
