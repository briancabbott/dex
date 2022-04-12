cd api/v2
rm go.sum
go mod download
go mod verify
go mod tidy
cd ../..