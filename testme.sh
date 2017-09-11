#!/bin/bash

for i in {1..1000}; do
	http_proxy=104.198.132.138:4140 curl -s http://hello
done
