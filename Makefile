NAME := github.com/kare/environment

test:
	go test -v $(NAME)

compile:
	go build $(NAME)

install:
	go install $(NAME)

vet:
	go vet -n -x $(NAME)

fmt:
	go fmt $(NAME)

dep-install:
	go get github.com/dmotylev/goproperties

clean:
	go clean $(NAME)
