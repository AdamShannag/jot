package model

type Project struct {
	Name     string    `json:"name"`
	Services []Service `json:"services"`
}
