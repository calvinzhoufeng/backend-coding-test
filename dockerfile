FROM golang:buster AS buildenv

# The local workdir is used
ENV SRC_DIR=/go/src/firemark
RUN mkdir $SRC_DIR
COPY go.mod $SRC_DIR
WORKDIR $SRC_DIR

COPY . $SRC_DIR

RUN go mod download

# build binary
RUN go build \
    -o /firemark \
    ./src/cmd


# build working image
FROM golang:buster

COPY --from=buildenv /firemark /app/firemark

WORKDIR /app

ENTRYPOINT ["/app/firemark"]
