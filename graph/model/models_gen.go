// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	VideoID int    `json:"video_id"`
	Comment string `json:"comment"`
	Like    int    `json:"like"`
	Dislike int    `json:"dislike"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
	Second  int    `json:"second"`
	Minute  int    `json:"minute"`
	Hour    int    `json:"hour"`
}

type DetailPlaylist struct {
	ID         string `json:"id"`
	PlaylistID int    `json:"playlist_id"`
	VideoID    int    `json:"video_id"`
}

type LikeComment struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CommentID int    `json:"comment_id"`
	Type      string `json:"type"`
}

type LikePost struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PostID int    `json:"post_id"`
	Type   string `json:"type"`
}

type LikeReply struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	ReplyID int    `json:"reply_id"`
	Type    string `json:"type"`
}

type LikeVideo struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	VideoID int    `json:"video_id"`
	Type    string `json:"type"`
}

type Playlist struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Second      int    `json:"second"`
	Minute      int    `json:"minute"`
	Hour        int    `json:"hour"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	Privacy     string `json:"privacy"`
	UserID      string `json:"user_id"`
	Views       int    `json:"views"`
}

type Post struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Date        string `json:"date"`
}

type Reply struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CommentID int    `json:"comment_id"`
	Reply     string `json:"reply"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Second    int    `json:"second"`
	Minute    int    `json:"minute"`
	Hour      int    `json:"hour"`
}

type Subscribe struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	SubscribeTo string `json:"subscribe_to"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	Membership string `json:"membership"`
	Subscriber int    `json:"subscriber"`
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
	Premium     string `json:"premium"`
	Category    string `json:"category"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	Second      int    `json:"second"`
	Minute      int    `json:"minute"`
	Hour        int    `json:"hour"`
	Duration    int    `json:"duration"`
}

type NewComment struct {
	UserID  string `json:"user_id"`
	VideoID int    `json:"video_id"`
	Comment string `json:"comment"`
	Like    int    `json:"like"`
	Dislike int    `json:"dislike"`
}

type NewPlaylist struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
	UserID      string `json:"user_id"`
	Views       int    `json:"views"`
}

type NewPost struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Date        string `json:"date"`
}

type NewReply struct {
	UserID    string `json:"user_id"`
	CommentID int    `json:"comment_id"`
	Reply     string `json:"reply"`
}

type NewUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	Membership string `json:"membership"`
	Subscriber int    `json:"subscriber"`
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
	Premium     string `json:"premium"`
	Category    string `json:"category"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
	Duration    int    `json:"duration"`
}
