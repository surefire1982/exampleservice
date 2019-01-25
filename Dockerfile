############################
# STEP 1 build executable binary
############################
FROM golang AS builder

COPY . $GOPATH/src/github.com/surefire1982/exampleservice/
WORKDIR $GOPATH/src/github.com/surefire1982/exampleservice/

# Fetch dependencies.
# Using go get.
RUN go get -d -v ./...
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/exampleservice api/main.go

COPY service.config.yaml /go/bin/service.config.yaml

###########################
# STEP 2 build a small image
###########################
FROM scratch

# Set DBHost to external mysql
ENV DBHOST="host.docker.internal"

WORKDIR /go/bin/

# Make port 80 available to the world outside this container
EXPOSE 8080
# Copy config settings 
COPY --from=builder /go/bin/service.config.yaml .

# Copy our static executable.
COPY --from=builder /go/bin/exampleservice .
# Run the hello binary.
ENTRYPOINT ["/go/bin/exampleservice"]