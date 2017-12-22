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
	fmt.Printf("FINISH retrieve all rows of data from %s files with T%.5fs seconds. \n", indicator, time.Since(start).Seconds())

}

func sequential_csv() { 
	
	// get the time before execution
	start := time.Now()
	
	retrieve_without_channel(LEO_DIRECTORY, LEO_INDICATOR);
	retrieve_without_channel(COMPANY_DIRECTORY, COMPANY_INDICATOR);
	retrieve_without_channel(NSPL_DIRECTORY, NSPL_INDICATOR);
	
		// obtain the time after execution
	fmt.Printf("T%.5fs seconds on retrieve all the data SEQUENTIALLY. \n", time.Since(start).Seconds())
}

