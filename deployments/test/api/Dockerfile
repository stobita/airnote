FROM golang:1.12.7-alpine3.10
WORKDIR /airnote
COPY . /airnote
RUN apk add --update --no-cache \
      git
CMD ash -c "go run cmd/airnote/airnote.go migrate && CGO_ENABLED=0 go test ./... -cover"
