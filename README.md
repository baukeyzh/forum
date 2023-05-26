# FORUM
## Objectives

This project consists in creating a web forum that allows :

    communication between users.
    associating categories to posts.
    liking and disliking posts and comments.
    filtering posts.


## Instructions for user registration:

    Must ask for email
        When the email is already taken return an error response.
    Must ask for username
    Must ask for password
        The password must be encrypted when stored (this is a Bonus task)

## Requirements

    You must use SQLite.
    You must handle website errors, HTTP status.
    You must handle all sort of technical errors.
    The code must respect the good practices.
    It is recommended to have test files for unit testing.


## How to use:
Create docker image
` make build
`
Run docker image
` make docker-run
`
- Go to the 'http://localhost:8082/' and go ahead)

For stop container
` make docker-stop
`
For delete container and images
` make docker-delete
`

## Authors
- [@AliAkhmet](https://01.alem.school/git/AliAkhmet)
- [@zbaukey](https://01.alem.school/git/zbaukey)