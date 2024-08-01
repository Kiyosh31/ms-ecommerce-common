push:
	go mod tidy
	git add .
	git commit -m "Update dependencies and code"
	git push
	$(eval LATEST_TAG := $(shell git describe --tags `git rev-list --tags --max-count=1`))
	$(eval MAJOR := $(shell echo $(LATEST_TAG) | cut -d. -f1 | sed 's/v//'))
	$(eval MINOR := $(shell echo $(LATEST_TAG) | cut -d. -f2))
	$(eval PATCH := $(shell echo $(LATEST_TAG) | cut -d. -f3))
	$(eval NEW_PATCH := $(shell echo $$(($(PATCH) + 1))))
	$(eval NEW_TAG := v$(MAJOR).$(MINOR).$(NEW_PATCH))
	git tag $(NEW_TAG)
	git push origin --tags

.PHONY: push