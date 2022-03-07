package modules

import (
	"fmt"
	"github.com/redpandaxl/taxonomy_crawl/models"
	"log"
)

func FilterByProvider(as []models.Segment, ss []models.Segment) (out []models.Segment) {
	f := make(map[string]struct{}, len(ss))
	for _, u := range as {
		f[u.Provider] = struct{}{}
	}
	for _, u := range ss {
		if _, ok := f[u.Provider]; ok {
			out = append(out, u)
			PrintTree(out, "null", 0)
		}
	}
	return
}

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

func FilterProvider(segments []models.Segment) {
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

	TestPrintTree(f, "null", 0)
}

func TestPrintTree(tbl map[string][]models.Segment, parent string, depth int) {
	for k, r := range tbl {
		log.Println("k: ", k)
		log.Println("r: ", r)
	}
	//PrintTree(tbl, "null", 0)
}
