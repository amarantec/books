#!/bin/sh
curl -X POST -H "Content-Type: application/json" http://192.168.2.22:8080/user-signup -d '{"name": "test2", "email": "test2@test.com", "password": "test2"}' 
