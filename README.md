# Notion HTMX Blog

A modern blog platform that uses Notion as a CMS, powered by Go and HTMX for dynamic frontend updates.

## Technologies

- Go - Backend server
- HTMX - Frontend interactivity
- Tailwind CSS - Styling
- Notion API - Content Management System

## Prerequisites

- Go 1.21 or higher
- Node.js and npm (for Tailwind CSS)
- Notion API key and database

## Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/notion-htmx-blog.git
```

2. Install dependencies:
```bash
go mod tidy 
```

3. Set up environment variables:
```bash
cp .env.example .env
```

4. Set up Notion API:
- Create a new Notion integration and get an API key.
- Create a new database in Notion and get the database ID.
- Set the database ID in the `main.go` file.

5. Run the server:
```bash
go run main.go
```

6. Access the blog at `http://localhost:8080`.


## Project Structure
```
.
├── cmd/
│   └── server/        # Main application entry point
├── internal/
│   ├── handlers/      # HTTP handlers
│   └── notion/        # Notion API integration
├── web/
│   ├── templates/     # HTML templates
│   └── static/        # Static assets
│       ├── css/       # Compiled CSS
│       └── js/        # JavaScript files
└── .env              # Environment variables
```

## Development

- Run `make dev`