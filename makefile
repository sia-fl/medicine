target  = app/cli/main.go

proto:
	protoc --go_out=./ \
	--go_opt=Mdist/proto/comm.proto=src/proto/ \
	--go_opt=Mdist/proto/ids.proto=src/proto/ \
	dist/proto/ids.proto \
	dist/proto/comm.proto;
	protoc --go-grpc_out=./ \
	--go-grpc_opt=Mdist/proto/comm.proto=src/proto/ \
	--go-grpc_opt=Mdist/proto/ids.proto=src/proto/ \
	dist/proto/comm.proto \
	dist/proto/ids.proto;

build:
	go build -o dist/commodity.exe $(target)
	$(go_exec) go build -o dist/commodity $(target)

.PHONY:proto test