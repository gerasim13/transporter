FROM golang:1.4.2
MAINTAINER Pavel Litvinenko <gerasim13@gmail.com>

RUN go-wrapper download github.com/compose/transporter/...
RUN go-wrapper install github.com/compose/transporter/...