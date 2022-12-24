files_src = src/server.go
REMOVE = rm -rf

.PHONY: all clean

all: backend.build

clean:
	echo cleaning backend build output...
	${REMOVE} build/*
	${REMOVE} backend.build
	echo finished cleaning backend build output

backend.build: ${files_src} go.mod go.sum
	echo building backend...
	go build -v -o build/ ${files_src}
	touch $@
	mkdir -p ../out
	cp -r -a -u build/. ../out/
	git -C .. submodule update --init
	mkdir -p ../out/kerncraft/examples
	cp -r -a -u ../dependencies/kerncraft/examples/. ../out/kerncraft/examples/
	echo finished building backend
