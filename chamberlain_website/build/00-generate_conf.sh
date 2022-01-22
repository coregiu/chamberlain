#!/bin/sh
echo "------generate the coregiu's configuration------"
export REAL_IP_VAR='$remote_addr'
envsubst < /etc/nginx/chamberlain.template > /etc/nginx/conf.d/default.conf