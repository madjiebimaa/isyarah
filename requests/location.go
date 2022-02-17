package requests

type LocationCreate struct {
	Location    string `json:"location" binding:"require"`
	Description string `json:"description" binding:"require,min=5"`
}
