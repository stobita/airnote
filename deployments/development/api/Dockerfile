FROM golang:1.12.7
COPY . /airnote
WORKDIR /airnote
RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u -t github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
CMD ["realize", "start", "--run", "--server"] 
