FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/second-kill ./cmd/main.go

FROM kasmweb/hunchly:develop-rolling

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/second-kill /app/second-kill

ENTRYPOINT [ "./second-kill" ]
CMD [" "]