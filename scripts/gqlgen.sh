#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f graph/generated/generated.go \
    graph/model/models_gen.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"