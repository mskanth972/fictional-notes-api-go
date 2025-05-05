# containerized-api-server
A containerized web application with RESTful CRUD operations, JWT-based authentication, and rate limiting, deployed on AWS using Docker and cloud-native best practices.

# Fictional Notes API – Adobe Take-Home Assessment

This project is a containerized web application built using **Golang** and the **Gin** framework. It simulates a simple note-taking backend with features like **CRUD operations**, **token-based authentication using JWT**, and **rate limiting**. The application also includes **Swagger UI documentation** for easy API exploration and testing.

---

## 🧩 Project Overview

The API is designed to demonstrate:

- Secure login functionality using JWT tokens
- Standard RESTful CRUD operations (`/notes`)
- Authorization protection on all routes except login
- Simple rate limiting (5 requests per minute)
- Self-documented API via Swagger

---

## 🚀 How to Build and Run the Project

### Prerequisites
- [Docker](https://www.docker.com/) must be installed and running on your system.

### Steps

1. **Build the Docker image**
   ```bash
   docker build -t fictional-notes-api-go .
   ```

2. **Run the container**
   ```bash
   docker run -p 8000:8000 fictional-notes-api-go
   ```

3. **Access the API**
   Open your browser and navigate to:
   ```
   http://localhost:8000/swagger/index.html
   ```

---

## 🔐 Authentication Flow

1. **Login**
   - Endpoint: `POST /login`
   - Credentials:
     ```json
     {
       "username": "admin",
       "password": "password"
     }
     ```
   - Returns:
     ```json
     {
       "access_token": "<JWT>",
       "token_type": "bearer"
     }
     ```

2. **Authorize**
   - Use the **Authorize** button in Swagger UI.
   - Enter:
     ```
     Bearer <access_token>
     ```

3. **Access Notes Endpoints**
   - All `/notes` routes require the token.

---

## 📚 Available Endpoints

| Method | Path          | Description         | Auth Required |
|--------|---------------|---------------------|---------------|
| POST   | `/login`      | Obtain JWT Token    | ❌ No         |
| GET    | `/notes`      | List all notes      | ✅ Yes        |
| POST   | `/notes`      | Create a note       | ✅ Yes        |
| DELETE | `/notes/{id}` | Delete a note by ID | ✅ Yes        |

---

## ✅ Assessment Requirements Coverage

- [x] API with fictional CRUD operations
- [x] Token-based authentication using JWT
- [x] Rate limiting per client
- [x] Containerized (Docker)
- [x] Cloud-deployable
- [x] API documentation via Swagger

---

## 📎 Final Notes

- The application resets all notes and rate limits upon restart.
- It is self-contained and doesn’t require any external database or services.
