FROM golang:1.12.6-alpine3.10

WORKDIR /i_like_money

COPY web.go web.go

RUN go build

CMD [ "./i_like_money" ]
