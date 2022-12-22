files_root = package-lock.json package.json svelte.config.js tsconfig.json vite.config.js
files_src = src/app.d.ts src/app.html src/routes/+layout.js src/routes/+page.svelte
files_static = static/favicon.png
REMOVE = rm -rf

.PHONY: all clean

all: frontend.build

clean:
	echo cleaning frontend build output...
	${REMOVE} build/*
	${REMOVE} frontend.build
	echo finished cleaning frontend build output

frontend.build: ${files_root} ${files_static} ${files_src}
	echo building frontend...
	npm run build
	touch $@
	echo finished building frontend

