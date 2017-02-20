package main

import "time"

type Response struct {
	Name      string    `json:"name"`
	Flag      bool      `json:"flag"`
	CreatedAt time.Time `json:"createdAt"`
}

type Responses []Response
