##
## Build
##
FROM golang:1.18-alpine AS build

WORKDIR /app

RUN apk update && apk add git
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /app/go-docker

##
## Deploy
##
FROM golang:1.18-alpine

WORKDIR /app

COPY --from=build /app/go-docker /app

EXPOSE 3000

CMD ["/app/go-docker"]