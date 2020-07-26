FROM golang:1.14.4-alpine3.12 as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /app

COPY . .

RUN go build -o Photon_Server main.go

FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app && mkdir /log && \
  mkdir /web && mkdir /upload_files

WORKDIR /app 

EXPOSE 3000

COPY --from=builder /app/Photon_Server /app

CMD /app/Photon_Server
