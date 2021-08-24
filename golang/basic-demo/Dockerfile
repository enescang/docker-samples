FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /hello-docker

ENV PORT=3000
EXPOSE 3000

CMD [ "/hello-docker" ]