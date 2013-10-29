NAME := github.com/kare/environment

compile:
	go build $(NAME)

install:
	go install $(NAME)

test:
	go test -v $(NAME)

vet:
	go vet -n -x $(NAME)

fmt:
	go fmt $(NAME)

dep-install:
	go get github.com/dmotylev/goproperties
