push:
	@go mod tidy
	@git add .
	@git commit -m "new thing"
	@git push
	@git tag "1.5"
	@ git push --tags