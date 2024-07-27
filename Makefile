push:
	go mod tidy
	git add .
	git commit -m "Update dependencies and code"
	git push origin main
	git tag "v1.5"
	git push origin --tags