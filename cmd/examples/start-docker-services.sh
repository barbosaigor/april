#!/usr/bin/env bash

echo -e "Pulling lightweight image for testing\n"
docker pull alpine:latest

echo -e "Creating services containers\n"
services=(payment fees profile inventory shipping storefront)
for svc in ${services[@]}; do
	docker run -d --name $svc alpine:latest tail -f /dev/null
done
