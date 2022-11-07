package models

type WatchItem struct {
	Id      int
	Table   string
	ItemId  int
	Name    string
	ByDate  string
	Date    string
	ByCount string
	Count   int
}
