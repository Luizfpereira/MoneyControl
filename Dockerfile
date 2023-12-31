FROM golang:1.20.3

WORKDIR /go/src
RUN ln -sf /bin/bash /bin/sh

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]



