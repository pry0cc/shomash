#!/bin/bash

for domain in $(cat level_1.txt)
do
    echo "Scanning $domain..."
    mkdir -p output/$domain/
    python3 ~/tmp/domained/domained.py -d $domain --quick --noeyewitness 2>&1 > /dev/null
    cat output/$domain-all.txt | grep "$domain" | sort -u | zdns A | tee "output/$domain/dns.txt" | jq '.[].answers?[]?' | jq -r 'select(.type == "A") | .answer' | tee "output/$domain/ips_log.txt" | shomash -d "$domain"
    rm -f "output/$domain-*"
done