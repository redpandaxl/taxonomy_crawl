package modules

import (
	"fmt"
	"taxonomy_crawl/models"
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
