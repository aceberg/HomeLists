package models

type Conf struct {
	DbPath   string
	GuiIP    string
	GuiPort  string
	Theme    string
	Color    string
	NodePath string
}

type Item struct {
	ID    int
	Date  string
	Name  string
	Color string
	Count int
	Place string
	Sort  int
}

type Table struct {
	ID    string
	Name  string
	Date  string
	Lines int
}

type WatchItem struct {
	ID      int
	Table   string
	ItemID  int
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
