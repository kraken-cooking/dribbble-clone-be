# Dribbble Clone Backend

A Go-based backend service for a Dribbble-like platform, featuring authentication, user profiles, and shot (image) management.

## 🚀 Features

- User Authentication (Login/Signup)
- Profile Management
- Shot Upload and Management
- JWT-based Authorization
- PostgreSQL Database
- Docker Support

## 🛠 Tech Stack

- **Go** - Backend Programming Language
- **Gin** - Web Framework
- **GORM** - ORM Library
- **PostgreSQL** - Database
- **JWT** - Authentication
- **Docker** - Containerization

## 📋 Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Git

## 🔧 Installation & Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd dribbble-clone/be
   ```

2. **Set up environment variables**
   ```bash
   # For development
   cp .env.dev.example .env.dev
   
   # For production
   cp .env.example .env
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

## 🚀 Running the Application

### Development Mode

1. **Start the database**
   ```bash
   docker-compose -f docker-compose.dev.yml --env-file .env.dev up -d
   ```

2. **Run the application locally**
   ```bash
   go run cmd/main.go
   ```

### Production Mode

1. **Start the entire stack**
   ```bash
   docker-compose up -d
   ```

## 🔌 API Endpoints

### Authentication
- `POST /auth/signup` - Register a new user
- `POST /auth/login` - Login user

### Profile
- `GET /profile` - Get user profile
- `PUT /profile` - Update user profile

### Shots
- `POST /shots` - Upload a new shot
- `GET /shots` - List all shots
- `GET /shots/:id` - Get specific shot

## 🗄️ Environment Variables

### Production (.env)
```env
DB_HOST=postgres
DB_USER=<your-db-user>
DB_PASSWORD=<your-db-password>
DB_NAME=<your-db-name>
DB_PORT=5432
JWT_SECRET=<your-jwt-secret>
```

## 📁 Project Structure 