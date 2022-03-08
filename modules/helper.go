package modules

import (
	"fmt"
	"github.com/redpandaxl/taxonomy_crawl/models"
)

func PrintTree(tbl []models.Segment, parent string, depth int) {
	for _, r := range tbl {
		if r.ParentID == parent {
			for i := 0; i < depth; i++ {
				fmt.Print("--")
			}
			fmt.Print(r.Name, " (Provider ID: "+r.Provider, ")\n")
			PrintTree(tbl, r.ID, depth+1)
		}
	}
}

func FilterProvider(segments []models.Segment) map[string][]models.Segment {
	f := make(map[string][]models.Segment)
	for _, seg := range segments {
		segArray, ok := f[seg.Provider]
		if ok {
			//segArray is already valid
		} else {
			segArray = make([]models.Segment, 0)
		}
		segArray = append(segArray, seg)
		f[seg.Provider] = segArray
	}
	for k, r := range f {
		fmt.Println("-----Provider ID: ", k, "-----")
		PrintTree(r, "null", 0)

	}
	return f
}
