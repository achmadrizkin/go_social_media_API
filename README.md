# Golang Social Media API
make golang API using GIN, and validator

## How to use this
Please create database go-ecommerce, or if u not it will error. For table, it will auto add(auto migrate) automatic.

And after that, download this repo, and copy this text, and run in terminal. and its done.

    go run main.go

## Implementation in Android (Coming Soon ...)
...

## Api Spec
please change the localhost:3000 to 10.0.2.2:3000 if u want to use emulator

* List Of User
    * POST http://localhost:3000/v1/users/
    * GET ALL USERS http://localhost:3000/v1/users/
    * Search User http://localhost:3000/v1/users/:user



## How to POST
    {
          "id": 1,
          "name_user": "achmadrizkin",
          "email_user": "achmadrizki@gmail.com",
          "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          "following": 21,
          "followers": 2324
    }