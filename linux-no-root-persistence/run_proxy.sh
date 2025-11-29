#!/bin/bash

PROXY_BINARY="./proxy-linux"
LOG_FILE="./proxy.log"

echo "Starting persistent SOCKS5 proxy..." >> $LOG_FILE

while true
do
    echo "Attempting to run $PROXY_BINARY at $(date)" >> $LOG_FILE

    # Run the proxy, redirecting output to the log file
    $PROXY_BINARY >> $LOG_FILE 2>&1

    # $? holds the exit status of the last command (i.e., the proxy)
    EXIT_CODE=$?

    echo "Proxy crashed with exit code $EXIT_CODE at $(date). Restarting in 5 seconds..." >> $LOG_FILE

    # Wait a few seconds before restarting to avoid a rapid restart loop
    sleep 10 
done
