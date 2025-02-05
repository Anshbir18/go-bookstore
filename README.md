# 📚 Go Bookstore API

A simple **Book Management REST API** built with **Golang**, using **Gorilla Mux** for routing, **GORM** for database interactions, and **MySQL** for data storage.

## 🚀 Features

✅ **CRUD Operations**: Create, Read, Update, and Delete books  
✅ **Gorilla Mux**: Efficient request routing  
✅ **GORM**: ORM for MySQL database management  
✅ **Structured Project Architecture** for maintainability  

## 🛠️ Tech Stack

- **Golang**  
- **Gorilla Mux** (Router)  
- **GORM** (ORM)  
- **MySQL** (Database)  

## 📂 Project Structure  

```bash
go-bookstore/
│── cmd/
│   ├── main.go             # Entry point of the application
│── pkg/
│   ├── config/             # Database configuration
│   │   ├── app.go
│   ├── controllers/        # Handles API requests
│   │   ├── book-controller.go
│   ├── models/             # Database models
│   │   ├── book.go
│   ├── routes/             # API routes
│   │   ├── bookstore-routes.go
│   ├── utils/              # Utility functions
│   │   ├── utils.go
│── go.mod                  # Go module file
│── go.sum                  # Go dependencies
```



## 🔗 API Endpoints

| Method | Endpoint              | Description            |
|--------|-----------------------|------------------------|
| **GET**    | `/book/`           | Fetch all books        |
| **POST**   | `/book/`           | Add a new book        |
| **GET**    | `/book/{id}`       | Retrieve a book by ID  |
| **PUT**    | `/book/{id}`       | Update book details    |
| **DELETE** | `/book/{id}`       | Remove a book         |

## ⚡ Installation & Setup

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/Anshbir18/go-bookstore.git
cd go-bookstore
```

## 2️⃣ Install Dependencies  
Ensure you have Go installed. Then run:  

```sh
go mod tidy
```

## 3️⃣ Set Up MySQL Database
Create a MySQL database and update the connection details in `config/app.go`:
```sh
dsn := "user:password@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

## 4️⃣ Run the Server
```sh
go run cmd/main.go
```

The server will start on `localhost:8080`.


