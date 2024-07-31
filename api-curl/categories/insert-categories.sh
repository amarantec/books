#!/usr/bin/bash

echo "Informe o nome da categoria:"
read name
echo "Informe o url da categoria:"
read url

curl -i -X POST -H "Content-Type: application/json" http://localhost:8080/insert-category -d "{\"name\": \"$name\", \"url\": \"$url\"}"
