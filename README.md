# Estore Backend

**Estore Backend** is a demo microservice backend for an e-commerce store.

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

How to use
----------
1. Use docker to build the project, and all microservices.

```
docker compose build
```

2. Use docker to spin up all microservices.

```
docker compose up
```