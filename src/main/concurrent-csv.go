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

func retrieve_data_with_channel(directory string, indicator string, msg chan string) {
	
	fmt.Printf("BEGIN retrieve data from %s files. \n", indicator);

	csvFile, err := os.Open(directory)
	checkErr(err, "Open CSV")

	defer csvFile.Close()
	
		// get the time before execution
	start := time.Now()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))
	

	for {
		
		_ , err := reader.Read()
		
		// Stop at EOF.
		if err == io.EOF {
			break
		}
	}

	// obtain the time after execution
	fmt.Printf("FINISH retrieve all rows of data from %s files with %.5fs seconds.", indicator, time.Since(start).Seconds())
	msg <- " " 

}

// select function
func goSelect(ch_company, ch_leo, ch_nspl chan string) {
	

	for i := 0; i < 3; i++ {

		select {
		case msg1 := <-ch_leo:
			fmt.Println(msg1)
		case msg2 := <-ch_company:
			fmt.Println(msg2)
		case msg3 := <-ch_nspl:
			fmt.Println(msg3)

		}
	}
}

// This function read all CSV data concurrently 
func concurrent_csv() {
	
			// get the time before execution
	start := time.Now()
	
	// make three channel for three functions
	ch_company := make(chan string)
	ch_leo := make(chan string)
	ch_nspl := make(chan string)
	
	
	go retrieve_data_with_channel(LEO_DIRECTORY, LEO_INDICATOR, ch_leo);
	go retrieve_data_with_channel(COMPANY_DIRECTORY, COMPANY_INDICATOR, ch_company);
	go retrieve_data_with_channel(NSPL_DIRECTORY, NSPL_INDICATOR, ch_nspl);
	
	goSelect(ch_company, ch_leo, ch_nspl) 
	
		// obtain the time after execution
	fmt.Printf("T%.5fs seconds on retrieve all the data CONCURRENTLY. \n", time.Since(start).Seconds())
}

/**

BEGIN retrieve data from postcode files. 
BEGIN retrieve data from subject files. 
BEGIN retrieve data from company files. 
FINISH retrieve all rows of data from subject files with 0.12362s seconds. 
FINISH retrieve all rows of data from postcode files with 15.21926s seconds. 
FINISH retrieve all rows of data from company files with 36.22334s seconds. 
T36.22355s seconds on retrieve all the data CONCURRENTLY. 

real	0m36.478s
user	0m52.337s
sys	0m0.719s

**/



