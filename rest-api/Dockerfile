FROM golang:alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /rest-api

EXPOSE 8010

CMD [ "/rest-api" ]