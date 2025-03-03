# Simple User Management API using MVC in Golang

## 🚀 Overview
This project demonstrates how to implement a **Simple User Management API** in Golang using the **MVC architecture pattern**. The architecture separates the application's business logic, data access, and presentation layers.

## 📌 Project Structure
```
user_management/
│── models/           # User Model (Entities)
│   └── user.go
│── controllers/      # API Handlers (Controllers)
│   └── user_controller.go
│── database/      # Data Access Layer
│   └── db.go
│── main.go          # Application Entry Point
└── go.mod           # Go Modules File
```

---

## 🔑 Technology Stack
- Golang
- Gin (Web Framework)

## Install
- `go get github.com/gin-gonic/gin`
- `go get gorm.io/gorm`
- `go get gorm.io/driver/sqlite`


---

## 🔥 How It Works
### 1. **Model Layer (User Entity)**
Defines the core business model.

`models/user.go`
```go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### 2. **Repository or Database Layer (Data Access)**
Handles user data storage and retrieval.

`database/db.go`
```go
package database

import (
    "backend-architecture/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    db.AutoMigrate(&models.User{})
    DB = db
}
```

### 3. **Controller Layer (API Handlers)**
Exposes the API endpoints.

`controllers/user_controller.go`
```go
package controllers

import (
    "backend-architecture/models"
    "backend-architecture/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&user)
    c.JSON(http.StatusCreated, user)
}
```

### 4. **Route**
Starts the server and injects dependencies.

`routes/routes.go`
```go
package routes

import (
    "backend-architecture/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
    return r
}
```

### 5. **Main Application (Entry Point)**
Starts the server and injects dependencies.

`main.go`
```go
package main

import (
    "backend-architecture/database"
    "backend-architecture/routes"
)

func main() {
    database.InitDB()
    r := routes.SetupRouter()
    r.Run(":8080")
}```

---

## 🎯 How to Run the Project
### 1. Clone the Repository
```bash
git clone https://github.com/penadidik/mvc_and_clean_architecture_golang.git
cd mvc_and_clean_architecture_golang
```

### 2. Install Dependencies
```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### 3. Run the Application
```bash
go run main.go
```

### 4. Test API with CURL
**Create User:**
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"John Doe", "email":"john@example.com"}'
```

**Retrieve Users:**
```bash
curl -X GET http://localhost:8080/users
```

---

## ✅ Why Use MVC?
- Separation of Concerns
- Modular Code
- Easy to Maintain and Extend

---

## 📄 License
This project is licensed under the MIT License.

---

## 📌 Author
[Didik Maryono]

Let's Build Simple and Clean Systems! 🚀


