# test-website

## Description

This project is for experimenting with creating a website using Go, React, and Postgresql.

## Installation

### Requirements

-   NodeJS
-   NPM
-   Go
    -   [echo](https://echo.labstack.com/)
    -   [pgx](https://github.com/jackc/pgx)
-   ReactJS
    -   [React-Bootstrap](https://react-bootstrap.github.io/)
    -   [Daemonite-Material](daemonite.github.io/material/) (CSS Only)
    -   [toastr](www.toastrjs.com)
-   Postgresql 12.3

### Steps

1. Install the npm packages

    ```
    npm install
    ```

1. Install the go modules

    ```
    go get github.com/labstack/echo/v4
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
    npm run start
    ```

-   Run the website ui development server
    ```
    npm run start-ui
    ```
