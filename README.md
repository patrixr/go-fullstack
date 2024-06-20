# go-fullstack

## About

`go-fullstack` is a full-stack setup that leverages Pocketbase, Tailwind CSS, htmx, Alpine.js, and Golang. This project includes a ready-to-go Docker Compose setup with a Cloudflare tunnel for easy deployment.

1. [Features](#features)
2. [Project Structure](#project-structure)
3. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Dependencies](#dependencies)
   - [Getting Started Guide](#getting-started-guide)
   - [Development Workflow](#development-workflow)
       - [Live Reload](#live-reload)
       - [Compile Tailwind CSS](#compile-tailwind-css)
       - [Generate templ files](#generate-templ-files)
       - [Docker](#docker)
4. [Data Migrations](#data-migrations)
5. [Contributing](#contributing)

## Features

- **Pocketbase**: Handles the backend and database.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **htmx**: Enhances HTML with AJAX, CSS Transitions, WebSockets, and Server-Sent Events.
- **Alpine.js**: Lightweight JavaScript framework for handling UI components.
- **Golang**: Powers the backend server.
- **Docker Compose**: Simplifies deployment and environment configuration (optional).
- **Cloudflare Tunnel**: Securely exposes your local server to the internet (configured in Docker compose)
- **Air**: Provides live reload for development.

## Project Structure

- **Pages**: Located in `internal/views/pages`.
- **Components**: Located in `internal/views/components`.
- **Fragments**: Located in `internal/views/fragments`. Fragments are small pieces of HTML that can be returned for usage with htmx.
- **Layouts**: Located in `internal/views/layouts`.
- **Router Configuration**: Configure routes in `internal/core/router`.

## Getting started

### Prerequisites

- Golang
- Tailwind CSS (can be installed as a standalone binary or via Node.js)
- Docker (optional)
- Docker Compose (optional)

### Dependencies

- **Tailwind CSS**: Used for styling. [Tailwind Installation Guide](https://tailwindcss.com/docs/installation)
  - Can be installed via npm/yarn or as a standalone binary.
- **Templ**: A Golang templating language that compiles to Go functions. [Templ Guide](https://templ.guide)
- **Air**: Used for live reloading during development.

### Getting Started

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/yourusername/go-fullstack.git
    cd go-fullstack
    ```

2. **Install Dependencies**:
    ```sh
    make tidy
    ```

3. **Run Migrations**:
    ```sh
    make migrate
    ```

4. **Build the Project**:
    ```sh
    make build
    ```

5. **Run the Application**:
    ```sh
    make run
    ```

### Development Workflow

To streamline your development workflow, you can use the following commands:

#### Live Reload

Air provides live reloading during development. To start the development server with live reload, you will have to install Air first:

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

and then run:

```sh
make dev
```

#### Compile Tailwind CSS

```sh
make tailwind
```

#### Generate templ files

```sh
make templ
```

#### Docker

To deploy the application, use Docker Compose which is configured to set the ENV variable to prod by default.

```sh
docker-compose up -d
```

## Data migrations

The project supports "auto migrations" from Pocketbase if the `ENV` environment variable is set to `dev`. This feature allows you to modify records in the Pocketbase admin and automatically generate migration files in `internal/migrations`.

**Note**: The `ENV` variable defaults to `dev`. Ensure to set it to `prod` when releasing. Docker Compose does this for you by default.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.
