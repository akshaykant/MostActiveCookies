package main

import (
	"csv2record/max"
	"csv2record/search"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)
type InputData struct {
	filename string
	date string
}

func getCommandLineData()(InputData, error){
	if len(os.Args) < 2 {
		return InputData{}, errors.New("A filepath argument is required. Use -f flag with filename and -d flag with date")
	}
	//flags for command-line arguments
	var filename string
	flag.StringVar(&filename, "f", "", "file name")
	var date string
	flag.StringVar(&date, "d", "", "date in UTC")

	flag.Parse()

	//Check if input file is valid CSV and exists
	var err error
	var enrichedFileName string
	if enrichedFileName, err = checkIfValidFile(filename); err != nil {
		return InputData{}, err
	}

	//check if date is in valid format
	var enrichedDate string
	if enrichedDate, err = checkIfValidDate(date, "date"); err != nil {
		exitGracefully(err)
	}

	return InputData{enrichedFileName, enrichedDate}, nil
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func checkIfValidFile(filename string) (string, error) {
	// Check if file extension is CSV
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return "", fmt.Errorf("File %s is not CSV. Please enter a valid CSV file. ", filename)
	}

	// Check if file does exist
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return "", fmt.Errorf("File %s does not exist ", filename)
	}
	return filename, nil
}

func checkIfValidDate(date string, flag string) (string, error) {
	var t time.Time
	var err error

	if flag == "date"{
		//Parse a time value from string in yyyy-mm-dd format
		t, err = time.Parse("2006-01-02", date)
		t.Nanosecond()
	}
	if flag =="timestamp"{
		t, err = time.Parse("2006-01-02T15:04:05", date[:19])
	}
	if err != nil {
		return "", err
	}

	return t.UTC().String()[:10], nil
}

func processCsvFile(data InputData) [][]string {

	file, err := os.Open(data.filename)

	check(err)

	defer file.Close()

	// Get Headers
	var headers, line []string
	cookieList := make([][]string, 0)

	reader := csv.NewReader(file)
	reader.Comma = ','

	headers, err = reader.Read()
	check(err)

	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			exitGracefully(err)
		}

		cookie, timestamp, err := processLine(headers, line)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}

		cookieList = append(cookieList, []string{cookie, timestamp})

	}

	return cookieList
}

func processLine(headers []string, dataList []string) (string, string, error) {
	if len(dataList) != len(headers) {
		return "", "", errors.New("Line doesn't match headers format. Skipping ")
	}

	cookie := dataList[0]
	timestamp := dataList[1]

	//check for date formatting
	var enrichedDate string
	var err error

	if enrichedDate, err = checkIfValidDate(timestamp, "timestamp"); err != nil {
		return "", "", err
	}

	return cookie, enrichedDate, nil
}

func check(e error) {
	if e != nil {
		exitGracefully(e)
	}
}

func main() {

	inputData, err := getCommandLineData()

	if err != nil {
		exitGracefully(err)
	}

	//Read the CSV file and convert in into slice/array
	cookieList := processCsvFile(inputData)

	//Search the input date in the array, assuming array is sorted by date.
	//Also date in array need to be converted for each timestamp into date before comparison
	//Binary Search modified
	evalList := search.EvalList{
		CookieList: cookieList,
		Date: inputData.date,
	}

	resultList , err := evalList.GetList()
	if err != nil {
		exitGracefully(err)
	}

	list := max.EvalList{
		CookieList: resultList,
	}

	mostOccurrence, err := list.ActiveCookie()

	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Printf("There are %d most Active Cookies on date %s \n", len(mostOccurrence), inputData.date)
	for i := range mostOccurrence{
		fmt.Println(i+1, " ", mostOccurrence[i])
	}

}
