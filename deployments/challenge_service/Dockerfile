FROM golang:1.20-alpine3.16 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ARG BUILD_REF
WORKDIR /service
COPY . .

#
RUN go build cmd/main.go


# Run the Go Binary in Alpine.
FROM alpine:3.17
COPY --from=build /service /service
WORKDIR /service
CMD ["./main"]
