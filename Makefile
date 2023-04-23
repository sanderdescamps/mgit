BINARY_NAME=mgit

env-clean:
	unset $(env | grep -i ^mgit | awk -F '=' '{print $$1}')


BINARY_NAME=mgit

_build:
	@mkdir -p ./dist
	@echo "=> build ${BINARY_NAME}"
	@GOARCH=amd64 GOOS=linux go build -o ./dist/${BINARY_NAME}-linux-amd64 main.go

_install:
	@echo "=> install ${BINARY_NAME}"
	@cp ./dist/${BINARY_NAME}-linux-amd64 ~/.local/bin/mgit
	@chmod +x ~/.local/bin/mgit

_auto_complete:
	@echo "=> install autocomplete"
	@mgit completion bash > /tmp/mgit-completion
	@echo "install autocomplete requires root permissions"
	@sudo cp /tmp/mgit-completion /etc/bash_completion.d/mgit-completion

build: _build
install: _build _install _auto_complete


