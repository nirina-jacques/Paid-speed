FROM golang:1.24-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl
WORKDIR /usr/app
COPY . .
RUN go get ./ &&\
    go build &&\
    CGO_ENABLED=0 go build -a -installsuffix cgo -o /app
FROM scratch
WORKDIR /
COPY --from=builder /app /app
EXPOSE 9191
CMD [ "/app" ]