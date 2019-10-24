FROM golang
# RUN apk add git gcc
RUN mkdir app
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
COPY . /app
WORKDIR /app
CMD ["go", "run", "server.go"]

EXPOSE 8000