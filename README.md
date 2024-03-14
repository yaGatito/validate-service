# validate-service

Technical task for COMPANY.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Installation
1. Download & install [Docker](https://www.docker.com/get-started/) 
2. Clone the git repository or download a .zip file.
3. Proceed to the root of the project folder.
4. Using Terminal run next 2 commands sequentially
    ```shell
    docker build -t validate-service .
    docker run -e VALIDATE_SERVICE_PORT=8081 -p 8080:8081 --rm validate-service
    ```

## Usage
Fork my postman [collection](https://elements.getpostman.com/redirect?entityId=20487409-5f38dab1-0cc5-41d8-9eb5-c0770ec5088a&entityType=collection) to be able to use prepared collection
### Schema:
HTTP Method: GET
URL: /validate
Query params: number, year, month
```
ex: http://localhost:8080/validate?number=123123123123&year=2026&month=2
```

## Configuration
Specify `-e VALIDATE_SERVICE_PORT=[INTERNAL-port]` and `-p [OUTER-port]:[INTERNAL-port]` flags on `docker run` to configure internal and outer ports of the container.
