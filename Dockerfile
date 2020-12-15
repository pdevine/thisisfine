FROM golang:alpine3.12
WORKDIR /project
COPY tif.go .
RUN apk add --no-cache git
RUN go get github.com/pdevine/go-asciisprite
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o thisisfine tif.go

FROM scratch
COPY --from=0 /project/thisisfine /thisisfine
ENV TERM="xterm-256color"
CMD ["/thisisfine"]
