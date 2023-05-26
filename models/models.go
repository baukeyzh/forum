package models

import (
	"errors"
	"time"
)

type User struct {
	Id       int       `json:"id" db:"id"`
	UserName string    `json:"userName"  db:"username"`
	Email    string    `json:"email"  db:"email"`
	PassHash string    `json:"password_hash"  db:"password_hash"`
	Token    string    `json:"token"  db:"token"`
	ExpireAt time.Time `json:"expire_at"  db:"expire_at"`
}

type Post struct {
	Id         int       `json:"id"  db:"id"`
	AuthorID   int       `json:"authorId"  db:"user_id"`
	AuthorName string    `json:"authorName"  db:"user_name"`
	Date       time.Time `json:"date"  db:"date"`
	Title      string    `json:"title"  db:"title"`
	Content    string    `json:"content"  db:"content"`
	Categories string    `json:"categories"  db:"categories"`
	Likes      int       `json:"likes"  db:"likes"`
	Dislikes   int       `json:"dislikes"  db:"dislikes"`
	MyLikeId   int       `json:"myLikeId"  db:"my_like_id"`
}

type Comment struct {
	Id         int       `json:"id"  db:"id"`
	AuthorID   int       `json:"authorId"  db:"user_id"`
	AuthorName string    `json:"authorName"  db:"user_name"`
	Date       time.Time `json:"date"  db:"date"`
	PostID     string    `json:"postId"  db:"post_id"`
	Content    string    `json:"content"  db:"content"`
	Likes      int       `json:"likes"  db:"likes"`
	Dislikes   int       `json:"dislikes"  db:"dislikes"`
}

type PostAndComments struct {
	Post_info Post      `json:"post_info"`
	Comments  []Comment `json:"comments"`
	IsAuth    bool      `json:"autorized"`
	UserId    int       `json:"userId"`
}

type LikePost struct {
	Id       int    `json:"id"  db:"id"`
	AuthorID int    `json:"authorId"  db:"author_id"`
	PostID   string `json:"postId"  db:"post_id"`
	Type     bool   `json:"type"  db:"type"`
}

type LikeComment struct {
	Id        int    `json:"id"  db:"id"`
	AuthorID  int    `json:"authorId"  db:"author_id"`
	CommentID string `json:"commentId"  db:"comment_id"`
	Type      bool   `json:"type"  db:"type"`
}

type Categories struct {
	Id   int    `json:"id"  db:"id"`
	Name string `json:"authorId"  db:"name"`
}

type RegistrationPage struct {
	SuccessMessage string `json:"successMessage"`
	ErrorMessage   string `json:"errorMessage"`
}

var ErrorUnauthorized = errors.New("Unauthorized")
