package main

const (
	COMPANY_INDICATOR string = "company"
	LEO_INDICATOR	  string = "subject"
	NSPL_INDICATOR    string = "postcode"
	COMPANY_DIRECTORY string = "/home/yinghua/Documents/FYP1/FYP-data/company-data/company-data-full.csv"
	LEO_DIRECTORY     string = "/home/yinghua/Documents/FYP1/FYP-data/subject-data/institution-subject-data.csv"
	NSPL_DIRECTORY    string = "/home/yinghua/Documents/FYP1/FYP-data/postcode-data/UK-NSPL.csv"
)

// function to check error and print error messages
func checkErr(err error, message string) {
	if err != nil {
		panic(message + " err: " + err.Error())
	}
}

func main() {
	
	sequential_csv()
	concurrent_csv()

}

