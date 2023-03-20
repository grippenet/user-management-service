##########################
# STAGE 1
##########################
FROM golang:1.18-alpine as builder
ARG DIR=/go/src/github.com/influenzanet/user-management-service
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN mkdir -p "$DIR"
ENV GO111MODULE=on
ADD . $DIR
WORKDIR $DIR
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN mkdir bin

WORKDIR $DIR/tools/create-admin-user
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $DIR/bin/create-admin-user .

WORKDIR $DIR/tools/key-generator
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $DIR/bin/key-generator .

WORKDIR $DIR/tools/weekday-assign
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $DIR/bin/weekday-assign .

##########################
# STAGE 2
##########################
FROM alpine:latest
# we need those on scratch for connecting to mongoDB:
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/influenzanet/user-management-service/bin /app/
WORKDIR /app
CMD ["sh"]
