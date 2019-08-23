FROM golang:1.12.7
COPY . /airnote
WORKDIR /airnote
RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u -t github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
RUN go get github.com/cortesi/modd/cmd/modd
CMD ["modd", "-f", "configs/modd.conf"]
