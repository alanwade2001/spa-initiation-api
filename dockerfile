FROM golang:alpine AS build

COPY go.mod go.sum /go/src/app/

WORKDIR /go/src/app/
RUN go mod download

COPY . . 
RUN go install 
RUN ls -ltr /go/bin/

FROM alpine

COPY --from=build /go/bin/spa-initiation-api /bin
COPY app.env /

#RUN ls -ltr /
CMD ["spa-initiation-api"]