# Greenlight

A modern REST API for managing a movie database, built with Go and PostgreSQL. Greenlight provides a robust backend service with user authentication, authorization, and comprehensive movie management capabilities.

## Features

- 🎬 **Movie Management**: Complete CRUD operations for movies
- 👥 **User Authentication**: Secure user registration and login system
- 🔐 **Authorization**: Role-based permissions system
- 📧 **Email Integration**: User activation via email
- 🚦 **Rate Limiting**: API rate limiting to prevent abuse
- 🔍 **Filtering & Pagination**: Advanced movie search and filtering
- 📊 **Metrics**: Application metrics via expvar
- 🌐 **CORS Support**: Cross-origin resource sharing configuration
- 📝 **Structured Logging**: JSON-based logging
- 🛡️ **Input Validation**: Comprehensive request validation

## Tech Stack

- **Language**: Go 1.24.4
- **Database**: PostgreSQL
- **Router**: httprouter
- **Email**: SMTP integration
- **Migrations**: SQL migration files
- **Encryption**: bcrypt for password hashing

## Project Structure

```
├── cmd/
│   ├── api/           # API application entry point
│   └── examples/      # Example implementations (CORS)
├── internal/
│   ├── data/          # Data models and database operations
│   ├── jsonlog/       # JSON logging utilities
│   ├── mailer/        # Email sending functionality
│   └── validator/     # Input validation
├── migrations/        # Database migration files
├── mocks/            # Test mocks
├── remote/           # Remote deployment scripts
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
└── Makefile          # Build and development commands
```

## API Endpoints

### Health Check
- `GET /v1/healthcheck` - API health status

### Movies
- `GET /v1/movies` - List movies (with filtering and pagination)
- `POST /v1/movies` - Create a new movie
- `GET /v1/movies/:id` - Get a specific movie
- `PATCH /v1/movies/:id` - Update a movie
- `DELETE /v1/movies/:id` - Delete a movie

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
- Make (optional, for using Makefile commands)

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

3. Set up your environment variables in `.envrc`:
```bash
export GREENLIGHT_DB_DSN="postgres://username:password@localhost/greenlight?sslmode=disable"
```

4. Run database migrations:
```bash
make db/migrations/up
```

### Running the Application

#### Using Make (Recommended)
```bash
# Run the API server
make run/api

# Connect to the database
make db/psql

# View all available commands
make help
```

#### Direct Go Command
```bash
go run ./cmd/api -db-dsn="postgres://username:password@localhost/greenlight?sslmode=disable"
```

### Configuration Options

The application accepts the following command-line flags:

- `-port`: Server port (default: 4000)
- `-env`: Environment (development|staging|production)
- `-db-dsn`: PostgreSQL DSN
- `-db-max-open-conns`: Maximum open database connections
- `-db-max-idle-conns`: Maximum idle database connections
- `-db-max-idle-time`: Maximum idle connection time
- `-limiter-rps`: Rate limiter requests per second
- `-limiter-burst`: Rate limiter burst size
- `-limiter-enabled`: Enable/disable rate limiting
- `-smtp-host`: SMTP server host
- `-smtp-port`: SMTP server port
- `-smtp-username`: SMTP username
- `-smtp-password`: SMTP password
- `-smtp-sender`: SMTP sender address
- `-cors-trusted-origins`: Trusted CORS origins

## Database Schema

The application uses PostgreSQL with the following main tables:

- **movies**: Store movie information (title, year, runtime, genres)
- **users**: User accounts and authentication
- **tokens**: Authentication tokens
- **permissions**: User permissions system

Migration files are located in the `migrations/` directory and handle:
- Table creation
- Constraints and indexes
- Permissions setup

## Authentication & Authorization

Greenlight implements a token-based authentication system:

1. Users register and receive an activation email
2. After activation, users can authenticate to receive tokens
3. Tokens are required for protected endpoints
4. Permissions control access to specific operations

### Permissions

- `movies:read` - Read movie data
- `movies:write` - Create, update, delete movies

## Development

### Available Make Commands

```bash
make help                 # Show available commands
make run/api             # Run the API server
make db/psql             # Connect to database
make db/migrations/new   # Create new migration
make db/migrations/up    # Apply migrations
make db/migrations/down  # Rollback migrations
make audit               # Run quality control checks
make build/api           # Build API binary
```

### Testing

Run tests with:
```bash
go test ./...
```

### Code Quality

The project includes quality control checks:
```bash
make audit
```

## API Usage Examples

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

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Built following Go best practices
- Inspired by modern REST API design principles
- Uses industry-standard authentication patterns
