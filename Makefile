basedir=$(PWD)

all:
	go install $(basedir)/...

build:
	go build $(basedir)/...

test:
	go test $(basedir)/...

clean:
	go clean
