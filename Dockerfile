# Specifies a parent image
FROM golang:1.23.5-alpine
 
# Creates an app directory to hold your app’s source code
WORKDIR /app

RUN go install github.com/air-verse/air@latest
 
# Copies everything from your root directory into /app
# COPY go.mod go.sum
COPY . .
 
# Installs Go dependencies
RUN go mod download

COPY . .
 
EXPOSE 3000
 
CMD [ "air", "-c", ".air.toml" ]
