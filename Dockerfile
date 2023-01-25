FROM golang:1.19 as dev

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o parkinglot-service ./src/

FROM dev as debug

#RUN go get  github.com/go-delve/delve/cmd/dlv@v1.7.3
#RUN go get  github.com/cespare/reflex@v0.3.1

RUN go get -d -v ./...

ENTRYPOINT [ "/go/app/parkinglot-service" ]
#CMD reflex -R "__debug_bin" -s -- bash -c "dlv debug --headless --continue --accept-multiclient --listen :40000 --api-version=2 --log ./src/"

FROM dev as build
