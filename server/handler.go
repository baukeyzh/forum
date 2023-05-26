package server

import (
	"html/template"
	"log"
	"net/http"

	"forum/repository"
)

var (
	tpl  *template.Template
	port = ":8081"
)

type Handler struct {
	repos *repository.Repository
}

// NewHandler create Handler struct with repos parameter
func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: repos}
}

// Server func - all handlers
func Server(h *Handler) {
	// Middleware for identification the user by cookie
	member := h.identification

	tpl = template.Must(template.ParseGlob("templates/*.html"))
	//---- POST and GET ----
	// Auth Handlers
	http.HandleFunc("/login", h.gestLogin)
	http.HandleFunc("/registration", h.gestRegistration)

	//---- GET ONLY ----
	// homepage
	http.HandleFunc("/", h.homePage)
	// get all posts
	http.HandleFunc("/posts", h.getAllPosts)
	// gel one past page
	http.HandleFunc("/post-page", h.getPostAndComments)

	//---- POST ONLY ----
	// logout
	http.HandleFunc("/logout", member(h.memberLogout))
	// add post
	http.HandleFunc("/v1/post/create", member(h.memberPostCreate))
	// add comment
	http.HandleFunc("/v1/comment/create", member(h.memberCommentCreate))
	//  add likes
	http.HandleFunc("/v1/post/like", member(h.memberLikeForPost))
	http.HandleFunc("/v1/comment/like", member(h.memberLikeForComment))

	// handle css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//
	log.Printf("Starting a web server on http://localhost%s/ ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
