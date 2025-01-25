# User System API

A Web API application for managing users and their addresses, implemented using Golang and the Echo framework.

## Features
- User management with CRUD operations.
- Address management associated with users.
- Environment-based configuration using `.env` files.
- Database integration with PostgreSQL.

## Getting Started

Follow the instructions below to set up and run the project on your local machine.

### Prerequisites
- [Golang](https://golang.org/doc/install) (Version 1.19 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Copy the `.env.example` file to `.env`:
   ```bash
   cp .env.example .env
   ```

3. Update the `.env` file with your custom configuration (if necessary).

### Environment Variables

The application uses the following environment variables:

```plaintext
APP_NAME="UserSystem"       # Name of the application
APP_ENV=prod                # Environment (e.g., dev, test, prod)
APP_PORT=3000               # Port for the API server

DB_HOST=localhost           # Database host
DB_PORT=5432                # Database port
DB_USER=postgres            # Database user
DB_PASSWORD=postgres        # Database password
DB_NAME=users               # Database name
DB_SSL_MODE=disable         # Database SSL mode
```

### Running the Application

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. The API will be available at:
   ```
   http://localhost:3000 #3000 is example port
   ```

### Database Setup

1. Ensure that your PostgreSQL database is running.
2. Use the credentials provided in your `.env` file to create the database.

### Technologies Used

- **Language**: Golang
- **Framework**: Echo
- **Database**: PostgreSQL
- **Configuration**: `.env` files

### License
This project is licensed under the MIT License.
