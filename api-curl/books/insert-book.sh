#!/bin/sh
curl -X POST -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQHRlc3QuY29tIiwiZXhwaXJlIjoxNzIxNzA4OTY4LCJ1c2VySWQiOjJ9.O_k2A5_rApo9zZeSThjtmBPzrnmAIQFyhncHUEArwfU" http://192.168.2.22:8080/insert-book -d '{"title": "Test4", "description": "Testando4", "genre": ["test4 gen"], "author": ["test4 author"], "category_id": 1}'

