
# Go CRUD API - Movie Management

This is a simple **CRUD (Create, Read, Update, Delete)** API for managing movies, built with **Go (Golang)**. The project utilizes the **gorilla/mux** router to handle routing and allows performing various operations on a list of movies. Each movie has a title, an ISBN, and a director.

## Project Overview

This project demonstrates a basic RESTful API that allows users to manage movies. The API supports the following operations:

- **Get all movies**
- **Get a movie by its ID**
- **Create a new movie**
- **Update an existing movie**
- **Delete a movie**

## Prerequisites

- [Go 1.16+](https://golang.org/dl/) installed on your local machine.
- A working Go environment setup.

The project uses the `gorilla/mux` package for routing. Install it by running:

```bash
go get github.com/gorilla/mux
```

## Setup and Installation

1. Clone the repository:

```bash
git clone https://github.com/Moozaheed/Golang-Learning-Projects.git
cd Golang-Learning-Projects/go-curd
```

2. Install dependencies:

Install the necessary Go package (gorilla/mux):

```bash
go get github.com/gorilla/mux
```

3. Run the server:

To start the API server, run:

```bash
go run main.go
```

The server will start and listen on port 8000.

## API Endpoints

### `GET /movies`

Fetch all movies.

#### Response Example:
```json
[
  {
    "id": "1",
    "isbn": "448743",
    "title": "Movie One",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  },
  {
    "id": "2",
    "isbn": "448744",
    "title": "Movie Two",
    "director": {
      "firstname": "Steve",
      "lastname": "Smith"
    }
  }
]
```

---

### `GET /movies/{id}`

Fetch a specific movie by its ID.

#### URL Params:
- `id` (string): The ID of the movie.

#### Response Example:
```json
{
  "id": "1",
  "isbn": "448743",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

---

### `POST /movies`

Create a new movie.

#### Request Body:
```json
{
  "isbn": "448749",
  "title": "Movie Five",
  "director": {
    "firstname": "Mark",
    "lastname": "Taylor"
  }
}
```

#### Response Example:
```json
{
  "id": "5",
  "isbn": "448749",
  "title": "Movie Five",
  "director": {
    "firstname": "Mark",
    "lastname": "Taylor"
  }
}
```

---

### `PUT /movies/{id}`

Update an existing movie.

#### URL Params:
- `id` (string): The ID of the movie.

#### Request Body:
```json
{
  "isbn": "448749",
  "title": "Updated Movie Five",
  "director": {
    "firstname": "Mark",
    "lastname": "Taylor"
  }
}
```

#### Response Example:
```json
{
  "id": "5",
  "isbn": "448749",
  "title": "Updated Movie Five",
  "director": {
    "firstname": "Mark",
    "lastname": "Taylor"
  }
}
```

---

### `DELETE /movies/{id}`

Delete a movie by its ID.

#### URL Params:
- `id` (string): The ID of the movie.

#### Response Example:
```json
[
  {
    "id": "1",
    "isbn": "448743",
    "title": "Movie One",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  },
  {
    "id": "2",
    "isbn": "448744",
    "title": "Movie Two",
    "director": {
      "firstname": "Steve",
      "lastname": "Smith"
    }
  }
]
```