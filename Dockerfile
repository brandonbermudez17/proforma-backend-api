FROM golang:1.18-alpine3.16

ENV APP_FILE_PATH cmd/main.go

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /backend-exe ${APP_FILE_PATH}

EXPOSE 8080

CMD [ "/backend-exe" ]