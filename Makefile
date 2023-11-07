build:
	env GOOS=linux go build -ldflags="-s -W" -o build/main main.go
deploy_prod: build
	serverless deploy --stage dev --aws-profile default