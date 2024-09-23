.PHONY: install run run-forced-seed clean

install:
	npm install

build:
	npm run build

run: build
	npm start

run-forced-seed:
	FORCED_SEED=true npm start

clean:
	rm -rf node_modules dist