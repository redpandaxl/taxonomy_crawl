package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type segment struct {
	ID       string `csv:"ID"`
	ParentID string `csv:"Parent_ID"`
	Name     string `csv:"Name"`
}

func main() {
	taxonomy, err := os.OpenFile("taxonomy.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer taxonomy.Close()

	segments := []*segment{}

	if err := gocsv.UnmarshalFile(taxonomy, &segments); err != nil { // Load clients from file
		panic(err)
	}

	for _, segment := range segments {

		fmt.Print(segment.Name, "\n\n")

	}

}

// ReadCsv accepts a file and returns its content as a multi-dimension type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

//func printTree(tbl, parent string, depth int) {
//	for _, r := range tbl {
//		if r.ParentID == parent {
//			for i := 0; i < depth; i++ {
//				fmt.Print("--")
//			}
//			fmt.Print(r.Name, "\n\n")
//			printTree(tbl, r.ID, depth+1)
//		}
//	}
//}

//func printTree(tbl, parent string, depth int) {
//	if tbl.ParentID == "null" {
//		for i := 0; i < depth; i++ {
//			fmt.Print("--")
//		}
//		fmt.Print(tbl.Name, "\n\n")
//		//printTree(segment, segment.ID, depth+1)
//	}
