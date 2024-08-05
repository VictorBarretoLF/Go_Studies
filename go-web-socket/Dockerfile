FROM golang:1.22.2

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra-cli@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 \
    go install github.com/jackc/tern/v2@latest \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

CMD ["tail", "-f", "/dev/null"]
