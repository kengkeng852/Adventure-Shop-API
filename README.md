# Adventure Shop API

## Project Description

The Adventure Shop API is a RESTful service designed to manage a virtual shop for players, featuring OAuth2 authentication, item management, and player inventory tracking. This project is built using Golang, with PostgreSQL for data persistence, and includes middleware for player and admin authorization.

## Features

### OAuth2 Authentication

- Player login via Google OAuth2.
- Admin login via Google OAuth2.
- Callback handling for player and admin login.
- Logout functionality.

### Item Shop Management

- List available items in the shop.
- Buy items from the shop.
- Sell items to the shop.

### Item Management

- Create new items.
- Edit existing items.
- Archive items.

### Player Coin Management

- Add coins to a player's account.
- View player's coin balance.

### Inventory Management

- List items in a player's inventory.

### Middleware

- Player authorization middleware.
- Admin authorization middleware.

## Getting Started

### Prerequisites

- Go 1.16+
- PostgreSQL

### Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/kengkeng852/adventure-shop-api.git
   cd adventure-shop-api
   ```

2. **Install Dependencies**:

   ```bash
   go mod download
   ```

3. **Set Up the Database**:

   - Update the `config.yaml` file with your PostgreSQL credentials:
     ```yaml
     database:
       host: localhost
       port: 5432
       user: YOUR_DATABASE_USER
       password: YOUR_DATABASE_PASSWORD
       dbname: YOUR_DATABASE_NAME
       sslmode: disable
       schema: public
     ```
   - Run the database migrations:
     ```bash
     go run databases/migration/main.go
     ```

4. **Run the Application**:

   ```bash
   go run main.go
   ```

   The server will start on the port specified in the `config.yaml` file.

## Configuration

The configuration file (`config.yaml`) includes settings for the server, OAuth2 authentication, database connection, and application state. Below is an example:

```yaml
server:
  port: 8000
  allowOrigins:
    - "*"
  bodyLimit: "10M" # MiB
  timeout: 30 # Seconds

oauth2:
  playerRedirectUrl: "http://localhost:8000/v1/oauth2/google/player/login/callback"
  adminRedirectUrl: "http://localhost:8000/v1/oauth2/google/admin/login/callback"
  clientId: "YOUR_GOOGLE_CLIENT_ID"
  clientSecret: "YOUR_GOOGLE_CLIENT_SECRET"
  endpoints:
    authUrl: "https://accounts.google.com/o/oauth2/auth?access_type=offline&approval_prompt=force"
    tokenUrl: "https://oauth2.googleapis.com/token"
    deviceAuthUrl: "https://oauth2.googleapis.com/device/code"
  scopes:
    - "https://www.googleapis.com/auth/userinfo.email"
    - "https://www.googleapis.com/auth/userinfo.profile"
  userInfoUrl: "https://www.googleapis.com/oauth2/v2/userinfo"
  revokeUrl: "https://accounts.google.com/o/oauth2/revoke"

database:
  host: localhost
  port: 5432
  user: YOUR_DATABASE_USER
  password: YOUR_DATABASE_PASSWORD
  dbname: YOUR_DATABASE_NAME
  sslmode: disable
  schema: public

state:
  secret: "YOUR_SECRET_KEY"
  expiresAt: 120 #seconds
  issuer: "adventureshop"
```

## API Endpoints

### OAuth2

- `GET /v1/oauth2/google/player/login`
- `GET /v1/oauth2/google/admin/login`
- `GET /v1/oauth2/google/player/login/callback`
- `GET /v1/oauth2/google/admin/login/callback`
- `DELETE /v1/oauth2/google/logout`

### Item Shop

- `GET /v1/item-shop`
- `POST /v1/item-shop/buying`
- `POST /v1/item-shop/selling`

### Item Management

- `POST /v1/item-managing`
- `PATCH /v1/item-managing/:itemID`
- `DELETE /v1/item-managing/:itemID`

### Player Coin

- `POST /v1/player-coin`
- `GET /v1/player-coin`

### Inventory

- `GET /v1/inventory`






