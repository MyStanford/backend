package person

type Person struct {
	Name        string `json:"name"`
	Prompt      string `json:"prompt"`
	Description string `json:"description"`
	Emotion     int    `json:"emotion"`
}
