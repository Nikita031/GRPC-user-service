FROM golang:1.22.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /grpc-user-service

EXPOSE 50051

CMD [ "/grpc-user-service" ]
