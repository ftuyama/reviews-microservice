NAME = weaveworksdemos/reviews
DBNAME = weaveworksdemos/reviews-db
INSTANCE = reviews
TESTDB = weaveworkstestreviewsdb
OPENAPI = $(INSTANCE)-testopenapi
GROUP = weaveworksdemos

default: docker


pre:
	go get -v github.com/Masterminds/glide

deps: pre
	glide install

rm-deps:
	rm -rf vendor

test:
	@docker build -t $(INSTANCE)-test -f ./Dockerfile-test .
	@docker run --rm -it $(INSTANCE)-test /bin/sh -c 'glide novendor| xargs go test -v'

cover:
	@glide novendor|xargs go test -v -covermode=count

coverprofile:
	go get github.com/modocache/gover
	go test -v -covermode=count -coverprofile=profile.coverprofile
	go test -v -covermode=count -coverprofile=db.coverprofile ./db
	go test -v -covermode=count -coverprofile=mongo.coverprofile ./db/mongodb
	go test -v -covermode=count -coverprofile=api.coverprofile ./api
	go test -v -covermode=count -coverprofile=reviewss.coverprofile ./reviewss
	gover
	mv gover.coverprofile cover.profile
	rm *.coverprofile


dockerdev:
	docker build -t $(INSTANCE)-dev .

dockertestdb:
	docker build -t $(TESTDB) -f docker/reviews-db/Dockerfile docker/reviews-db/

dockerruntest: dockertestdb dockerdev
	docker run -d --name my$(TESTDB) -h my$(TESTDB) $(TESTDB)
	docker run -d --name $(INSTANCE)-dev -p 8084:8084 --link my$(TESTDB) -e MONGO_HOST="my$(TESTDB):27017" $(INSTANCE)-dev

docker:
	docker build -t $(NAME) -f docker/reviews/Dockerfile-release .

dockerlocal:
	docker build -t $(INSTANCE)-local -f docker/reviews/Dockerfile-release .

cleandocker:
	-docker rm -f my$(TESTDB)
	-docker rm -f $(INSTANCE)-dev
	-docker rm -f $(OPENAPI)
	-docker rm -f reviews-mock

clean: cleandocker
	rm -rf bin
	rm -rf docker/reviews/bin
	rm -rf vendor
