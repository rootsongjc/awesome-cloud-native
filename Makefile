.PHONY: build preview test status clean help

# ─── Configuration ───
PORT     ?= 8080
README   := README.md
TMPL_DIR := tmpl
SCRIPTS  := scripts

# ─── Targets ───

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

build: ## Generate tmpl/index.html from README.md
	go run repo.go

preview: build ## Build and start local preview server (port configurable: make preview PORT=3000)
	@echo "Previewing at http://localhost:$(PORT)"
	@cd $(TMPL_DIR) && python3 -m http.server $(PORT)

test: ## Run Go tests (alphabetical order + duplicate link checks)
	go test ./...

status: ## Run project status checker (reads GITHUB_TOKEN from .env)
	@if [ -f .env ]; then set -a && . ./.env && set +a && python3 $(SCRIPTS)/check_status.py; else python3 $(SCRIPTS)/check_status.py; fi
	@if [ -f status.json ]; then cp status.json $(TMPL_DIR)/status.json; echo "Copied status.json to $(TMPL_DIR)/"; fi

clean: ## Remove generated files
	rm -f $(TMPL_DIR)/index.html status.json
