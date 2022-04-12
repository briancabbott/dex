export CGO_ENABLED=0

go mod tidy -compat=1.17
go mod tidy

make bin/kind
make bin/protoc-gen-go
make bin/protoc
make bin/gotestsum
make bin/grpc-client

make build
make buildmake testall
make testall
make generate
make build

make bin/dex

make kind

make bin/example-app

make deps
make proto-internal
make proto
make lint


make release-binary
make release-binary
make release-binary
