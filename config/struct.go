package config

type User struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Position string `json:"position"`
}