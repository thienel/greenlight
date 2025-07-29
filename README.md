# Greenlight

A REST API for managing a movie database, built with Go and PostgreSQL.

## Features

- Movie management with CRUD operations
- User authentication and registration
- Token-based authorization with permissions
- Email user activation
- Rate limiting
- Movie filtering and pagination
- Application metrics
- CORS support
- JSON logging
- Input validation
- Comprehensive testing

## Tech Stack

- Go 1.24.4
- PostgreSQL
- httprouter
- bcrypt password hashing
- SMTP email integration
- SQL migrations

## API Endpoints

### Health Check
- `GET /v1/healthcheck` - API health status

### Movies
- `GET /v1/movies` - List movies with filtering and pagination
- `POST /v1/movies` - Create a movie (requires movies:write permission)
- `GET /v1/movies/:id` - Get a movie (requires movies:read permission)
- `PATCH /v1/movies/:id` - Update a movie (requires movies:write permission)
- `DELETE /v1/movies/:id` - Delete a movie (requires movies:write permission)

### Users
- `POST /v1/users` - Register a new user
- `PUT /v1/users/activated` - Activate user account

### Authentication
- `POST /v1/tokens/authentication` - Authenticate and get token

### Metrics
- `GET /debug/vars` - Application metrics

## Getting Started

### Prerequisites

- Go 1.24.4 or later
- PostgreSQL

### Installation

1. Clone the repository:
```bash
git clone https://github.com/thienel/greenlight
cd greenlight
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables:
```bash
export GREENLIGHT_DB_DSN="postgres://username:password@localhost/greenlight?sslmode=disable"
```

4. Run database migrations:
```bash
make db/migrations/up
```

### Running the Application

```bash
# Using Make
make run/api

# Or directly with Go
go run ./cmd/api -db-dsn="postgres://username:password@localhost/greenlight?sslmode=disable"
```

## Configuration

The application accepts the following command-line flags:

- `-port`: Server port (default: 4000)
- `-env`: Environment (development|staging|production)
- `-db-dsn`: PostgreSQL DSN
- `-db-max-open-conns`: Maximum open database connections (default: 25)
- `-db-max-idle-conns`: Maximum idle database connections (default: 25)
- `-db-max-idle-time`: Maximum idle connection time (default: 15m)
- `-limiter-rps`: Rate limiter requests per second (default: 2)
- `-limiter-burst`: Rate limiter burst size (default: 4)
- `-limiter-enabled`: Enable/disable rate limiting (default: true)
- `-smtp-host`: SMTP server host
- `-smtp-port`: SMTP server port (default: 2525)
- `-smtp-username`: SMTP username
- `-smtp-password`: SMTP password
- `-smtp-sender`: SMTP sender address
- `-cors-trusted-origins`: Trusted CORS origins (space separated)

## Database

The application uses PostgreSQL with these tables:

- **movies**: Store movie information (title, year, runtime, genres)
- **users**: User accounts and authentication
- **tokens**: Authentication tokens
- **permissions**: User permissions system

Migration files in the `migrations/` directory handle schema creation and updates.

## Authentication

The application uses token-based authentication:

1. Users register and receive an activation email
2. After activation, users can authenticate to receive tokens
3. Tokens are required for protected endpoints
4. Permissions control access to specific operations:
   - `movies:read` - Read movie data
   - `movies:write` - Create, update, delete movies

## Development

### Available Make Commands

```bash
make help                # Show available commands
make run/api             # Run the API server
make db/psql             # Connect to database
make db/migrations/new   # Create new migration
make db/migrations/up    # Apply migrations
make audit               # Run quality control checks
```

### Testing

Run tests with:

```bash
# Run all tests
go test ./...

# Run with race detection
go test -race ./...

# Run quality audit (includes tests, formatting, static analysis)
make audit
```

The test suite includes unit tests for all packages. Integration and end-to-end tests require a database connection.

## API Examples

### Create a Movie
```bash
curl -X POST http://localhost:4000/v1/movies \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "title": "The Shawshank Redemption",
    "year": 1994,
    "runtime": "142 mins",
    "genres": ["Drama"]
  }'
```

### List Movies with Filters
```bash
curl "http://localhost:4000/v1/movies?title=shawshank&genres=drama&page=1&page_size=5&sort=year" \
  -H "Authorization: Bearer <token>"
```

### Register a User
```bash
curl -X POST http://localhost:4000/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```
