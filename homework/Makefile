
.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage.out
	go tool cover -func=coverage.out

clean_test_cache:
	go clean -testcache
	rm coverage/coverage.out
