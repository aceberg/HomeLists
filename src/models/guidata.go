package models

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
