#!/bin/bash

http_proxy=104.198.132.138:4140
for i in {1..1000}; do
	if [ $((i%5)) -eq 0 ]
	then 
		http_proxy=$http_proxy curl -s http://insult-now\?setTimeout\=long
	else
                http_proxy=$http_proxy curl -s http://insult-now
	fi
	
	http_proxy=$http_proxy curl -s http://insult-rude
done
