FROM golang:alpine as builder

WORKDIR /build

COPY . .

ENV GOPATH=/

RUN apk add build-base && apk add --no-cache bash && apk add sqlite && apk add sqlite-dev
RUN go env -w CGO_ENABLED=1 && go build -o todo ./cmd/main.go


FROM alpine:latest as production

COPY --from=builder /build/todo /build/todo

CMD [ "./build/todo" ]
