### Build
FROM golang:1.17.1-alpine as build

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

### Serve
FROM alpine:3.15.4
WORKDIR /app

ARG SERVER_PORT=4000
ARG SQL_ADDR
RUN test -n "$SQL_ADDR"

COPY --from=build /go/src/app/bin/server /app

RUN apk add --no-cache bash
RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh -O /usr/local/bin/wait-for-it \
    && chmod +x /usr/local/bin/wait-for-it

COPY ./cmd/server/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 4000

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/app/server"]