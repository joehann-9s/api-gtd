# GTD Task Management REST API  [![Version](https://img.shields.io/badge/Version-1.0-green.svg)](https://github.com/joehann-9s/api-gtd/releases/tag/v1.0) [![Go Version](https://img.shields.io/badge/Go-1.20-blue?logo=go)](https://golang.org/doc/go1.20) [![PostgreSQL Version](https://img.shields.io/badge/PostgreSQL-15-blue?logo=postgresql)](https://www.postgresql.org/docs/15/)


![Project slide](assets/slide-gtd.png)

A REST API for GTD-based task management, following the principles of David Allen's Getting Things Done (GTD) methodology. This API allows you to manage and organize your tasks using the GTD approach.


# Getting Started
## Installation
1. Clone this repository to your local machine.
    ```bash
    git clone https://github.com/joehann-9s/api-gtd.git
    ```

2. Run the following command to install the dependencies:
    ```bash
    cd api-gtd
    go version
    go mod download
    ```

3. Create and configure the `.env` file.
    ```bash
    touch .env
    ```
    after that, use this template:
    ```bash
    DSN ="host=postgres user=postgres password=password dbname=example-db port=5432 sslmode=disable TimeZone=America/Lima"
    JWT_SECRET="clave_secreta_del_token"
    PORT="3000"
    ```
    be sure to change these values with your own.

4. Start the server using the following command:
    - Development mode:
        ```bash
        air
        ```
    - Production mode:
        ```bash
        go build
        ./api-gtd
        ```

# API Documentation
You can find the API documentation on [postman](https://documenter.getpostman.com/view/27585203/2s93m612DE).




