# Restaurant Management Backend API 🍔🍴

This is a backend API for a restaurant management system, implemented in Golang.

## Overview

This project provides a backend API for a restaurant management system. It includes controllers, models, routes, and middleware to handle various restaurant-related operations. 
More about backend api architecture in GO: https://github.com/amitshekhariitbhu/go-backend-clean-architecture

## Features

- Create, update, and delete restaurant items (e.g., menu items).
- Manage orders and their statuses.
- User authentication and authorization.
- ...

## Project Structure

The project is organized as follows:

- `controllers/`: Contains controller logic for different endpoints.
- `database/`: Database setup and migrations.
- `helpers/`: Utility functions.
- `middleware/`: Middleware for authentication, logging, etc.
- `models/`: Data models and database schema.
- `routes/`: Defines API routes and connects them to controllers.
- `main.go`: Main application entry point.
- `go.mod`, `go.sum`: Dependency management files.

## Prerequisites

- Golang (version X.X.X) 🐹
- MongoDB (or other preferred database) 🛢️

## Getting Started

Clone this repository & run the commands:

```bash
git clone https://github.com/shnartho/restaurant-management-backend.git
cd restaurant-management-backend

// Install project dependencies:
go mod tidy
Set up your database configuration in database/db.go.

// Run database migrations:
go run database/migrations.go

// Start the server:
go run main.go
```

## The API server should now be running at http://localhost:8080 🚀.

