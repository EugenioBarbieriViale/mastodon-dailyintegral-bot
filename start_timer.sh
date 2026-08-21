#!/bin/bash

set -xe

systemctl --user daemon-reload
systemctl --user enable --now mastodon-bot.timer

chmod +x run.sh

exit
