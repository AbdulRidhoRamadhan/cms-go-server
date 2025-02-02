# CMS-Go-Project

A Content Management System (CMS) built with Go, featuring user authentication, product management, and category organization.

## About This Project

This project is a rewrite of a CMS backend that was originally built using Node.js (Express.js). The goal of this migration to Go (Gin) was to improve performance, efficiency, and maintainability while also serving as a learning experience for backend migration between different web frameworks.

Through this project, I explored:

- The differences in architecture, request handling, and database interactions between Node.js and Go.
- Best practices in containerization with Docker and automating CI/CD using GitHub Actions.

This showcases my ability to adapt to new technologies, work with modern DevOps practices, and apply best practices in backend development.

## Features

- 🔐 JWT Authentication & Role-based Authorization (Admin/Staff)
- 📦 Product & 🏷️ Category Management
- 🖼️ Image Upload (Cloudinary)
- 🔍 Product Search & Filtering
- 📝 CRUD Operations
- 🐳 Docker Containerization
- 🔄 CI/CD with GitHub Actions
- 📊 Database Migrations

## Tech Stack

- **Backend**: Go (Gin Framework)
- **Database**: PostgreSQL (GORM)
- **Authentication**: JWT
- **Image Storage**: Cloudinary
- **Containerization**: Docker
- **CI/CD**: GitHub Actions

## 🔑 Credentials

To try out the application, use the following credentials:

- **Admin Interface**:
  - Email: `demo@mail.com`
  - Password: `demo123`

- **Public Interface**:
  - No login required
 
**Note:** Accounts with admin roles have full access to all system features and settings, but for security reasons, credentials for role admin accounts are not included in this demo.

The demo account provided is a staff role account, which has the following limitations:
- Can update, delete, and update images using the upload feature for posts they have created.
- Cannot create other staff accounts.
- Cannot update, delete, and update images using the upload feature for posts created by other users.
## Prerequisites

- Go 1.22+
- PostgreSQL
- Docker (optional)
- Cloudinary Account

## Setup

1. Clone the repository:

   ```bash
   https://github.com/AbdulRidhoRamadhan/cms-go-server.git
   ```

2. Create a `.env` file:

   ```env
   JWT_SECRET=your_jwt_secret
   PORT=80
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   CLOUDINARY_URL=your_cloudinary_url
   ADMIN_DEFAULT_PASSWORD=your_admin_password
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```
