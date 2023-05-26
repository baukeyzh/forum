package service

import (
	"errors"
	"fmt"
	"time"

	"forum/models"
	"forum/repository"
)

// GetAllPosts from posts and likes tables
func GetAllPosts(repos *repository.Repository, currentUserId int) ([]models.Post, error) {
	allPosts, err := repos.Posts.GetAllPosts(currentUserId)
	if err != nil {
		fmt.Println(err.Error())
		return allPosts, errors.New("can't get all posts")

	}
	return allPosts, nil
}

// GetPostById from posts, comments and likes tables
func GetPostById(repos *repository.Repository, id int) (models.Post, error) {
	post, err := repos.Posts.GetPostById(id)
	if err != nil {
		fmt.Println(err.Error())
		return post, errors.New("can't get post by id")

	}
	return post, nil
}

// AddPost to posts table
func AddPost(repos *repository.Repository, post models.Post, categories []int) (int, error) {
	post.Date = time.Now()
	id, err := repos.Posts.CreatePost(post)
	if err != nil {
		return 0, fmt.Errorf("DB can't add post: %w", err)
	}
	for _, catId := range categories {
		if err := repos.Posts.AddCategoryToPost(id, catId); err != nil {
			return 0, fmt.Errorf("DB can't add category: %w", err)
		}
	}
	return id, nil
}
