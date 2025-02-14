# URL Shortener

URL Shortener is a web application that allows you to shorten long URLs into short, shareable links. The project consists of a backend developed in Go and a frontend using SvelteKit with Vite and Tailwind CSS.

## Features

- **URL Shortening**: Convert a long URL into a short URL.
- **Redirection**: Access the long URL using the short URL.
- **User Interface**: A simple and efficient interface to interact with the application.
- **URL Storage**: Uses a SQLite database to store URL mappings.

## Prerequisites

- Go 1.23 or later
- Node.js and npm (for the frontend)
- Docker and Docker Compose (for deployment)

## Installation

### Backend

1. **Clone the repository**:

   ```bash
   git clone https://github.com/gkehren/url-shortener.git
   cd url-shortener/backend
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   - The project uses SQLite for URL storage. Ensure the `urls.db` file is accessible.

### Frontend

1. **Navigate to the frontend directory**:

   ```bash
   cd ../frontend
   ```

2. **Install dependencies**:

   ```bash
   npm install
   ```

3. **Start the development server**:

   ```bash
   npm run dev
   ```

## Usage

1. **Start the backend**:

   ```bash
   go run cmd/url-shortener/main.go
   ```

2. **Access the application**:

   - Backend: `http://localhost:8080`
   - Frontend: `http://localhost:3000`

## Docker

To simplify deployment, use Docker Compose:

1. **Build and start containers**:

   ```bash
   docker-compose up --build
   ```

## API

### Endpoints

- **Shorten URL**: `POST /shorten`
  - Request: `{"long_url": "http://example.com"}`
  - Response: `{"short_url": "abC123"}`

- **Resolve URL**: `GET /:shortURL`
  - Redirects to the associated long URL.

## Contribution

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.# URL Shortener

URL Shortener is a web application that allows you to shorten long URLs into short, shareable links. The project consists of a backend developed in Go and a frontend using SvelteKit with Vite and Tailwind CSS.

## Features

- **URL Shortening**: Convert a long URL into a short URL.
- **Redirection**: Access the long URL using the short URL.
- **User Interface**: A simple and efficient interface to interact with the application.
- **URL Storage**: Uses a SQLite database to store URL mappings.

## Prerequisites

- Go 1.23 or later
- Node.js and npm (for the frontend)
- Docker and Docker Compose (for deployment)

## Installation

### Backend

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener/backend
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   - The project uses SQLite for URL storage. Ensure the `urls.db` file is accessible.

### Frontend

1. **Navigate to the frontend directory**:

   ```bash
   cd ../frontend
   ```

2. **Install dependencies**:

   ```bash
   npm install
   ```

3. **Start the development server**:

   ```bash
   npm run dev
   ```

## Usage

1. **Start the backend**:

   ```bash
   go run cmd/url-shortener/main.go
   ```

2. **Access the application**:

   - Backend: `http://localhost:8080`
   - Frontend: `http://localhost:3000`

## Docker

To simplify deployment, use Docker Compose:

1. **Build and start containers**:

   ```bash
   docker-compose up --build
   ```

## API

### Endpoints

- **Shorten URL**: `POST /shorten`
  - Request: `{"long_url": "http://example.com"}`
  - Response: `{"short_url": "abC123"}`

- **Resolve URL**: `GET /:shortURL`
  - Redirects to the associated long URL.

## Contribution

Contributions are welcome! Please open an issue or submit a pull request.