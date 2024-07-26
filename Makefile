push:
	@go mod tidy
	@git add .
	@git commit -m "new thing"
	@git push