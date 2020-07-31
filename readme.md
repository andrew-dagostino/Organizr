# test-website

## Description

This project is for experimenting with creating a website using Go, React, and Postgresql.

## Installation

### Requirements

* NodeJS
* NPM
* Go
    * [echo](https://echo.labstack.com/)
* ReactJS
    * [Material-UI](https://material-ui.com/)
* Postgresql - todo

### Steps

1. Install the npm packages
    ```
    npm install
    ```

1. Install the go modules
    ```
    go get github.com/labstack/echo/v4
    ```
1. postgresql setup - todo

## Building & Running

### Building

* Build both the website ui and server executable
    ```
    npm run build
    ```

* Build website ui separately
    ```
    npm run build-ui
    ```

* Build server executable separately
    ```
    npm run build-server
    ```

### Running

* Run the server executable
    ```
    npm run start
    ```

* Run the website ui development server
    ```
    npm run start-ui
    ```