godev:
	bee run

goclean:
	go clean -i ./...

goinstall:
	glide install

gotest:
	go test -p 1 ./controllers
	go test -p 1 ./models

gopackage:
	bee package

travis: gotest
