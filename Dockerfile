FROM golang:1.14.2-alpine as builder
MAINTAINER "Gilson Gabriel <gilson@codemastersolucoes.com>"

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-w -s" -installsuffix cgo -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]