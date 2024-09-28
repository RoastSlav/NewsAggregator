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

## Future Improvements

- **User Profiles**
  - Users can save articles.
  - Users can like articles.
  - Users can comment on articles.
- **Search and Filtering**
  - Add a `GET /articles/search` endpoint to allow searching for articles based on keywords, author, or source.
  - Include filtering options (e.g., date range, category) in the `GET /articles` endpoint to refine search results.
- **Pagination**
  - Implement pagination for the `GET /articles` endpoint using query parameters like `page` and `limit` to control the number of articles returned per request.
- **Categories and Tags**
  - Enhance articles with categories (e.g., Technology, Health, Sports) and tags (e.g., AI, Python, Economy).
  - Implement endpoints for filtering articles by category or tags (e.g., `GET /articles?category=technology`).
- **Scheduled Fetching**
  - Use a task scheduler (e.g., cron jobs or Go's `time` package) to periodically fetch articles from external APIs.
  - Introduce an admin endpoint to manually trigger article fetching.
- **Article Analytics**
  - Track article views, most-read articles, and user interaction.
  - Provide an analytics endpoint (e.g., `GET /articles/analytics`) for aggregated statistics like total views, top authors, and popular topics.
- **Bookmarking and Read-Later**
  - Introduce a bookmarking system for users to save articles to read later.
  - Provide endpoints like `POST /articles/{id}/bookmark` and `GET /users/{id}/bookmarks` to manage bookmarks.
- **Authentication and Authorization**
  - Implement user registration and login using JSON Web Tokens (JWT).
  - Restrict certain endpoints (e.g., saving articles, commenting) to authenticated users.
  - Introduce user roles (e.g., admin, editor) with varying access levels to manage content.
- **Comments and Discussions**
  - Develop a commenting system for users to comment on articles.
  - Enable liking or disliking comments and include a moderation system to filter inappropriate content.
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