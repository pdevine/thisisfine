FROM golang:alpine3.14
WORKDIR /project
COPY tif.go .
COPY go.* ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o thisisfine tif.go

FROM scratch
COPY --from=0 /project/thisisfine /thisisfine
ENV TERM="xterm-256color"
CMD ["/thisisfine"]
