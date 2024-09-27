# NewsAggregator

NewsAggregator is a Go-based application that fetches articles from various news sources and serves them via a RESTful API.

## Features

- Fetch articles from external news APIs.
- Serve articles through a RESTful API.
- Connect to a MySQL database for storing articles.

## Requirements

- Go 1.23.1 or later
- MySQL database
- Environment variables for database configuration

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/NewsAggregator.git
    cd NewsAggregator
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your environment variables in a `config.env` file:
    ```env
    DB_USER=your_db_user
    DB_PASS=your_db_password
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    DB_NAME=your_db_name
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start, and you can access the API at `http://localhost:8080`.

## API Endpoints

- `GET /articles`: Retrieve all articles.

## Project Structure

- `main.go`: Entry point of the application.
- `internal/api.go`: Contains the API handlers.
- `internal/articles`: Contains logic for fetching and processing articles.
- `database/database.go`: Handles database connection.
- `go.mod`: Go module dependencies.