FROM golang:1.18

COPY go.mod /code/go.mod
COPY go.sum /code/go.sum
WORKDIR /code
RUN ls
RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon
COPY . /code

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
