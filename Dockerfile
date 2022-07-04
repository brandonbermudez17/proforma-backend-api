FROM golang:1.18-alpine3.16

ENV APP_FILE_PATH cmd/main.go

RUN mkdir  /app

ADD . /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#COPY *.go ./

RUN go build -o /apiperfildata ${APP_FILE_PATH}
# RUN go build -o /backend-exe ${APP_FILE_PATH}
# RUN go build -o cmd/main .
# RUN go build -o cmd/main.go

EXPOSE 4000

# CMD [ "/main" ]
CMD [ "/apiperfildata" ]