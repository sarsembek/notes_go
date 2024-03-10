package model

type Pokemon struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Type1       string  `json:"type1"`
	Type2       string  `json:"type2"`
	Height      float64 `json:"height"`
	Weight      float64 `json:"weight"`
	BaseExp     int     `json:"base_experience"`
	CaptureRate int     `json:"capture_rate"`
	HP          int     `json:"hp"`
	Attack      int     `json:"attack"`
	Defense     int     `json:"defense"`
	SpAttack    int     `json:"special_attack"`
	SpDefense   int     `json:"special_defense"`
	Speed       int     `json:"speed"`
}
