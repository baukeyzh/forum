package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

type commentSQL struct {
	db *sql.DB
}

// New create new database.
func NewCommentSQL(db *sql.DB) *commentSQL {
	return &commentSQL{db: db}
}

// INSERT INTO comments (user_id, date, post_id , content) values (1, "2023-05-01 13:35:04.556898354+06:00" , 1, "golang top, js for girls");
// CreateComment
func (r *commentSQL) CreateComment(comment models.Comment) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (user_id, date, post_id, content) values (?,?,?,?) RETURNING id`, commentsTable)

	row := r.db.QueryRow(query, comment.AuthorID, comment.Date, comment.PostID, comment.Content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// SELECT c.* , u.username as user_name,(SELECT Count (*) FROM comments_likes cl WHERE cl.comment_id = c.id and type = true) as likes,(SELECT Count (*) FROM comments_likes cl WHERE cl.comment_id = c.id and type = false) as dislikes FROM comments c LEFT JOIN users u ON u.id = c.user_id WHERE post_id=7;
// GetCommentsByPostId
func (r *commentSQL) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	var allComments []models.Comment

	query := fmt.Sprintf(`
	SELECT c.* , u.username as user_name,
		(SELECT Count (*) FROM %s cl WHERE cl.comment_id = c.id and type = true) as likes,
		(SELECT Count (*) FROM %s cl WHERE cl.comment_id = c.id and type = false) as dislikes
	FROM %s c 
	LEFT JOIN %s u ON u.id = c.user_id
	WHERE post_id=?;`, commentsLikesTable, commentsLikesTable, commentsTable, usersTable)
	rows, err := r.db.Query(query, postId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("no posts found: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("can't get posts: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		if err = rows.Scan(
			&comment.Id,
			&comment.AuthorID,
			&comment.Date,
			&comment.PostID,
			&comment.Content,
			&comment.AuthorName,
			&comment.Likes,
			&comment.Dislikes,
		); err != nil {
			return nil, fmt.Errorf("can't scan all comments: %w", err)
		}
		allComments = append(allComments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("can't get all comments: %w", err)
	}
	return allComments, nil
}
