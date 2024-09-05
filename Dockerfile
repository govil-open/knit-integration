# Build stage
FROM golang:1.23-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20
ENV USER=docker
ENV UID=12345
ENV GID=23456
RUN addgroup -S "$USER" && \
    adduser --disabled-password --gecos "" --home "/app" --ingroup "$USER" --no-create-home --uid "$UID" "$USER"
WORKDIR /app
RUN chmod a+w /app
COPY --from=builder /app/main .
COPY env/* env/
COPY resources/messages/* resources/messages/
COPY db/migration/* db/migration/
EXPOSE 8080
USER docker
CMD [ "/app/main" ]