# todo: We should version all of these dependencies.

FROM golang:1.23

# Base libraries
RUN apt-get update
RUN apt-get install -y protobuf-compiler

# Go protoc dependencies
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

# Common base proto dependencies
RUN go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
RUN git clone --branch v1.1.0 https://github.com/envoyproxy/protoc-gen-validate
RUN git clone https://github.com/googleapis/googleapis

RUN mkdir /proto-common
RUN cp -r protoc-gen-validate/validate /proto-common/validate
RUN cp -r googleapis/google /proto-common

COPY generate.sh generate.sh
RUN chmod +x generate.sh

ENTRYPOINT [ "./generate.sh" ]
