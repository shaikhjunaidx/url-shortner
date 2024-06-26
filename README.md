# URL Shortener

A URL shortener service built with Go, Redis, and Go Fiber. This service allows users to shorten URLs and redirects to the original URL when accessed through the shortened link. It also includes rate limiting and domain validation features.

## Features

- Shorten URLs and generate custom short links
- Redirect to original URLs using the short link
- Rate limiting to prevent abuse
- Domain validation to avoid certain domains
- Dockerized for easy setup and deployment

## Installation

### Prerequisites

- Go 1.16+
- Docker and Docker Compose

### Steps

1. Clone the repository:

    ```sh
    git clone https://github.com/shaikhjunaidx/url-shortener.git
    cd url-shortener
    ```

2. Build and run the Docker containers:

    ```sh
    docker-compose up --build
    ```

## Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory and set the following variables:

```env
DB_ADDR="db:6379"
DB_PASS=""
APP_PORT=":3000"
DOMAIN="localhost:3000"
API_QUOTA=10
```

## Endpoints

### `POST /api/v1`

- **Description**: Shortens a given URL.

- **Request Body**:

    ```json
    {
        "url": "https://example.com",
        "short": "customShortLink",  // optional
        "expiry": 24                // optional, in hours
    }
    ```

- **Response**:

    ```json
    {
        "url": "https://example.com",
        "short": "http://localhost:3000/abc123",
        "expiry": 24,
        "rate_limit": 10,
        "rate_limit_reset": 30
    }
    ```

### `GET /:url`

- **Description**: Redirects to the original URL for the given short URL.
- **Response**: Redirects to the original URL.


## Usage

### Using Postman

1. **Install Postman**:
   - Download and install Postman from [Postman’s official website](https://www.postman.com/downloads/).

2. **Create a new POST request** to shorten a URL:
   - **URL**: `http://localhost:3000/api/v1`
   - **Method**: `POST`
   - **Headers**: 
     - `Content-Type: application/json`
   - **Body**: Select `raw` and `JSON` format, then enter the following JSON:
     ```json
     {
         "url": "https://example.com",
         "short": "customShortLink",  // optional
         "expiry": 24                // optional, in hours
     }
     ```

3. **Send the request** and you should receive a response similar to:
    ```json
    {
        "url": "https://example.com",
        "short": "http://localhost:3000/abc123",
        "expiry": 24,
        "rate_limit": 10,
        "rate_limit_reset": 30
    }
    ```

4. **Create a new GET request** to test URL redirection:
   - **URL**: `http://localhost:3000/abc123`
   - **Method**: `GET`

   Opening this URL in your browser should redirect you to the original URL, `https://example.com`.


## Summary

### Features

- **URL Shortening**: Converts long URLs into short, easy-to-share links.
- **Custom Short Links**: Allows users to create custom short links.
- **Redirection**: Automatically redirects short URLs to their original long URLs.
- **Rate Limiting**: Prevents abuse by limiting the number of URL shortening requests.
- **Domain Validation**: Ensures that certain domains can be avoided or handled specifically.
- **Dockerized**: Easily deployable using Docker and Docker Compose.
- **Configuration**: Easily configurable via environment variables.


### Potential Use Cases

- **Social Media**: Share concise, manageable links within character limits.
- **Marketing Campaigns**: Manage and track performance of multiple campaign links.
- **Email Campaigns**: Clean, trackable links for email content.
- **Event Management**: Easy-to-share registration and information links for events.

