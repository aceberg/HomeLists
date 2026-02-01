package models

type Conf struct {
	DbPath  string
	GuiIP   string
	GuiPort string
	Theme   string
}

type Item struct {
	Id    uint16
	Date  string
	Name  string
	Color string
	Count uint16
	Place string
	Sort  uint16
}

type Table struct {
	Id    string
	Name  string
	Date  string
	Lines uint16
}

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

type LineColor struct {
	Name string
	Code string
}

type GuiData struct {
	Config       Conf
	CurrentTable string
	TableList    []Table
	ItemList     []Item
	WatchList    []WatchItem
	OneItem      Item
	Themes       []string
	Colors       []LineColor
}
