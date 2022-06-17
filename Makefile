run:
	go run github.com/deyring/argos/cmd
test:
	go test github.com/deyring/argos/...
build-image:
	docker build -t deyring/argos:latest -f docker/Dockerfile .