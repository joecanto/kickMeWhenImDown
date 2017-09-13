#!/bin/bash

http_proxy=104.198.132.138:4140
for i in {1..1000}; do
	http_proxy=$http_proxy curl -s http://insult-rude
	http_proxy=$http_proxy curl -s http://insult-now
done
