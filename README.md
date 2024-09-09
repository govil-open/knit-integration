# knit-integration

Aim of this project is to provide a base code structure for developing any new microservice. You only need to focus on actual business logic implementation as all other boiler plate code implementation is taken care of.

## Getting Started

Follow below steps to get the boiler plate code running on your local machine.

### Prerequisites

- Golang
- Postgres
- Docker

### Local setup

Install swag by swaggo (when writing a new API, we use swag to generate the swagger documentation)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Run a postgres server using docker with below command

```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
```

Login to the postgres database and create a Database with name 'go_boiler_plate_db'

```bash
docker exec -it postgres psql -U postgres
psql > CREATE DATABASE go_boiler_plate_db;
```

Open terminal in your project root directory and run below command

```bash
make server
```

## Project Overview

- .github/workflows - This is for github quality checks. There are total three workflows.

    1. Audit - Any code pushed will go through a set of quality checks like lint checks and tests execution. Refer .golanci.yml for the list of enabled linters.
    2. Static Analysis - This will run the db migrations, tests and generate a coverage report. Then it runs Sonar analysis for quality gate checks. Some the checks include code coverage > 80% etc.
    3. API Documentation - This is manual process. Go to GitHub -> Actions -> API Documentation -> Click on 'Run workflow' drop down -> select branch -> run workflow. This will generate the api documentation. [Link](https://api-doc.open.money/go-boiler-plate.html) to documentation.

- config - This package contains scripts for configuration loading, database initialilzation and migration, logger and monitoring setup. Viper is used for configuration loading, logrus for logging, sentry for application monitoring. locale.go is the script for loading configuration related to localisation. Any error shown to the user is localised. (Resources->messages->en.json)

- db - This package contains scripts for database migration, querying and mocking. sqlc is the library used for sql injection. You just need to write the .sql scripts for all the required database changes and use sqlc to generate the corresponding go scripts.
    - To generate migration schema files run 'make migrate FILE_NAME={}' from the root folder of the application. This will generate empty .sql files in the migration folder. Update the sql files to include your migration sql scripts.
    - Include all your query sqls in the query package and then run 'make sqlc' command to generate the corresponding .go files.
- env - All the environment variables are in this package.

- middleware - This package contains the scripts for all the activites required before request reaches the controller or after sending the response. This project has middleware for health check before the request reaches controller. Another middleware is for error handling.

- model - This package contains structure (data transfer objects) for api request and response.

- server - This package contains scripts for initialization of the server, setting up routes, adding validators.

- test - This package contains all the functional tests.

- main.go - This is the entry point of the application which calls all the initialization scripts and starts the http server.

- Makefile - This file defines the set of tasks to be executed.

## Makefile

Run the application

```bash
make server
```

Generate scripts from sql

```bash
make sqlc
```

Generate sql file in migrations folder

```bash
make FILE_NAME={} migrate
```

Run only migration upward

```bash
make DB_URL={} migrateup
```

Run rollback migration

```bash
make DB_URL={} migratedown
```

Run linters

```bash
make lint
```

Run sonar analysis

```bash
make SONAR_HOST={} SONAR_LOGIN={} GIT_COMMIT_ID={} PROJECT_KEY={} sonar
```

Run all tests

```bash
make test
```

Mock database

```bash
make mock
```

Generate Swagger documentation

```bash
make swagger
```

Format Swagger documentation

```bash
make swagger-format
```
