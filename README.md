# CSV-to-JSON
Code reads in a CSV file and outputs the data into an array of json objects.

## Usage
Swap data.csv with the csv of your liking. Next commit will accept command line argument for csv file path, web browser file upload capability and progress. 
To run:
```linux
    go build
    ./CSV-to-JSON
```

To see benchmarks on your data file:
```linux
    go test -bench=.
```

The sample data file contains 36,634 records.
