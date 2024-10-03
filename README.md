# NewsAggregator

NewsAggregator is a Go-based application that fetches articles from various news sources and serves them via a RESTful
API. I'm using this project to learn Go and improve my skills in building web applications.

## Features

- Fetch articles from external news APIs.
- Serve articles through a RESTful API.
- Connect to a MySQL database for storing articles.
- User profiles with registration and login.
- Pagination for articles.
- Search for articles.
- Automatic cron updates for fetching articles.
- Categories for articles (e.g., Technology, Health, Sports, AI, Python, Economy).
- Users can save articles to read later.
- Users can comment on articles
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
- `GET /articles/{id}`: Retrieve a specific article by ID.
- `GET /articles/search`: Search and filter articles.
- `POST /articles/like/{id}`: Like or dislike an article.
- `POST /articles/comment/{id}`: Comment on an article.
- `GET /articles/comments/{id}`: Retrieve comments for an article.
- `POST /articles/read-later/{id}`: Add or remove an article from read later.
- `GET /articles/read-later`: Retrieve read later articles.
- `GET /articles/category/{name}`: Retrieve articles by category.
- `GET /articles/category`: Retrieve all categories.
- `POST /user/register`: Register a new user.
- `POST /user/login`: Log in an existing user.

## Project Structure

- `cmd/newsAggregator/main.go`: Entry point of the application.
- `internal/`: Core application logic.
    - `articles/`: Contains files related to article handling.
        - `articleRepository.go`: Handles database interactions for articles.
        - `articleService.go`: Business logic for articles.
        - `models.go`: Data models for articles.
    - `database/`: Database connection and setup.
        - `database.go`: Handles the database connection.
    - `routes/`: API routes and handler registration.
        - `routes.go`: Defines routes for the application.
    - `users/`: Contains user-related files.
        - `models.go`: User data models.
        - `sessionRepository.go`: Session management for users.
        - `userRepository.go`: Database interactions for users.
        - `userService.go`: Business logic for users.
    - `util/`: Utility functions.
        - `errorUtil.go`: Error handling utilities.
        - `httpBodyUtil.go`: Helper functions for HTTP body processing.
- `config.env`: Environment variables.
- `go.mod`: Go module dependencies.

## Libraries Used

- `github.com/go-sql-driver/mysql`: MySQL driver for Go.
- `github.com/jmoiron/sqlx`: General purpose extensions to database/sql.
- `github.com/joho/godotenv`: Go port of Ruby's dotenv library (loads environment variables from `.env`).
- `golang.org/x/crypto`: Supplementary Go cryptography libraries.
- `filippo.io/edwards25519`: Go implementation of the Edwards-curve Digital Signature Algorithm (EdDSA).
- `github.com/robfig/cron/v3`: Cron library for scheduling tasks.

## Future Improvements

- **Article Analytics**
    - Track article views, most-read articles, and user interaction.
    - Provide an analytics endpoint (e.g., `GET /articles/analytics`) for aggregated statistics like total views, top authors, and popular topics.
- **Rate Limiting and Throttling**
    - Introduce rate limiting to prevent abuse of the API by controlling the number of requests a user can make within a given time frame.
    - Use middleware to enforce rate limits and return appropriate responses.
- **Article Caching**
    - Integrate caching for articles using an in-memory store (e.g., Redis) to speed up repeated requests and reduce database load.
    - Implement cache invalidation when articles are updated or new ones are fetched.
- **Notifications**
    - Implement real-time notifications using WebSockets or Server-Sent Events (SSE) to notify users when new articles of interest are available.
    - Introduce email notifications or push notifications for subscribed users.
- **Content Recommendations**
    - Implement a recommendation system using collaborative or content-based filtering algorithms to suggest related articles based on user history.
- **Article Summarization**
    - Use Natural Language Processing (NLP) libraries (e.g., spaCy, Golang text analysis packages) to generate article summaries.
    - Add an endpoint (`GET /articles/{id}/summary`) to provide the summarized version of an article.