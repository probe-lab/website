SHELL=/usr/bin/env bash


PLOTS?=./content.en/plots
PLOTDEFS?=./config/plotdefs
CONF?=./config

ASHBY?=./ashby
NEBULA_IPFS?=
PARSEC?=
TIROS?=

.PHONY: serve
serve:
	hugo server

.PHONY: build
build:
	@${ASHBY} batch -s nebula_ipfs=${NEBULA_IPFS} -s parsec=${PARSEC} -s tiros=${TIROS} --version -vv --out ${PLOTS} --in ${PLOTDEFS} --conf ${CONF}
