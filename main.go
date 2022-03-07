package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/*
	Author: Amir Montgomery
	Date: 3/7/2022
	Obj: Read two CSVs and create a slice with all of the differences and line number
*/

type Row struct {
	LDAP        string
	SAM         string
	DisplayName string
	Mail        string
	Index       int
}

type CSV struct {
	Row []Row
}

func main() {

	/*
		1. Load .csv
		2. Create CSV struct from loaded .csv
		3. Read .csv line by line into Row struct
		4. Append Row structs into CSV.Row

		https://gosamples.dev/read-csv/
	*/

	// open file
	f, err := os.Open("C:/Users/amontgomery/OneDrive - rams.nfl.com/3_1TO.csv")
	if err != nil {
		fmt.Print(err)
	}

	// close file
	defer f.Close()

	var fullFile *CSV = new(CSV)
	i := 0
	var tempRow *Row = new(Row)

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("No more data.")
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		// rec == [ADsPath sAMAccountName displayName mail]
		// split into slice, fill out Row struct, append to fullFile.Row
		//fmt.Println(rec)
		tempRow.LDAP = rec[0]
		tempRow.SAM = rec[1]
		tempRow.DisplayName = rec[2]
		tempRow.Mail = rec[3]
		tempRow.Index = i

		fullFile.Row = append(fullFile.Row, *tempRow)
		i++
	}

	fmt.Print(fullFile)
}
