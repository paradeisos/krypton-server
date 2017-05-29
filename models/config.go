package models

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Mode     string `json:"mode"`
	Pool     int    `json:"pool"`
	Timeout  int    `json:"timeout"`
}
