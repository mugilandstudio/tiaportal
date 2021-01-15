# serverをコンパイルする
.PHONY: build
build:
	go build -o bin/tia cmd/tia.go

# protocを実行する。
.PHONY: proto
proto:
	protoc -I./proto/ \
		-I./third_party/googleapis/ \
		--go_out=./proto/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto/ \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./proto/ \
     	--grpc-gateway_opt=logtostderr=true \
     	--grpc-gateway_opt=paths=source_relative \
     	--grpc-gateway_opt=generate_unbound_methods=true \
		--js_out=import_style=commonjs:./proto \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./proto \
		./proto/*.proto 