FROM golang:1.24.4

WORKDIR /app

COPY . .

RUN chmod +x wait-for-it.sh

RUN go get -v

RUN go build -o main .

CMD ["./wait-for-it.sh", "postgresql_db:5432", "--", "./main"]
