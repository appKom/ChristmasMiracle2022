
FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /cm_backend

EXPOSE 8080

CMD ["/cm_backend"]
