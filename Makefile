push:
	go mod tidy
	git add .
	git commit -m "Update dependencies and code"
	git push