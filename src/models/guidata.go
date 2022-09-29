package models

type GuiData struct {
	Config 		 Conf
	CurrentTable string
	TableList 	 []Table
	ItemList 	 []Item
	OneItem		 Item
	Themes 		 []string
}