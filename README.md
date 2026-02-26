# Simple User Management API (Go + Gin)

A lightweight RESTful API built with **Go** and the **Gin Gonic** framework for managing user data. This project demonstrates basic CRUD operations using an in-memory data store.

---

## 🚀 Features

- **Create User**: Register new users with unique emails.
- **Read All Users**: List all users currently in the system.
- **Read User by ID**: Fetch specific user details.
- **Update User**: Partial updates (PATCH) for email or password.
- **Delete User**: Remove users from the system by ID.
- **Validation**: Basic checks for duplicate emails and required fields.

---

## 🛠 Tech Stack

* **Language:** Go (Golang)
* **Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
* **Storage:** In-memory Slice (Reset on server restart)
* **Port:** 8888

---

## 📋 Data Models

**User Object**
| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | Integer | Unique identifier |
| `email` | String | User's email address |
| `password` | String | User's password |