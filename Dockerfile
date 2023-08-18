FROM golang:1.17-alpine3.15

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 9000

CMD [ "go", "run", "main.go" ]