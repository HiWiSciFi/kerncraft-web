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
	cp -a build/. ../out/
	echo finished building backend
