package command

type CreateJobPositionHistoryCommand struct {
	Lunch         float64 `json:"lunch"`
	BreakFast     float64 `json:"breakfast"`
	Dinner        float64 `json:"dinner"`
	Accommodation float64 `json:"accommodation"`
}
