#!/bin/bash

set -e
systemctl status --user mastodon-bot.timer

cd /home/eu/programming/mastodon-dailyintegral-bot/ || exit 1

/usr/sbin/python3 ./src/scrape/main.py
/usr/sbin/go run ./src/bot/main.go

exit
