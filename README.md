# Golang Social Media API
make golang API using GIN, and validator

## How to use this
Please create database go-social-media, or if u not it will error. For table, it will auto add(auto migrate) automatic.

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
* List of All Explore
    * POST http://localhost:3000/v1/post/
    * GET BY NOT USER, AND ORDER BY LIKE http://localhost:3000/v1/explore/:user
    * GET BY EMAIL AND ORDER BY CreateAt http://localhost:3000/v1/explore/user/:email
* List of All Reels
    * POST http://localhost:3000/v1/reels/
    * GET BY NOT USER, AND ORDER BY LIKE http://localhost:3000/v1/reels/e/:user


## How to POST, and UPDATE USER
    {
          "id": 1,
          "name_user": "achmadrizkin",
          "email_user": "achmadrizki@gmail.com",
          "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          "following": 21,
          "followers": 2324
    }

## Response List Of User
    {
        "data": [
            {
                "id": 1,
                "name_user": "achmadrizkin",
                "email_user": "achmadrizki@gmail.com",
                "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
                "following": 21,
                "followers": 2324
            },
            {
                "id": 2,
                "name_user": "cristianoronaldoo",
                "email_user": "cristianoronaldo@gmail.com",
                "image_url": "https://akcdn.detik.net.id/community/media/visual/2021/11/25/cristiano-ronaldo-1.jpeg?w=700&q=90",
                "following": 132,
                "followers": 13231234
            }
        ]
    }

## How to post Explore
    {
        "id": 1,
        "name_user": "achmadrizkin",
        "email_user": "achmadrizki@gmail.com",
        "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
        "image_post": "https://ebooks.gramedia.com/ebook-covers/50217/image_highres/ID_KSH2019MTH12SH.jpg",
        "description_post": "New Book Hehe",
        "Like_post": 12345,
    }

## Response Explore
    {
        "data": [
            {
                "id": 3,
                "name_user": "achmadrizkin",
                "email_user": "achmadrizki22@gmail.com",
                "image_post": "https://ebooks.gramedia.com/ebook-covers/50217/image_highres/ID_KSH2019MTH12SH.jpg",
                "description_post": "New Book Hehe",
                "like_post": 123132213123,
                "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
                "following": 21,
                "followers": 2324,
                "create_at": "2021-12-03T19:06:34.202+07:00",
                "update_at": "0001-01-01T00:00:00Z"
            },
            {
                "id": 4,
                "name_user": "achmadrizkin",
                "email_user": "achmadrizki22@gmail.com",
                "image_post": "https://ebooks.gramedia.com/ebook-covers/50217/image_highres/ID_KSH2019MTH12SH.jpg",
                "description_post": "New Book Hehe",
                "like_post": 832230,
                "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
                "following": 21,
                "followers": 2324,
                "create_at": "2021-12-03T19:10:44.478+07:00",
                "update_at": "0001-01-01T00:00:00Z"
            },
            {
                "id": 1,
                "name_user": "achmadrizkin",
                "email_user": "achmadrizki@gmail.com",
                "image_post": "https://ebooks.gramedia.com/ebook-covers/50217/image_highres/ID_KSH2019MTH12SH.jpg",
                "description_post": "New Book Hehe",
                "like_post": 12345,
                "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
                "following": 21,
                "followers": 2324,
                "create_at": "2021-12-03T17:39:04.172+07:00",
                "update_at": "0001-01-01T00:00:00Z"
            },
            {
                "id": 2,
                "name_user": "achmadrizkin",
                "email_user": "achmadrizki22@gmail.com",
                "image_post": "https://ebooks.gramedia.com/ebook-covers/50217/image_highres/ID_KSH2019MTH12SH.jpg",
                "description_post": "New Book Hehe",
                "like_post": 10,
                "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
                "following": 21,
                "followers": 2324,
                "create_at": "2021-12-03T19:05:32.003+07:00",
                "update_at": "0001-01-01T00:00:00Z"
            }
        ]
    }

## POST OR UPDATE REELS
    {
        "id": 1,
        "name_user": "achmadrizkin",
        "email_user": "achmadrizki22@gmail.com",
        "image_url": "https://images.unsplash.com/photo-1584462746497-276f4aeb9fca?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
        "video_post": "https://assets.mixkit.co/videos/preview/mixkit-man-under-multicolored-lights-1237-large.mp4",
        "description_post": "A man with a small beard and mustache wearing a white sweater, sunglasses, and a backwards black baseball cap turns his head in different directions under changing colored lights",
        "Like_post": 12345,
        "following": 21,
        "followers": 2324
    }