#!/bin/bash

for domain in $(cat level_1.txt); do mkdir -p output/$domain/; timeout 5 crobat-client -s $domain | grep "$domain"; done | zdns A | tee output/$domain/dns.txt | jq -r 'select(.data.answers[0].type == "A") | .data.answers[0].answer' | tee output/$domain/ips_log.txt | go run main.go -d $domain
