FROM golang:1.19
LABEL authors="moolcoov"

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o ./expressionist-orchestra

EXPOSE 8080
CMD ["./expressionist-orchestra"]
