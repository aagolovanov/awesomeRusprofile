
clean:
	cd ./pkg && rm -f *.pb.go


protogen: clean
	cd api; protoc --go_out=../pkg --go_opt=paths=source_relative \
                --go-grpc_out=../pkg --go-grpc_opt=paths=source_relative \
                *.proto

run:
	go run ./cmd/main.go