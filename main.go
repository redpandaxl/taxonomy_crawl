package main

import (
	"github.com/gocarina/gocsv"
	"os"
	"taxonomy_crawl/models"
	"taxonomy_crawl/modules"
)

func main() {
	taxonomy, err := os.OpenFile("taxonomy.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(taxonomy *os.File) {
		err := taxonomy.Close()
		if err != nil {

		}
	}(taxonomy)

	var segments []models.Segment

	if err := gocsv.UnmarshalFile(taxonomy, &segments); err != nil {
		// Load clients from file
		panic(err)

	}

	modules.PrintTree(segments, "null", 0)
	//fmt.Println(modules.FilterByProvider(segments, []models.Segment{}))

}