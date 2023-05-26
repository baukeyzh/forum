package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable             = "users"
	postsTable             = "posts"
	categoriesTable        = "categories"
	categoriesToPostsTable = "posts_categories"
	commentsTable          = "comments"
	postsLikesTable        = "posts_likes"
	commentsLikesTable     = "comments_likes"
)

type Storage struct {
	Db *sql.DB
}

type Authorization interface {
	CreateUser(User models.User) (int, error)
	GetUser(Email string) (models.User, error)
	AddToken(User models.User) error
	GetUserByToken(Token string) (models.User, error)
	DeleteToken(User models.User) error
}

type Posts interface {
	CreatePost(post models.Post) (int, error)
	GetAllPosts(currentUserId int) ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	AddCategoryToPost(postId, catId int) error
}

type Comments interface {
	CreateComment(comment models.Comment) (int, error)
	GetCommentsByPostId(postId int) ([]models.Comment, error)
}
type Likes interface {
	AddLikePost(like models.LikePost) (int, error)
	AddLikeComment(like models.LikeComment) (int, error)
	GetLikeByPostUser(postId, userId int) (models.LikePost, error)
	GetLikeByCommentUser(commentId, userId int) (models.LikeComment, error)
}

type Repository struct {
	Authorization
	Posts
	Comments
	Likes
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}

	return &Storage{Db: db}, nil
}

// Init all
func (s *Storage) Init(initSqlFileName string) error {
	file, err := ioutil.ReadFile(initSqlFileName)
	if err != nil {
		log.Fatalf("Can't read SQL file %v", err)
	}

	// Execute all
	_, err = s.Db.Exec(string(file))
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	return nil
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		Posts:         NewPostsSQL(db),
		Comments:      NewCommentSQL(db),
		Likes:         NewlikeSQL(db),
	}
}
