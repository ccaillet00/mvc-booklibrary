# Stage Build
FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY . .

RUN go build -o /out/booklibrary .

# Stage run

FROM scratch

WORKDIR /data

COPY --from=build /out/booklibrary /app/booklibrary

# Default to HTTP mode; can be overridden via APP_MODE env var (cli, http, both)
ENV APP_MODE=http

EXPOSE 8080

CMD [ "/app/booklibrary" ]