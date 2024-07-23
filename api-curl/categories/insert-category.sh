#!/bin/sh

curl -X POST -H "Content-Type: application/json" http://192.168.2.22:8080/insert-category -d '{"name": "test1", "url": "test1"}'
