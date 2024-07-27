push:
	@go mod tidy
	@git add .
	@git commit -m "new thing"
	@git push
	@git tag "1.4"
	@ git push --tags