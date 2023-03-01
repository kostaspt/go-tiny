### Build
FROM golang:1.19.1-alpine3.16 as build

LABEL org.opencontainers.image.source="REPOSITORY_URL"

WORKDIR /go/src/app

# Install system dependencies
RUN apk add --no-cache make git

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Make sure we do have version available
ARG VERSION
ENV VERSION=$VERSION

# Generate production build
COPY . .
RUN make build

### Deploy
FROM alpine:3.17.2
WORKDIR /app

# Define and verify args
ARG SQL_ADDR
RUN test -n "$SQL_ADDR"
ENV SQL_ADDR=$SQL_ADDR

COPY --from=build /go/src/app/bin /app

RUN apk add --no-cache bash
RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh -O /usr/local/bin/wait-for-it \
    && chmod +x /usr/local/bin/wait-for-it

COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]