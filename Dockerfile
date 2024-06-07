FROM huecker.io/library/golang:alpine as BUILDER

WORKDIR /build

COPY . .
ENV GOPATH=/
RUN go build -o tasks ./cmd/main.go


FROM huecker.io/library/alpine:latest as production

COPY --from=builder /build/tasks /build/tasks


