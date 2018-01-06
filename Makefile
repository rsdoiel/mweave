#
# Simple Make for supporting compilation to Linux, Mac OS X and Windows
# on amd64, Raspbian on ARM7 and ARM6 and Linux on ARM64
#

PROJECT = mweave

MOTTO = "An experimental literate programming tool"

VERSION = $(shell grep -m 1 'Version =' mweave.go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

CLI_NAMES = mweave

build: $(CLI_NAMES) README.md

mweave: bin/mweave$(EXT)

bin/mweave$(EXT): mweave.go cmd/mweave/mweave.go
	env CGO_ENABLED=0 go build -o bin/mweave$(EXT) cmd/mweave/mweave.go

README.md: bin/mweave$(EXT) README.mweave
	./bin/mweave -weave -i README.mweave -o README.md
	git add README.md

test:
	go test

install:
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmd/mweave/mweave.go

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d node_modules ]; then rm -fR node_modules; fi


generate_usage_pages: mweave
	bash gen-usage-pages.bash

website: generate_usage_pages
	bash mk-website.bash

publish: generate_usage_pages
	bash mk-website.bash
	bash publish.bash

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/mweave cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/mweave.exe cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/mweave cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mweave cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/raspbian-arm6:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/bin/mweave cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm6.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 GOARM=6 go build -o dist/bin/mweave cmd/mweave/mweave.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md docs/* bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist/docs
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	bash gen-usage-pages.bash
	cp -v docs/mweave.md dist/docs/

release: generate_usage_pages distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/raspbian-arm6 dist/linux-arm64

