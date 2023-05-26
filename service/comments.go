package service

import (
	"errors"
	"fmt"
	"time"

	"forum/models"
	"forum/repository"
)

// GetAllComments from comments and likes tables
func GetAllComments(repos *repository.Repository, postId int) ([]models.Comment, error) {
	comments, err := repos.Comments.GetCommentsByPostId(postId)
	if err != nil {
		fmt.Println(err.Error())
		return comments, errors.New("can't get all comments")

	}
	return comments, nil
}

// AddComment to comments table
func AddComment(repos *repository.Repository, comm models.Comment) (int, error) {
	comm.Date = time.Now()
	id, err := repos.Comments.CreateComment(comm)
	if err != nil {
		return 0, fmt.Errorf("DB can't add token: %w", err)
	}
	return id, nil
}
