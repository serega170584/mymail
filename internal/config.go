package internal

type Config struct {
	App struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Mail struct {
		Port     int    `json:"port"`
		From     string `json:"from"`
		To       string `json:"to"`
		Host     string `json:"host"`
		Password string `json:"password"`
	}
}
