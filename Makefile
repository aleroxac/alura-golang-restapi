## ---------- UTILS
.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean all temp files
	@rm -rf .temp



## ---------- SETUP
.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest
	cd src && go mod tidy



## ---------- MAIN
.PHONY: build
build:
	@[ -d .build ] && rm -rf .build || true
	@[ -d src/tmp ] && rm -rf src/tmp || true
	@mkdir .build && cp -r src/* .build
	@docker build -f docker/Dockerfile -t alura-golang .build
	@rm -rf .build



.PHONY: run
run:
	@cd src && air

.PHONY: up
up:
	@docker-compose -f docker/docker-compose.yaml up -d

.PHONY: down
down:
	@docker-compose -f docker/docker-compose.yaml down
	@pkill air
