package field

type Field struct {
	ID      int    `json:"id"`
	PlaceID int    `json:"place_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}
