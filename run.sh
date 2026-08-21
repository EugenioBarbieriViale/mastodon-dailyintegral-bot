#!/bin/bash

set -e

path=/home/eu/programming/mastodon-dailyintegral-bot/
bot_binary=main

cd "$path" || exit 1

/usr/sbin/python3 ./src/scrape/main.py

go build -o "$bot_binary" ./src/bot/main.go
./"$bot_binary"
