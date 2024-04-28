# GO_GORM
STILL WORK IN PROGRESS (i have learned to set up both GORM and gin web frame work and make all crud APIs like the other repository https://github.com/Arya-S-Patil/go_crud_social now am trying to connect database)

---
markdown
Copy code
# Go CRUD API with GORM and Gin

![Go version](https://img.shields.io/badge/Go-v1.16-blue)
![Gin version](https://img.shields.io/badge/Gin-v1.7.4-green)
![GORM version](https://img.shields.io/badge/GORM-v1.21.7-orange)

## Work in Progress

This is a work-in-progress project aimed at building a CRUD API using the GORM ORM library and the Gin web framework in Go. The project's objective is to connect to a PostgreSQL database and perform CRUD operations on it.

## Initial Setup

 **Clone the Repository:**

   ```bash
   git clone <repository-url>
Install Dependencies:
bash
Copy code
go mod tidy
Set up PostgreSQL Database:Create a new PostgreSQL database using tools like pgAdmin or the PostgreSQL CLI.
##Database Connection
Provide Database Details:Open the .env.example file and provide your PostgreSQL database details:
env
Copy code
DB_HOST=<database-host>
DB_PORT=<database-port>
DB_USER=<database-user>
DB_PASSWORD=<database-password>
DB_NAME=<database-name>
Rename Environment File:Rename the .env.example file to .env.
Run the Application:
bash
Copy code
go run main.go
The Gin router runs on port 8080, and the Gorilla Mux router runs on port 8000. Both routers handle CRUD operations for the Post resource.
Routes
##The following routes are available in the API:

GET /api/users: Get all users.
GET /api/users/:id: Get a single user by ID.
POST /api/users: Create a new user.
PUT /api/users/:id: Update an existing user by ID.
DELETE /api/users/:id: Delete a user by ID.
