package cmd

type Job struct {
	Id          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Number      uint64  `json:"number"`
	Category    string  `json:"category"`
	Rate        float64 `json:"rate"`
}
