# Organizr

## Description

Building an organizational app to learn and experiment with Go and React.

## Installation

### Requirements

-   NodeJS
-   NPM
-   Go
    -   [echo](https://echo.labstack.com/)
    -   [pgx](https://github.com/jackc/pgx)
-   ReactJS
    -   [Semantic UI](https://react.semantic-ui.com/)
-   Postgresql 12.7

### Steps

1. Install the npm packages

    ```
    npm install
    ```

1. Install the go modules

    ```
    go get github.com/labstack/echo/v4
    go get github.com/labstack/gommon/log@v0.3.0
    go get github.com/jackc/pgx/v4
    go get golang.org/x/crypto/bcrypt
    ```

1. Set the `PGHOST` and `PG_URL` environment variables

    ```
    PGHOST: <db_host>
    PG_URL: user=<db_user> password=<db_user_pwd> host=<db_host> port=<db_port> dbname=<db_name>
    ```

1. Set the `JWT_SECRET` environment variable
    ```
    JWT_SECRET: <secret>
    ```

## Building & Running

### Building

-   Build both the website ui and server executable

    ```
    npm run build
    ```

-   Build website ui separately

    ```
    npm run build-ui
    ```

-   Build server executable separately
    ```
    npm run build-server
    ```

### Running

-   Run the server executable

    ```
    npm run start-server
    ```

-   Run the website ui development server
    ```
    npm run start-ui
    ```
