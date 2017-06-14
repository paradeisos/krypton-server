godev:
	bee run

goclean:
	go clean -i ./...

goinstall:
	glide install

gotest:
	go test -p 1 krypton-server/controllers
	go test -p 1 krypton-server/models

gopackage:
	bee package

travis: gotest
