## Project Title: GoEchoAuth

GoEchoAuth is a secure authentication API built with Go and Echo framework, utilizing PostgreSQL as the database. The primary purpose of this project is to demonstrate a simple user registration and login system with a focus on security practices. The application incorporates GORM for seamless interaction with the database and offers a foundation for developers to understand and extend user authentication features in their own Go projects.

### Key Features

- User registration with secure password storage.
- User login with authentication and authorization.
- Utilizes Docker for easy setup and deployment.
- Demonstrates best practices for Go web development.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Endpoints](#endpoints)

Feel free to explore the repository to understand how to implement secure user authentication in Go applications.

## Introduction

Briefly introduce your project. What does it do, and why is it useful? Include any relevant background information.

## Getting Started

Provide instructions on how to get the project up and running on a local machine. Include information on prerequisites and installation steps.

### Prerequisites

List any software, libraries, or services that need to be installed before running the project.

- Docker
- Go

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/wussh/user-management-backend.git
   ```

2. Navigate to the project directory:

   ```bash
   cd user-management-backend
   ```

3. Build and run the Docker containers:

   ```bash
   docker-compose up --build
   ```

   This will initialize the PostgreSQL database and start the application.

### Usage

To interact with the application, you can use the provided API endpoints. Below are examples of how to use the application with cURL commands:

#### Register a New User

```bash
curl -X POST \
  http://localhost:8080/register \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "new_user",
    "password": "secure_password"
}'
```

This command registers a new user by sending a POST request to the `/register` endpoint with a JSON payload containing the `username` and `password`. Adjust the values accordingly.

#### Login as an Existing User

```bash
curl -X POST \
  http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "existing_user",
    "password": "existing_password"
}'
```

This command logs in an existing user by sending a POST request to the `/login` endpoint with a JSON payload containing the `username` and `password`. Modify the values as needed.

### Additional Notes

- The application uses [Echo](https://echo.labstack.com/), a high-performance HTTP framework for Go, for handling HTTP requests.
- PostgreSQL is used as the database, and the connection details are specified in the `main.go` file.
- User registration and login functionalities are implemented with GORM, a popular Go ORM library.

Adjust the example commands and payload data based on your specific use case. Refer to the API documentation for a comprehensive overview of available endpoints and their functionalities.

## Configuration

Explain any configuration options or environment variables that can be set.

- `DATABASE_URL`: The URL for connecting to the PostgreSQL database.
- `PORT`: The port on which the server will listen.

## Endpoints

Describe the available API endpoints and their functionalities.

- **POST /register**: Register a new user.
  - Request body: JSON with `username` and `password`.
- **POST /login**: Log in an existing user.
  - Request body: JSON with `username` and `password`.