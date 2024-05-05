NAME = weaveworksdemos/reviews
DBNAME = weaveworksdemos/reviews-db

TAG=$(TRAVIS_COMMIT)

INSTANCE = reviews

.PHONY: default copy test

default: test

release:
	docker build -t $(NAME) -f ./docker/reviews/Dockerfile .

test:
	GROUP=weaveworksdemos COMMIT=test ./scripts/build.sh
	./test/test.sh unit.py
	./test/test.sh container.py --tag $(TAG)

dockertravisbuild: build
	docker build -t $(NAME):$(TAG) -f docker/reviews/Dockerfile-release docker/reviews/
	docker build -t $(DBNAME):$(TAG) -f docker/reviews-db/Dockerfile docker/reviews-db/
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	scripts/push.sh
