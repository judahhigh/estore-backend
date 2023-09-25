FROM golang:1.21
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ls -l
RUN go build -o main cmd/auth-server/auth-server.go
CMD ["/app/main"]