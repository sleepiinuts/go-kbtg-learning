package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type AgriHist struct {
	Year      int
	IndusCD   string
	IndusName string
	Size      string
	Variable  string
	Value     int
	Unit      string
}

func readCSV() {
	fr, err := os.Open("sampleCSV.csv")

	if err != nil {
		fmt.Printf("error open for read csv file: %v\n", err)
		return
	}

	defer fr.Close()

	reader := csv.NewReader(fr)

	fw, err := os.OpenFile("sample-out.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("error open for write: %v\n", err)
		return
	}

	defer fw.Close()

	writer := csv.NewWriter(fw)

	var agh AgriHist
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("error reading csv: %v\n", err)
			return
		}

		// read to struct
		toStruct(record, &agh)
		fmt.Println("read agh: ", agh)

		err = writer.Write(record)
		if err != nil {
			fmt.Printf("error writing csv: %v\n", err)
			return
		}

		writer.Flush()
	}
}

func toStruct(str []string, agh *AgriHist) error {

	var err error
	for i, s := range str {
		switch i {
		case 0:
			// Year
			agh.Year, err = strconv.Atoi(s)
			if err != nil {
				return err
			}
		case 1:
			// indust code
			agh.IndusCD = s
		case 2:
			// indust name
			agh.IndusName = s
		case 3:
			// indust size
			agh.Size = s
		case 5:
			agh.Value, err = strconv.Atoi(s)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
