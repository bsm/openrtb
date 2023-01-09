default: test

.minimal.makefile:
	curl -fsSL -o $@ https://gitlab.com/bsm/misc/raw/master/make/go/minimal.makefile

include .minimal.makefile
