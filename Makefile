SHELL=/usr/bin/env bash


PLOTS?=./content.en/plots
PLOTDEFS?=./config/plotdefs


ASHBY?=./ashby
NEBULA_IPFS?=
PARSEC?=
TIROS?=

.PHONY: serve
serve:
	hugo server

.PHONY: ashby
ashby:
	go build ./code/cmd/ashby

.PHONY: build
build: ashby
	@${ASHBY} batch -s nebula_ipfs=${NEBULA_IPFS} -s parsec=${PARSEC} -s tiros=${TIROS} --version -vv --out ${PLOTS} --in ${PLOTDEFS}
