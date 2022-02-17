package models

type Geometry struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Query struct {
	Results []struct {
		Bounds struct {
			Northeast Geometry `json:"northeast"`
			Southwest Geometry `json:"southwest"`
		} `json:"bounds"`
		Components struct {
			City        string `json:"city"`
			Continent   string `json:"continent"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			PostCode    string `json:"postcode"`
			State       string `json:"state"`
			StateCode   string `json:"state_code"`
		} `json:"components"`
		Confidence int      `json:"confidence"`
		Formatted  string   `json:"formatted"`
		Geometry   Geometry `json:"geometry"`
	} `json:"results"`
}
