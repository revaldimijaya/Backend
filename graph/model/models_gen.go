// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	VideoID int    `json:"video_id"`
	Comment string `json:"comment"`
	Like    int    `json:"like"`
	Dislike int    `json:"dislike"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	Membership string `json:"membership"`
}

type Video struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	URL         string `json:"url"`
	Watch       int    `json:"watch"`
	Like        int    `json:"like"`
	Dislike     int    `json:"dislike"`
	Restriction string `json:"restriction"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Thumnail    string `json:"thumnail"`
}

type NewComment struct {
	UserID  string `json:"user_id"`
	VideoID int    `json:"video_id"`
	Comment string `json:"comment"`
	Like    int    `json:"like"`
	Dislike int    `json:"dislike"`
}

type NewUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	Membership string `json:"membership"`
}

type NewVideo struct {
	UserID      string `json:"user_id"`
	URL         string `json:"url"`
	Watch       int    `json:"watch"`
	Like        int    `json:"like"`
	Dislike     int    `json:"dislike"`
	Restriction string `json:"restriction"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Thumnail    string `json:"thumnail"`
}
