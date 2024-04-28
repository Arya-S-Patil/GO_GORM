# GO_GORM
STILL WORK IN PROGRESS (i have learned to set up both GORM and gin web frame work and make all crud APIs like the other repository https://github.com/Arya-S-Patil/go_crud_social now am trying to connect database)

---

# Go CRUD API with GORM and Gin

This is a work-in-progress project for building a CRUD API using the GORM ORM library and the Gin web framework in Go. The project aims to connect to a PostgreSQL database and perform CRUD operations on it.

## Initial Setup

1. Clone this repository to your local machine:

    ```bash
    git clone <repository-url>
    ```

2. Install the required dependencies:

    ```bash
    go mod tidy
    ```

3. Set up your PostgreSQL database. You can use tools like pgAdmin or the PostgreSQL CLI to create a new database.

## Database Connection

The next step is to establish a connection to your PostgreSQL database. Follow these steps to set up the database connection:

1. Open the `.env.example` file and provide your PostgreSQL database details:

    ```bash
    DB_HOST=<database-host>
    DB_PORT=<database-port>
    DB_USER=<database-user>
    DB_PASSWORD=<database-password>
    DB_NAME=<database-name>
    ```

2. Rename the `.env.example` file to `.env`.

3. Once the database connection details are provided in the `.env` file, the application will be able to connect to the database.

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The API server will start running on `http://localhost:3000`.

## Routes

The following routes are available in the API:

- `GET /api/users`: Get all users.
- `GET /api/users/:id`: Get a single user by ID.
- `POST /api/users`: Create a new user.
- `PUT /api/users/:id`: Update an existing user by ID.
- `DELETE /api/users/:id`: Delete a user by ID.

---

Feel free to expand on this README file as you progress with your project. You can add more details about the project structure, middleware usage, error handling, authentication, etc., as needed. Good luck with your project!
