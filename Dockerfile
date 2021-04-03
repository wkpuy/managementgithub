# Dockerfile References: https://docs.docker.com/engine/reference/builder/ 
#Create a multi-stage Docker image for GO

# Start from golang:1.12-alpine base image
FROM golang:1.13-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Warunee Kosab <kpnmtu@gmail.com>"

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container


# Build the Go app
RUN go build -o management cmd/api/main.go

# ========================== end building stage ==========================

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" > /etc/timezone

LABEL maintainer="Warunee Kosab <kpnmtu@gmail.com>"

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache \
    bash \
    openssh

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy app structure and other resource
COPY . .

# copy go app file
COPY --from=builder /app/management ./management

# Expose port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["./management"]