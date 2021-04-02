# build proto files
build_proto:
	protoc --go_out=plugins=grpc:proto ./proto/messages.proto

# clean the cache
go_clean:
	-go clean -cache 

# clean the proto files
clean_generated:
	rm -f .generated
	find . -name "*pb*.go" -delete
