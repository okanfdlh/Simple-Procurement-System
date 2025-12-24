FROM golang:1.25-alpine

WORKDIR /app

# install git (sometimes needed for go mod)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build binary
RUN go build -o app

EXPOSE 3000

CMD ["./app"]
