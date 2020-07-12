// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Membership bool   `json:"membership"`
}

type Video struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	URL    string `json:"url"`
	Watch  int    `json:"watch"`
	Like   int    `json:"like"`
}

type NewUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Membership bool   `json:"membership"`
}

type NewVideo struct {
	UserID string `json:"user_id"`
	URL    string `json:"url"`
	Watch  int    `json:"watch"`
	Like   int    `json:"like"`
}
