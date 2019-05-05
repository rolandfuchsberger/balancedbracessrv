# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Source: https://www.callicoder.com/docker-golang-image-container-example/

# Start from golang v1.11 base image
FROM golang:1.11 as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/fuchsberger.email/balancedbracessrv/

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install packr2 utility
RUN go get -u github.com/gobuffalo/packr/v2/packr2

# Run packr2 to generate .go files out of boxes (not working with v1!)
RUN packr2

# Build the Go app with https://github.com/gobuffalo/packr (to include templates)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/balancedbracessrv .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/balancedbracessrv . 

EXPOSE 8080

CMD ["./balancedbracessrv"] 
