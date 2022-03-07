package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type segment struct {
	ID       string `csv:"ID"`
	ParentID string `csv:"Parent_ID"`
	Name     string `csv:"Name"`
	Provider string `csv:"Provider_ID"`
}

func main() {
	taxonomy, err := os.OpenFile("taxonomy.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer taxonomy.Close()

	var segments []segment

	if err := gocsv.UnmarshalFile(taxonomy, &segments); err != nil { // Load clients from file
		panic(err)

	}

	//for _, r := range segments {
	//
	//	}
	//}

	printTree(segments, "null", 0)

}

func printTree(tbl []segment, parent string, depth int) {
	for _, r := range tbl {
		if r.ParentID == parent {
			for i := 0; i < depth; i++ {
				fmt.Print("--")
			}
			fmt.Print(r.Name, " (ID: "+r.ID, " Parent ID: "+r.ParentID+")", "\n")
			printTree(tbl, r.ID, depth+1)
		}
	}
}
