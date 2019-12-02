#!/bin/sh

trap 'kill $(jobs -p)' EXIT

# Start the first process
nginx
status=$?
if [ $status -ne 0 ]; then
	echo "Failed to start nginx: $status"
	exit $status
fi

# Start the second process
sudo -u weekend ./backend &
status=$?
if [ $status -ne 0 ]; then
	echo "Failed to start backend: $status"
	exit $status
fi

# Naive check runs checks once a minute to see if either of the processes exited.
# This illustrates part of the heavy lifting you need to do if you want to run
# more than one service in a container. The container exits with an error
# if it detects that either of the processes has exited.
# Otherwise it loops forever, waking up every 60 seconds

while sleep 60; do
	ps aux > /tmp/ps.txt
	grep -q nginx /tmp/ps.txt
	NGINX_STATUS=$?
	grep -q backend /tmp/ps.txt
	BACKEND_STATUS=$?
	if [ $NGINX_STATUS -ne 0 ]; then
		echo "Error: nginx exited"
		exit 1
	fi
	if [ $BACKEND_STATUS -ne 0 ]; then
		echo "Error: backend exited"
		exit 1
	fi
done
