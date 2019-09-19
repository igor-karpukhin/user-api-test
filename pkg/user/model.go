package user

import "time"

type User struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Age      uint      `json:"age"`
	Hobbies  []string  `json:hobbies`
}
