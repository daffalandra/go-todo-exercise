# Go Product 📋

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat-square&logo=postgresql)
![Bootstrap](https://img.shields.io/badge/Bootstrap-5.3.3-7952B3?style=flat-square&logo=bootstrap)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

A simple, lightweight web application built with **Go**, **GORM**, and **PostgreSQL** to manage categories for a todo-like system. The app provides full CRUD (Create, Read, Update, Delete) functionality for categories, styled with a responsive **Bootstrap 5** interface. Perfect for learning Go web development or building a minimal task management system.

## 🚀 Features
- **Category Management**: Create, view, edit, and delete categories with a user-friendly interface.
- **ORM with GORM**: Seamless database operations with PostgreSQL using GORM.
- **Responsive UI**: Built with Bootstrap 5.3.3 for a clean, mobile-friendly design.
- **Environment Config**: Loads settings from a `.env` file for easy configuration.
- **Form Validation**: Client- and server-side validation for robust user input handling.

## 📂 Project Structure
```
go-todo-exercise/
├── config/                    # Database and environment setup
│   └── config.go
├── controllers/               # HTTP handlers
│   ├── categorycontroller/
│   │   └── categorycontroller.go
│   └── homecontroller/
│       └── homecontroller.go
├── entities/                  # Data models
│   └── entities.go
├── models/                    # Database operations
│   └── categorymodel/
│       └── categorymodel.go
├── views/                     # HTML templates
│   ├── category/
│   │   ├── index.html        # List categories
│   │   ├── add.html          # Add a category
│   │   └── edit.html         # Edit a category
│   └── home/
│       └── index.html        # Home page
├── .env                       # Environment variables
├── main.go                    # Application entry point
└── README.md                  # You're here!
```

## 🛠️ Prerequisites
- **Go**: 1.18 or higher
- **PostgreSQL**: 14 or higher
- **Git**: For cloning the repository
- **Dependencies**:
  - `github.com/joho/godotenv`
  - `gorm.io/gorm`
  - `gorm.io/driver/postgres`

## ⚙️ Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/daffalandra/go-todo-exercise.git
cd go-todo-exercise
```

### 2. Install Dependencies
Initialize the Go module and install dependencies:
```bash
go mod init github.com/daffalandra/go-todo-exercise
go get github.com/joho/godotenv gorm.io/gorm gorm.io/driver/postgres
```

### 3. Configure PostgreSQL
- Ensure PostgreSQL is running.
- Create a database:
  ```sql
  CREATE DATABASE go_product;
  ```
- Create a `.env` file in the project root:
  ```
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=your_username
  DB_PASSWORD=your_password
  DB_NAME=go_product
  ```
  Replace `your_username`, `your_password`, and `go_product` with your PostgreSQL credentials and database name.

### 4. Run the Application
Start the server:
```bash
go run main.go
```
The app will be available at `http://localhost:8080`.

## 🎮 Usage
Access the application in your browser:
- **Home Page**: `http://localhost:8080/` - Welcome page.
- **Categories**:
  - **List**: `http://localhost:8080/categories` - View all categories.
  - **Add**: `http://localhost:8080/categories/add` - Create a new category.
  - **Edit**: `http://localhost:8080/categories/edit?id=<id>` - Update a category.
  - **Delete**: `http://localhost:8080/categories/delete?id=<id>` - Remove a category.

The app features a responsive UI with Bootstrap, form validation, and error messages for a smooth user experience.

## 🗄️ Database Schema
GORM auto-migrates the following schema for the `categories` table:
```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
```
- Soft deletes are enabled via `deleted_at`. To disable, modify `entities.go` and use `db.Unscoped().Delete` in `categorymodel.go`.

## 🌐 Routes
| Method | Path                    | Description                     |
|--------|-------------------------|---------------------------------|
| GET    | `/`                     | Home page                      |
| GET    | `/categories`           | List all categories            |
| GET    | `/categories/add`       | Show form to add a category     |
| POST   | `/categories/add`       | Create a new category          |
| GET    | `/categories/edit`      | Show form to edit a category    |
| POST   | `/categories/edit`      | Update a category              |
| GET    | `/categories/delete`    | Delete a category              |

## 🔮 Future Enhancements
- Add product management for the `/products` route.
- Implement CSRF protection with `github.com/gorilla/csrf`.
- Use `github.com/gorilla/mux` for advanced routing.
- Cache templates for better performance.
- Add unit tests with `testing` and `httptest`.
- Deploy to a cloud platform with HTTPS.

## 🤝 Contributing
We welcome contributions! To contribute:
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/your-feature`.
3. Commit your changes: `git commit -m "Add your feature"`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a pull request.

Please ensure your code follows Go best practices and includes tests where applicable.

## 📜 License
This project is licensed under the [MIT License](LICENSE).

## 📬 Contact
For questions or feedback, reach out to [Daffa Landra](https://github.com/daffalandra) or open an issue on GitHub.

---