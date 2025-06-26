FROM golang:alpine

WORKDIR /app

COPY . .

#RUN go build
#RUN go install github.com/air-verse/air@latest
#RUN air init

#CMD ["air"]

RUN go build -o typotemplate

CMD ["./typotemplate"]
