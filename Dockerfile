FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Nam Nguyen <nam.nguyen.de@gmail.com>"

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app


RUN go get -u github.com/golang/dep/...

# Install the package
RUN dep ensure

# Build my app
RUN go build -o /app/main .

CMD ["/app/main"]
