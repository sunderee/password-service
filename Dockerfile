# ==============
# | FIRST STAGE |
# ==============
FROM golang:latest AS builder

WORKDIR /server

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# ===============
# | SECOND STAGE |
# ===============
FROM alpine:latest
RUN apk upgrade --no-cache -U && apk add --no-cache ca-certificates

RUN addgroup -S dev && adduser -S dev -G dev
USER dev
RUN mkdir /home/dev/server
WORKDIR /home/dev/server

# ============
# | FINISH UP |
# ============
COPY --from=builder /server/main .
EXPOSE 3333

CMD ["./main"] 