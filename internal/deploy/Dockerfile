FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN make build

CMD ["bash"]
