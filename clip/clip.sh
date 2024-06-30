#!/bin/sh

xdotool key Control+Insert; ./clip; xdotool key Shift+Insert; xkb-switch -n;
