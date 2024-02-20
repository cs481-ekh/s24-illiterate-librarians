package model

type Child struct {
	ParentId  string `json:"parentId"`
	ChildId   string `json:"childId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Grade     int8   `json:"grade"`
}
