FROM golang:1.14 as build-image

WORKDIR /go/src
COPY domain/ ./domain/
COPY go.mod main.go ./

RUN go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["channel-to-slack"]
