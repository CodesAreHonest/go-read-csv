package main 

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	_ "github.com/lib/pq"
)


func retrieve_without_channel(directory string, indicator string) {
	
	fmt.Printf("BEGIN retrieve data from %s files. \n", indicator);

	csvFile, err := os.Open(directory)
	checkErr(err, "Open CSV")

	defer csvFile.Close()
	
		// get the time before execution
	start := time.Now()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		_, err := reader.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

	}

	// obtain the time after execution
	fmt.Printf("FINISH retrieve all rows of data from %s files with %.5f seconds. \n", indicator, time.Since(start).Seconds())

}

func sequential_csv() { 
	
	// get the time before execution
	start := time.Now()
	
	retrieve_without_channel(LEO_DIRECTORY, LEO_INDICATOR);
	retrieve_without_channel(COMPANY_DIRECTORY, COMPANY_INDICATOR);
	retrieve_without_channel(NSPL_DIRECTORY, NSPL_INDICATOR);
	
		// obtain the time after execution
	fmt.Printf("%.5f seconds on retrieve all the data SEQUENTIALLY. \n", time.Since(start).Seconds())
}

/** 

yinghua@yinghua:~/gitRepo/go-read-csv/src/main$ go build *.go 
yinghua@yinghua:~/gitRepo/go-read-csv/src/main$ time go run *.go 

BEGIN retrieve data from subject files. 
FINISH retrieve all rows of data from subject files with 0.09179s seconds. 
BEGIN retrieve data from company files. 
FINISH retrieve all rows of data from company files with 32.64937s seconds. 
BEGIN retrieve data from postcode files. 
FINISH retrieve all rows of data from postcode files with 13.07156s seconds. 
45.81286s seconds on retrieve all the data SEQUENTIALLY. 

real	0m46.050s
user	0m46.651s
sys	0m0.612s

**/ 

