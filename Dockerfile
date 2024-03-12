# syntax=docker/dockerfile:1

FROM golang:1.20

# set and use /app as destination for COPY
WORKDIR /purreads

# copy env file
COPY .env.docker ./

COPY go.mod go.sum ./
COPY gqlgen.yml ./
COPY database ./database/
COPY graph ./graph/

# copy all go source files
COPY *.go ./

# download go modules
# RUN go get github.com/newrelic/go-agent/v3/newrelic
RUN go get github.com/newrelic/go-agent/v3/newrelic
RUN go mod download
COPY . .

# compile app
RUN CGO_ENABLED=0 GOOS=linux go build -o /start-app

EXPOSE 8080

# run app
CMD ["/start-app"]
