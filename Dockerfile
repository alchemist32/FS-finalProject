# syntax=docker/dockerfile:1

FROM golang:1.21 as build

WORKDIR /code
# ADD ./products-api ./
COPY ./products-api/go.mod ./products-api/go.sum ./

RUN go mod download
COPY ./products-api/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 9000

# Run
CMD ["/docker-gs-ping"]