BUILD_PATH := ./bin/craig-weekend

.PHONY: build
build:
	go build -o ${BUILD_PATH}

.PHONY: clean
clean:
	rm -rf bin

.PHONY: run
run:
	${BUILD_PATH}

.PHONY: fresh
fresh: build run
