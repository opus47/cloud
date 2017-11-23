FROM golang:latest

MAINTAINER Ryan Goodfellow rcgoodfellow@gmail.com

WORKDIR /go/src/github.com/rcgoodfellow/opus47/api

# Right now this is a combination build+deploy container. In the future
# consider separating the two
RUN go-wrapper download \
  github.com/go-swagger/go-swagger/cmd/swagger \
  github.com/rs/cors

RUN go-wrapper install \
  github.com/go-swagger/go-swagger/cmd/swagger \
  github.com/rs/cors

RUN mkdir -p ~/.ssh

COPY . . 
RUN swagger generate server -f openapi.yml
RUN go-wrapper download ./...
RUN go-wrapper download ./cmd/opus47-server
RUN go-wrapper install ./cmd/opus47-server
RUN ./genkeys.sh

EXPOSE 443 

CMD ["./run.sh"]
