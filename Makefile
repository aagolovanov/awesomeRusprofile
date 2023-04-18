
clean:
	cd ./pkg && rm -f *.pb.go


protogen: clean
	cd api; protoc -I google/api --go_out=../pkg --go_opt=paths=source_relative \
                --go-grpc_out=../pkg --go-grpc_opt=paths=source_relative \
                api.proto

run:
	go run ./cmd/main.go

testFunc:
	go test ./pkg



gen-gateway:
	cd api; protoc -I . -I google/api --grpc-gateway_out ../pkg \
        	--grpc-gateway_opt logtostderr=true \
        	--grpc-gateway_opt paths=source_relative \
        	--grpc-gateway_opt generate_unbound_methods=true \
        	api.proto


gen-swagger:
	cd api; protoc -I . --openapiv2_out ./openapiv2 \
                --openapiv2_opt logtostderr=true \
                api.proto