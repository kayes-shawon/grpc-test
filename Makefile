make-proto:
	@protoc \
			-I=./proto \
				--go_opt=module=github.com/kayes-shawon/grpc-test \
				--go_out=. \
				--go-grpc_opt=module=github.com/kayes-shawon/grpc-test \
				--go-grpc_out=. \
				./proto/hello.proto
