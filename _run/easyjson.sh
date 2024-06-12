#!/bin/sh


echo "Обход структуры Time-Rules"
easyjson -all -omit_empty struct.go
easyjson -all -omit_empty def.go

