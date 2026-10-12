#!/bin/sh

# Error Exit + No Unset + Pipefail
set -euo pipefail
pacman -Syu --noconfirm docker ufw fail2ban vim openssh kitty-terminfo
