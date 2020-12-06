# Help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: imports
imports: ## Run code cleanup with goimports
	@# Clean files that are not inside the vendor/ folder (faster!)
	@for FILENAME in $$(find . -type f -name '*.go' -not -path "./vendor/*"); do \
		echo "goimports -w $$FILENAME"; \
		goimports -w $$FILENAME; \
	done

.PHONY: test
test: ## Run unit tests
	@go test -v ./...
