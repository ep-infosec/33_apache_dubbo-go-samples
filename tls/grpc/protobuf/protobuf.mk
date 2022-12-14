.PHONY: compile
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC := $(shell which protoc)
ifeq ($(PROTOC),)
	PROTOC = must-rebuild
endif

UNAME := $(shell uname)

$(PROTOC):
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install protobuf-compiler
endif

$(PROTOC_GEN_GO):
	go get -u dubbo.apache.org/dubbo-go/v3/protocol/grpc/protoc-gen-dubbo

helloworld.pb.go: helloworld.proto | $(PROTOC_GEN_GO) $(PROTOC)
	protoc -I . helloworld.proto --dubbo_out=plugins=grpc+dubbo:.

.PHONY: compile
compile: helloworld.pb.go

