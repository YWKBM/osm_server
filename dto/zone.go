package dto

type Zone struct {
	Name string `json:"name"`
	Geo  string `json:"geo"`
}

type ZoneList struct {
	Items []ZoneListItem `json:"items"`
}

type ZoneListItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
