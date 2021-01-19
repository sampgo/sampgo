#!/bin/bash

git clone https://github.com/Zeex/sampgdk

# a_http.idl is cursed

find sampgdk -name "*.idl" | grep -v 'a_http' | xargs go run .
