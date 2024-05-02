FROM golang:alpine as app-builder
WORKDIR /go/tmp/app
COPY go.mod .
COPY go.sum .
COPY . .
COPY ./cmd/endpoint/main.go .
RUN ls -l
RUN go get .
RUN go list -m all
ENV TZ=Europe/Moscow
RUN CGO_ENABLED=0 go test -v ./...
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata -o main /go/tmp/app .

FROM alpine:latest
RUN apk --no-cache add tzdata
ENV TZ=Europe/Moscow
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app-builder /go/tmp/app/ /go/src/app/
COPY --from=app-builder /go/tmp/app/main /go/src/app/
WORKDIR /go/src/app/
CMD ["/go/src/app/main"]
