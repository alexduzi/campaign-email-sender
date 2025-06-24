# Campaign Email Sender

A Go-based microservice for managing and sending email campaigns with authentication, built using Clean Architecture principles.

## 🚀 Features

- **Campaign Management**: Create, read, update, and delete email campaigns
- **Authentication**: JWT-based authentication with Keycloak integration
- **Email Sending**: Automated email delivery with SMTP support
- **Background Worker**: Separate worker process for handling email delivery
- **Status Tracking**: Real-time campaign status monitoring
- **Database Integration**: PostgreSQL with GORM ORM
- **Testing**: Comprehensive unit tests with mocks
- **Development Tools**: Hot reload with Air

## 🏗️ Architecture

The project follows Clean Architecture principles with the following structure:

```
├── cmd/
│   ├── api/         # Main API server
│   └── worker/      # Background email worker
├── internal/
│   ├── contract/    # Data transfer objects
│   ├── domain/      # Business logic and entities
│   ├── endpoints/   # HTTP handlers
│   ├── infrastructure/ # External dependencies
│   └── test/        # Test utilities and mocks
├── docs/           # API documentation
└── docker-compose.yml # Development environment
```

## 🛠️ Prerequisites

- **Go 1.24+**
- **Docker & Docker Compose**
- **Git**

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd campaignemailsender
```

### 2. Set Up Environment Variables

Copy the example environment file and configure it:

```bash
cp .env.EXAMPLE .env
```

Edit the `.env` file with your configuration:

```env
DATABASE="host=localhost user=postgres password=1234567 dbname=campaigndb port=5433 sslmode=disable"
KEYCLOAK="http://localhost:8080/realms/provider"
CLIENTID="emailn"
EMAIL_SMTP="smtp.gmail.com"
EMAIL_USER="your-email@gmail.com"
EMAIL_PASSWORD="your-app-password"
```

### 3. Start the Development Environment

Start all required services (PostgreSQL, PgAdmin, Keycloak):

```bash
docker-compose up -d
```

This will start:
- **PostgreSQL** on port `5433`
- **PgAdmin** on port `5050` (admin@admin.com / password)
- **Keycloak** on port `8080` (admin / admin123)

### 4. Configure Keycloak

1. Access Keycloak at http://localhost:8080
2. Login with `admin` / `admin123`
3. Create a new realm called `provider`
4. Create a client with ID `emailn`
5. Create a user for testing

### 5. Install Dependencies

```bash
go mod download
```

### 6. Run Database Migrations

The application will automatically create the necessary tables on startup.

### 7. Start the API Server

Using Air for hot reload (recommended for development):

```bash
# Install Air if you haven't already
go install github.com/air-verse/air@latest

# Start with hot reload
air
```

Or run directly:

```bash
go run cmd/api/main.go
```

The API will be available at http://localhost:3000

### 8. Start the Email Worker (Optional)

In a separate terminal:

```bash
go run cmd/worker/main.go
```

## 📡 API Endpoints

### Authentication
All campaign endpoints require a valid JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/ping` | Health check |
| POST | `/campaigns` | Create a new campaign |
| GET | `/campaigns` | List all campaigns |
| GET | `/campaigns/{id}` | Get campaign by ID |
| DELETE | `/campaigns/{id}` | Delete a campaign |
| POST | `/campaigns/start/{id}` | Start a campaign |

### Example Requests

#### Get JWT Token
```http
POST http://localhost:8080/realms/provider/protocol/openid-connect/token
Content-Type: application/x-www-form-urlencoded

client_id=emailn&username=your-email@gmail.com&password=your-password&grant_type=password
```

#### Create Campaign
```http
POST http://localhost:3000/campaigns
Authorization: Bearer <your-token>
Content-Type: application/json

{
    "name": "My Campaign",
    "content": "Hello everyone! This is a test email.",
    "emails": ["recipient1@example.com", "recipient2@example.com"]
}
```

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test ./... -cover
go test ./... -coverprofile=coverfile
go tool cover -html="coverfile"
```

Run tests for a specific package:

```bash
go test ./internal/domain/campaign/...
```

## 📊 Database Access

Access PgAdmin at http://localhost:5050:
- **Email**: admin@admin.com
- **Password**: password

Database connection details:
- **Host**: pg-docker (or localhost if connecting externally)
- **Port**: 5432 (internal) / 5433 (external)
- **Database**: campaigndb
- **Username**: postgres
- **Password**: 1234567

## 🔧 Development Tools

### Hot Reload with Air

The project includes Air configuration for hot reloading during development. The configuration is in `.air.toml`.

### REST Client

Use the provided `docs/rest.http` file with your HTTP client (VS Code REST Client, IntelliJ HTTP Client, etc.) for testing the API.

## 📝 Campaign Workflow

1. **Create**: Campaign starts in "Pending" status
2. **Start**: Change status to "Started" to queue for sending
3. **Processing**: Worker picks up started campaigns and sends emails
4. **Complete**: Status changes to "Done" or "Fail" based on result

## 🔒 Security

- JWT authentication via Keycloak
- Email validation for all recipients
- Input validation using Go validator
- Secure password handling for SMTP

## 🚨 Common Issues

### Email Not Sending
- Verify SMTP credentials and server settings
- For Gmail, use App Passwords instead of regular passwords
- Check firewall settings for SMTP port (587)

### Database Connection Issues
- Ensure PostgreSQL container is running
- Check database connection string in `.env`
- Verify port 5433 is not in use by another service

### Authentication Errors
- Verify Keycloak is running and configured correctly
- Check that the realm and client are properly set up
- Ensure JWT token is valid and not expired