.DEFAULT_GOAL := help

ifeq ($(TAG),)
	TAG := "latest"
endif

.PHONY: help
help: ## Справка по командам
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build 
build: ## Сборка
	GO111MODULE=off go build -o awt main.go

.PHONY: docker 
docker: ## Сборка docker image
	docker build -t yuccastream/awt:$(TAG) .

.PHONY: push 
push: ## Push docker image to registy
	docker push yuccastream/awt:$(TAG)
