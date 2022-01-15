#!/bin/sh
echo "------generate the coregiu's configuration------"
envsubst < /etc/nginx/chamberlain.template > /etc/nginx/conf.d/default.conf