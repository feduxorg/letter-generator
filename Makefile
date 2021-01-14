build:
.PHONY: all
all: build

.PHONY: build
build:
	bin/build

.PHONY: test
test:
ifeq ($(DEBUG), 1)
	  bin/test -d
else
	  bin/test
endif


.PHONY: clean
clean:
	bin/clean
