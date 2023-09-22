#!/bin/bash

if pgrep -f "audio_server.py" > /dev/null
then
    pkill -f "audio_server.py"
    echo "stopped audio_server"
else
    echo "can't stop audio_server"
fi