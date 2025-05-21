FROM golang:1.24.3

WORKDIR /app

COPY . .

EXPOSE 8000

RUN go build -o main

CMD [ "./main" ]