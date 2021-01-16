#!/bin/bash

git clone https://github.com/Zeex/sampgdk

find sampgdk -name "*.idl" | xargs go run .
