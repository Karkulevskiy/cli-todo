FROM golang:alpine as builder

WORKDIR /build

COPY . .
ENV GOPATH=/
RUN go build -o todo ./cmd/main.go


FROM alpine:latest as production

COPY --from=builder /build/todo /build/todo

CMD [ "/build/todo" ]



