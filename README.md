# Estore Backend

**Estore Backend** is a demo microservice backend for an e-commerce store.

## Table of Contents
- [Motivation](#motivation)
- [Technology](#technology)
- [ProjectStructure](#project-structure)
- [HowToUse](#how-to-use)

## Motivation

This project started as a way to showcase language proficiency in Golang. Being that
Go is particularly well suited for server-side projects, it made sense to create
a demo project where the backend is written mostly, if not entirely, in Go. The reason
an e-commerce store was chosen as the example application was because of the incredible
number of aspects a backend would have to consider to support it. Stores like Amazon
have to serve millions of consumers, each of which require speed, security, and efficiency
for product search, transactions, reviews, etc. As a vendor, the platform must also provide
a means to connect consumers to their products with elegance and efficiency. Furthermore,
vendor revenue through transactions must be well maintained and secure. In essence, as a 
demo project, developing an estore backend is nothing short of a challenge.

## Technology

An e-commerce store, especially one that handles traffic at scale, can and should employ
a number of technologies on the backend to best ensure the success of the store. Below is
a list of many technologies planned for or already employed in the demo project.

* **Go or Golang**: This language enables efficient server-side programming with benchmarks approaching some of the fastest languages like C, C++, and Rust, but without all the syntactical bells and whistles that make programming in languages like C, C++, and Rust more difficult. Along with compile-time protections, Go's design lends itself to readable server-side code that tackles complex problems. Furthermore, community support for the language and it's server-side applications is broad.
* **Docker**: Many server-side processes can, and should, be containerized into small chunks that can be deployed within on-prem or cloud-based infrastructure to handle business functions at scale. In this project, business functions are broken apart into bite-sized microservices that can be containerized in Docker and deployed anywhere.
* **RabbitMQ**: Regardless of the number and size of microservices, it is advantageous to decouple service communication from the microservices themselves. RabbitMQ provides a means of connecting microservices together through message consumers and producers. An on-prem system could have microservices publish messages to pre-defined channels, which RabbitMQ would route to other microservices to effectuate changes in state, side-effects, et cetera. By abstracting the communication between microservices away from the microservices themselves, microservice maintainers only have to worry about where to publish their messages to and what these messages should contain. No longer do microservice maintainers have to seek and and connect to microservices directly through HTTP, websocket, or any other means.
* **GoKit**: GoKit is an all-in-one microservice development library for Go. It decouples business logic, from endpoints, and services allowing all kinds of flexibility in developing microservices.
* **Kubernetes**: Kubernetes allows microservices to be deployed to handle scaling elegantly on backend load. This pod orchestration tool is an absolute necessity for deploying and configuring microservices in the most cost-effective and clean manner possible.
* **PostgreSQL**: As an open-source scalable, ACID compliant, relational database technology, PostgreSQL is a great choice to develop a enterprise db. For the e-commerce store project I'm building, this technology enables relational table management of account information, preferences, stores, products, and more. I could have well made the db SQLite,a NoSQL variant, or even a simple in-memory map, but PostgreSQL is good enough and can eventually be hosted in an cloud environment like AWS RDS later if desirable.
* **Redis**: The plan is to create a token store to issue and track all access tokens required by the client to access database resources through the microservice architecture. Redis provides much faster response times relative to non-in-memory database alternatives. As such, it provides a means to store and deliver time-critical data to enable larger performance in systems that necissitate it. In this case, the redis database will simply handle tokens, but as such it shouldn't bog down the system for such a small scale use-case.

## Project structure

This project comprises several microservice servers that together bring life to an e-commerce store. The listing below
explains the various microservices implemented.

* **api/account**: This microservice is a CRUD REST api for interacting with a postgreSQL db. Clients can manage user-account related information with this microservice.
* **auth/account**: This microservice handles higher-level functions for user registration, login, and token refresh. Under the hood, this microservice relies on the api/account microservice for all user-management related activites. This service also relies on the auth/token microservice to issue tokens on registration and login to enable user-related operations from a frontend for the e-commerce store.
* **auth/token**: This microservice handles token creation, storage, and distribution. It exposes service endpoints
that allow clients to issue new tokens, refresh tokens, and revoke tokens.

## How to use

Currently the project relies on docker compose to initialize and start all microservices within the project. While the frontend has not been implemented, a tool like Postman or Curl can be used to trigger microservice endpoints once started.

For security, microservices are configured with godotenv to read sensitive information from the environment. To start up the project, a root-level .env file with the following variables must be created.

```
DB_HOST="database"                     # The host-name set on the account-api server for reaching the db
DB_PORT="5432"                         # The port exposing the postgreSQL db
DB_USER="postgres"                     # The username configured for postgreSQL
DB_PASSWORD="postgres_password"        # The password for the configured postgreSQL user
DB_NAME="backend_db_name"              # The postgreSQL database to store user information
SERVICE_PORT="8080"                    # The port the account-api server is reachable at
SECRETKEY="some_secret_key"            # A secret key to decode jwt tokens for microservice endpoint access
ACCOUNT_API_HOST="account-api-server"  # The name of the account-api server used by the auth-server to access
ACCOUNT_API_SCHEME="http"              # The scheme for accessing the account-api 
ACCOUNT_API_PORT="8080"                # The port to target by the auth-server to reach the account-api
ACCOUNT_SERVER_PORT="8081"             # The port the auth-server is reachable at
```

After setting up the above .env file and after initializing a postgreSQL database for the account-api to use, build the
project using docker compose and start it.

1. Use docker to build the project, and all microservices.

```
docker compose build
```

2. Use docker to spin up all microservices.

```
docker compose up
```

At this point microservices should be running and can be reachable from properly configured requests using a tool
like Postman or Curl.