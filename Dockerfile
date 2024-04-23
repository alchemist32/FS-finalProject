# syntax=docker/dockerfile:1

FROM golang:1.21 as build

WORKDIR /code
ADD ./products-api ./

RUN go mod tidy
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 9000

# Run
CMD ["/docker-gs-ping"]


FROM build
ARG SERVER_PORT
ENV SERVER_PORT=$SERVER_PORT