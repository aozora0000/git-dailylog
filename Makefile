build:
	go build -o git-dailylog --ldflags="-s -w -X main.Version=`git describe --tag`" cmd/main.go

clean:
	rm -rf git-dailylog
	rm -rf dist

release:
	goreleaser --rm-dist

install: build
	sudo mv git-dailylog `which git-dailylog`
