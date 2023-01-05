package place

type Place struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
}
