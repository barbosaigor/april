#!/usr/bin/env bash

echo -e "Pulling lightweight image for testing\n"
docker pull alpine:latest

echo -e "Creating services containers\n"
services=(payment fees profile inventory shipping storefront)
for svc in ${services[@]}; do
	docker run -d --name $svc alpine:latest tail -f /dev/null
done

echo -e "Creating services containers for selector test\n"
services=(123-payment-321 asd-feesadsasd 23123profile346356 asdasinventoryd asdasdadsshipping storefrontadsasdvas)
for svc in ${services[@]}; do
	docker run -d --name $svc alpine:latest tail -f /dev/null
done
