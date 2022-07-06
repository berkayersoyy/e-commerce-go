FROM golang:1.17-alpine

WORKDIR /app
ARG VERSION=dev

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /e-commerce-go -ldflags=-X=main.version=${VERSION} cmd/e-commerce-go/main.go

EXPOSE 8080

CMD [ "/e-commerce-go" ]