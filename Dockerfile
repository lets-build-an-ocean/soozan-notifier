FROM golang:alpine AS build

WORKDIR /app

COPY . .

# Build
RUN go build -o /go-sms


FROM alpine:latest
WORKDIR /app

COPY --from=build /go-sms .


EXPOSE 8080
CMD ["./go-sms"]