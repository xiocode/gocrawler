#!/bin/bash
remote=git@github.com:xiocode/gocrawler.git
git pull $remote master
git add .
git commit -m $1
git push $remote

