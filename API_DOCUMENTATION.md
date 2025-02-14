# URL Shortener

URL Shortener is a web application that allows you to shorten long URLs into short URLs, easy to share.

## Features

- URL shortening
- Redirection to the long URL from the short URL
- Storage of URL mappings in a SQLite database

## Prerequisites

- Go 1.23 or higher
- SQLite

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/gkehren/url_shortener.git
   cd url-shortener
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Configure the database:

   - The project uses SQLite for URL storage. A `urls.db` file will be created in the project's root directory.

## Usage

1. Start the application:

   ```bash
   go run cmd/url_shortener/main.go
   ```

2. Access the application via `http://localhost:8080`.

## URL Shortener API Documentation

# URL Shortener API Documentation

This documentation provides details on how to interact with the URL Shortener API. The API allows users to shorten long URLs and resolve short URLs to their original long URLs.

## Base URL

The base URL for all API endpoints is:

```
http://localhost:8080
```

## Endpoints

### 1. Shorten URL

- **URL**: `/shorten`
- **Method**: `POST`
- **Description**: Shortens a long URL and returns a short URL.
- **Request Body**:
  ```json
  {
    "long_url": "string"
  }
  ```
  - **long_url** (string, required): The long URL to be shortened.
- **Response**:
  - **Success**: `(Status: 200 OK)`
    ```json
    {
        "short_url": "string"
    }
    ```
    - **short_url** (string): The generated short URL.
  - **Error**: `(Status: 500 Internal Server Error)`
      ```json
      {
        "error": "string"
      }
      ```
      - **error** (string): A message describing the error.

### 2. Resolve URL

- **URL**: `/:short_url`
- **Method**: `GET`
- **Description**: Resolves a short URL to its corresponing long URL.
- **Path Parameter**:
  - **short_url** (string, required): The short URL to be resolved.
- **Response**:
  - **Success**: `(Status: 302 Found)`
    - Redirects to the original long URL.
  - **Error**: `(Status: 404 Not Found)`
      `URL not found`

## Examples

### Shorten URL

**Request**:

```curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"long_url": "https://www.google.com"}'```

**Response**:

```json
{
    "short_url": "gBti0D"
}
```

### Resolve URL

**Request**:

```curl -X GET http://localhost:8080/gBti0D```

**Response**:
- Redirects to `https://www.google.com`.

## Error Handling

- **404 Not Found**: The requested ressource could not be found.
- **500 Internal Server Error**: An error occurred on the server.