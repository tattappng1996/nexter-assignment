run:
	go run main.go

findfail:
	go test ./... | grep FAIL

test:
	go test ./... -v

mockgen:
	mockgen -package=mocks -source=usecase/usecase.go -destination=tests/mocks/usecase_mock.go

coverage:
	go test ./... -coverprofile=./tests/coverage.out
	go tool cover -html=./tests/coverage.out