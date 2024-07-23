#!/bin/sh
curl -X POST -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQHRlc3QuY29tIiwiZXhwaXJlIjoxNzIxNzA5NzEzLCJ1c2VySWQiOjJ9.nCf4-Wbu5dfLs2oAHZ7XCMD30PLE4e835VV8dQjhcHE" http://192.168.2.22:8080/insert-book -d '{"title": "Test5", "description": "Testando5", "genre": ["test5 gen"], "author": ["test5 author"], "category_id": 1}'

