package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gonutz/wui/v2"
)

/*
	Author: Amir Montgomery
	Date: 3/7/2022
	Obj: Read two CSVs and create a slice with all of the differences and line number
*/
type Row struct {
	ID    string
	Index int
}

type CSV struct {
	Rows []Row // slice of rows
}

func main() {

	/*
		1. Load .csv
		2. Create CSV struct from loaded .csv
		3. Read .csv line by line into Row struct
		4. Append Row structs into CSV.Row

		https://gosamples.dev/read-csv/

		Ideas to improve:
		wrap os.File in bufio.Reader
		wrap file handle in 20kb buffer
	*/

	// create dialogs and windows
	fileOpenDialog := wui.NewFileOpenDialog()
	window := wui.NewWindow()

	killMe := false
	window.SetOnClose(func() {

		killMe = true
		os.Exit(1)
	})

	var sliceFilePaths []string

	for len(sliceFilePaths) != 2 {
		_, sliceFilePaths = fileOpenDialog.ExecuteMultiSelection(window) // returns []string

		// if the box wasn't cancelled lmfao then check this
		if killMe == true {
			break
		} else {
			if len(sliceFilePaths) != 2 {
				fmt.Println(killMe)
				wui.MessageBoxError("Error", "This application wil only take 2 files. Please try again.")
			}
		}
	}

	fileOne := createNewCSVStruct(sliceFilePaths[0])
	fileTwo := createNewCSVStruct(sliceFilePaths[1])

	difference(fileOne, fileTwo)
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b *CSV) []string {
	mb := make(map[string]struct{}, len(b.Rows))
	for _, x := range b.Rows {
		mb[x.ID] = struct{}{}
	}
	var diff []string
	for _, x := range a.Rows {
		if _, found := mb[x.ID]; !found {
			tstring := fmt.Sprintf("Row %d: (%s)", x.Index, x.ID)
			diff = append(diff, tstring)
			//fmt.Println(x.ID)
			fmt.Println(tstring)
		}
	}
	return diff
}

/*
	[Map Notes] https://gobyexample.com/maps

	make(map[key type]value type)

	var file *CSV = { *Row, *Row, *Row }
	var file[x] *Row = [ID, Index]
	file[x].ID == string
	file[x].Index == int
*/

func createNewCSVStruct(file string) *CSV {
	f, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
	}

	// close file
	defer f.Close()

	var fullFile *CSV = new(CSV)
	i := 1
	var tempRow *Row = new(Row)

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			// fmt.Println("No more data.")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		// rec == [ADsPath sAMAccountName displayName mail]
		// split into slice, fill out Row struct, append to fullFile.Row
		//fmt.Println(rec)
		tempRow.ID = rec[0]
		tempRow.Index = i

		fullFile.Rows = append(fullFile.Rows, *tempRow)
		i++
	}

	return fullFile
}
