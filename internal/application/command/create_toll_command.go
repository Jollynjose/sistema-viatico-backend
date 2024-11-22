package command

type CreateTollCommand struct {
	Price float64 `json:"price"`
	Order int     `json:"order"`
}
