FROM golang:1.18 as builder

ARG DOCKER_GIT_CREDENTIALS

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY . ./
COPY ./cmd/ /app

# install git
RUN apt-get update
RUN apt-get install -y git

RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server
RUN ls /app
RUN ls /app/server

## Frontend React
FROM node:alpine as fb 

#Specify a working directory
WORKDIR /app
COPY ./web/ /app
#Copy the dependencies file
RUN npm install
#Copy remaining files
COPY . .
#Build the project for production
RUN npm run build 


#Copy production build files from builder phase to nginx

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
# RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
# COPY --from=builder /app/server /server

COPY --from=fb /app/build /web
RUN ls 
RUN ls /web

# Run the web service on container startup.
CMD ["/web"]