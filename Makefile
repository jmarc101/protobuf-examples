# Change this to your own Go module
GO_MODULE := grpc-go


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc \
    -I . -I google/type \
    --go_opt=module=grpc-go \
    --go_out=. \
    ./proto/basic/*.proto \
    ./proto/basic/first/*.proto \
    ./proto/basic/second/*.proto

.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run


.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
