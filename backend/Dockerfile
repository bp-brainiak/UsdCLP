FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/backend
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/backend 
#RUN cp $GOPATH/src/backend/.env.prod  go/bin/.env

FROM alpine:3.14
WORKDIR /go/bin/
COPY --from=builder /go/bin/backend /go/bin/backend
ADD prod_env .
ENTRYPOINT ["/go/bin/backend"]