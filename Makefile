VERSION_FILE := version.txt

push:
	go mod tidy
	git add .
	git commit -m "Update dependencies and code"
	git push
	chmod +x increment_version.sh
	$(eval NEW_VERSION := $(shell ./increment_version.sh $(VERSION_FILE)))
	git tag "$(NEW_VERSION)"
	git push origin --tags

# Optional clean target to remove generated files
clean:
	rm -f $(VERSION_FILE)

.PHONY: push clean