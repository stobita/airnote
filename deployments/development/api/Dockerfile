FROM golang:1.12.7-alpine3.10
RUN apk add --update --no-cache \
      git
COPY . /airnote
WORKDIR /airnote
RUN go get -u github.com/oxequa/realize
CMD ["realize", "start", "--run", "--server"] 
