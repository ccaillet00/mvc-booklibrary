# Stage Build
FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY *.go ./

RUN go build -o /out/booklibrary

# Stage run

FROM scratch

WORKDIR /data

COPY --from=build /out/booklibrary /app/booklibrary 

CMD [ "/app/booklibrary" ]