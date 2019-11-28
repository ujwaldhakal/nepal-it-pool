#!/bin/bash
if [ ${APP_ENV} = dev ]; \
then \
go run migration/developer.go; \
fi

if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
