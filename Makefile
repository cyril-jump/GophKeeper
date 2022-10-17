mock-gen:
	mockgen -source=internal/server/pkg/provider/interface.go -destination=internal/mocks/storage.go -package=mocks
test:
	go test ./...