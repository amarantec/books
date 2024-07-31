#!/usr/bin/bash

echo "Informe seu email"
read email

echo "Informe sua senha"
read password

curl -i -X POST -H "Content-Type: application/json" http://localhost:8080/user-login -d "{\"email\": \"$email\", \"password\": \"$password\"}"
