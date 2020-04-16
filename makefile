all: clean lint test build

help:
	@echo "test build          : $$ make build"
	@echo "pack build          : $$ make mkzip"
	@echo "run test            : $$ make test"
	@echo "gofmt (over write)  : $$ make lint"
	@echo "flush tmp files     : $$ make clean"

build:
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./bin/darwin-amd64/shis ./cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./bin/linux-amd64/shis ./cmd/main.go

lint:
	@echo ""
	@echo "-- formatted files ( by go fmt ) --------"
	@go fmt ./cmd/... ./shis/... ./test/...
	@echo "-----------------------------------------"

.PHONY: test
test:
	@echo ""
	@echo "-- test start ---------------------------"
	@go test -v ./test/...
	@echo "-----------------------------------------"

clean:
	@rm -rf ./bin
	@rm -rf shis-*


current_tag_version := `git describe --tags --abbrev=0`
current_tag_major_version := `git describe --tags --abbrev=0 | tr -d 'v' | awk -F '.' '{print $$1}'`
current_tag_minor_version := `git describe --tags --abbrev=0 | tr -d 'v' | awk -F '.' '{print $$2}'`

latest_changelog_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 1 | awk -F ' ' '{print $$2}'`
latest_changelog_major_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 1 | awk -F ' ' '{print $$2}' | awk -F '.' '{print $$1}'`
latest_changelog_minor_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 1 | awk -F ' ' '{print $$2}' | awk -F '.' '{print $$2}'`

packed_changelog_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 2 | tail -n 1 | awk -F ' ' '{print $$2}'`
packed_changelog_major_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 2 | tail -n 1 | awk -F ' ' '{print $$2}' | awk -F '.' '{print $$1}'`
packed_changelog_minor_version := `cat CHANGELOG | grep -e VERSION -e Version -e version | head -n 2 | tail -n 1 | awk -F ' ' '{print $$2}' | awk -F '.' '{print $$2}'`

pack: clean lint test
	@echo ""

	@if [ $(current_tag_major_version) -ne $(packed_changelog_major_version) ] || [ $(current_tag_minor_version) -ne $(packed_changelog_minor_version) ]; then\
		echo "current tag version error";\
		echo "    current git tag version       : "${current_tag_version}; \
		echo "    CHANGELOG last packed version : v"${packed_changelog_version}; \
		exit 1;\
	fi

	@if [ $(current_tag_major_version) -gt $(latest_changelog_major_version) ] \
	|| ( [ $(current_tag_major_version) -eq $(latest_changelog_major_version) ] && [ $(current_tag_minor_version) -ge $(latest_changelog_minor_version) ] ); then\
		echo "current tag version error";\
		echo "    current git tag version  : "${current_tag_version}; \
		echo "    CHANGELOG latest version : v"${latest_changelog_version}; \
		exit 1;\
	fi

	@if [ ! -z "`git status --porcelain`" ]; then\
		echo "git branch is not clean";\
		exit 1;\
	fi

	@echo "last packed version is : v"$(packed_changelog_version)
	@echo "incremented version is : v"$(latest_changelog_version)
	@echo ""
	@echo "$$ git tag v$(latest_changelog_version)"
	@git tag v$(latest_changelog_version)
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.versionString=v$(latest_changelog_version) -w -s" -o ./shis-v$(latest_changelog_version)/darwin-amd64/shis ./cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X main.versionString=v$(latest_changelog_version) -w -s" -o ./shis-v$(latest_changelog_version)/linux-amd64/shis ./cmd/main.go
	@zip ./shis-v$(latest_changelog_version).zip ./shis-v$(latest_changelog_version)/darwin-amd64/shis ./shis-v$(latest_changelog_version)/linux-amd64/shis
	@rm -r ./shis-v$(latest_changelog_version)
