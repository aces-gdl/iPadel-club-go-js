#!/bin/sh
cd ui && npm run build && cd .. && go mod tidy &&  go run main.go && fg
