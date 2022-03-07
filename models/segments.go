package models

type Segment struct {
	ID       string `csv:"ID"`
	ParentID string `csv:"Parent_ID"`
	Name     string `csv:"Name"`
	Provider string `csv:"Provider_ID"`
}

type Segments struct {
	segment Segment
}
