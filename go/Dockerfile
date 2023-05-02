FROM golang:1.20.3-alpine3.17
RUN apk --no-cache --update add npm make gcc g++ musl-dev openssl-dev git perl-utils curl
WORKDIR /go/src/sso-auth-server
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/.
CMD [ "./app" ]