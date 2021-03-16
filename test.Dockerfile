FROM golang:1.14.6 AS dev

ENV APP_PATH="/seatlect-service"

WORKDIR ${APP_PATH}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
CMD go test -v ./...