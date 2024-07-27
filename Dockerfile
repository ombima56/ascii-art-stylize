FROM ubuntu:latest
FROM golang:1.22.5

LABEL maintainer="ombimahillary6@gmail.com"
LABEL version="1.0"
LABEL environment="dev"

WORKDIR /ombima

COPY . .

RUN go build -o ascii-art-stylize .

EXPOSE 8080

VOLUME [ "home/downloads/doc" ]

CMD [ "./ascii-art-stylize" ]