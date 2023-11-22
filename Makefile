build_go:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o build/main cmd/main.go
deploy_dev: build
	serverless deploy --aws-profile profile_uala_global_labsupport_dev

clean-cache:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache

dependencies:
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/onsi/gomega/...

deps:dependencies
	@go mod tidy

cover:deps
	@go test ./... -coverprofile=c.out.tmp -coverpkg=./... && cat c.out.tmp | grep -v "_mock.go" > c.out

report:cover
	@go tool cover -func c.out | grep "total"

html-report:cover
	@go tool cover -html c.out

test:deps
	@go test ./...