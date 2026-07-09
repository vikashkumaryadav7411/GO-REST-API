
A robust, production-ready RESTful API for managing notes built with **Go (Golang)**, the **Gin Web Framework**, and **MongoDB**. 

The project follows a clean, decoupled architecture breaking features down into configuration, routing, request handling, and database repositories.

---

## Features

* **Complete CRUD operations** for managing persistent notes.
* **Environment variable configuration** utilizing `.env` files.
* **Structured routing and middleware** using the Gin framework.
* **MongoDB driver integration** complete with strict timeouts and context-safety.
* **Health check endpoint** for container orchestrators or monitoring setups.

---
Prerequisites
Go: Version 1.21 or higher installed.

MongoDB: A running local instance or a MongoDB Atlas connection string.

Installation & Setup
Clone the repository:

Bash
git clone 
cd notes-api
Configure Environment Variables:
Create a .env file in the root directory of the project:

Code snippet
PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=notes_db

Download Go Dependencies:

Bash
go mod download

Run the Application:

Bash
go run cmd/api/main.go


Technologies Used
Language: Go (Golang)

Web Framework: Gin Gonic

Database Client: Official MongoDB Go Driver

Environment Management: godotenv



## Project Structure

```text
├── cmd/
│   └── api/
│       └── main.go          # Application entrypoint
├── internal/
│   ├── config/
│   │   └── config.go        # Environment variable loader
│   ├── db/
│   │   └── mongo.go         # Database connection client lifecycle
│   ├── notes/
│   │   ├── handler.go       # HTTP request/response handling logic
│   │   ├── model.go         # Database and JSON representations
│   │   ├── repository.go    # Direct MongoDB interaction layers
│   │   └── routes.go        # HTTP endpoints group definitions
│   └── server/
│       └── router.go        # Core Engine setup & universal middleware
├── .env.example             # Template for your local environment configs
├── go.mod
└── go.sum




