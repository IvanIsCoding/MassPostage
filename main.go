package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type CSVLine struct {
	Columns map[string]string
}

func ReadCSV(filename string) ([]CSVLine, error) {

	// Open CSV File
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	// Load CSV into memory
	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	columnNames := lines[0]
	csvContent := []CSVLine{}

	for idx, line := range lines {

		if idx == 0 {
			continue // We skip the header column
		}

		// Map values to column names, and add to the result
		csvLine := CSVLine{}
		csvLine.Columns = make(map[string]string)
		for pos, columnName := range columnNames {
			csvLine.Columns[columnName] = line[pos]

		}

		csvContent = append(csvContent, csvLine)

	}

	return csvContent, nil

}

func ReadEML(filename string) (string, error) {

	text, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(text), nil

}

func WriteEML(filename string, message string, line CSVLine) error {

	// Replace place holders with personalized content
	finalMessage := message

	for key, value := range line.Columns {

		placeHolder := "{" + "{" + key + "}" + "}"
		finalMessage = strings.ReplaceAll(finalMessage, placeHolder, value)

	}

	// Write EML file
	err := ioutil.WriteFile(filename, []byte(finalMessage), 0644) // 0644 is the permission
	return err

}

func main() {

	// Processing CLI arguments
	args := os.Args[1:]

	csvFile := args[0]
	textFile := args[1]
	campaignName := args[2]

	// Reading csvFile and textFile
	csvSource, err := ReadCSV(csvFile)
	if err != nil {
		panic(err)
	}

	textSource, err := ReadEML(textFile)
	if err != nil {
		panic(err)
	}

	// Creating directory for the campaign
	err = os.Mkdir(campaignName, 0755) // 0755 is the directory permission
	if err != nil {
		panic(err)
	}

	// Writing files to campaign
	for i, csvLine := range csvSource {

		// creating path to the new file
		filename := fmt.Sprintf("email_%d.eml", i+1)
		filename = filepath.Join(campaignName, filename)

		err := WriteEML(filename, textSource, csvLine)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s has been created\n", filename)

	}

	fmt.Printf("Campaign %s has been completed\n", campaignName)

}
