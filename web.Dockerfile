FROM golang:1.14.6

ENV APP_PATH="/seatlect-service"

WORKDIR ${APP_PATH}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o web ./cmd/web/main.go
CMD ./web
