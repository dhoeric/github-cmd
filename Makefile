.PHONY: docker build push
ORG := dhoeric
PROJECT := github-cmd
TAG := latest
IMG_NAME := $(ORG)/$(PROJECT):$(TAG)

docker: build push

build:
	docker build -t $(IMG_NAME) .

push:
	docker push $(IMG_NAME)

