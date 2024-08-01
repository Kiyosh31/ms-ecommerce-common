push:
	go mod tidy
	git add .
	git commit -m "Update dependencies and code"
	git push
	git tag "v1.16"
	git push origin --tags