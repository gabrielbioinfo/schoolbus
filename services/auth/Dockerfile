FROM golang:1.18-alpine as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
# COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app/ ./cmd/server/main.go

FROM scratch
COPY --from=build-env /go/bin/app /app
EXPOSE 8080
ENTRYPOINT ["/app/main"]