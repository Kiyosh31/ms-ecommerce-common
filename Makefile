push:
	@go mod tidy
	@git add .
	@git commit -m "new thing"
	@git tag v1.3
	@git push