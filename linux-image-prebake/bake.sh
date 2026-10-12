#!/bin/sh

# Error Exit + No Unset + Pipefail
set -euo pipefail

systemctl enable --now docker
docker pull lloesche/valheim-server

# Hardening
# SSH Hardening Config
install -m 0644 sshd-hardening.conf /etc/ssh/sshd_config.d/00-hardening.conf

# SSHD Jail
printf '[sshd]\nenabled = true\nbackend = systemd\n' > /etc/fail2ban/jail.local

# UFW Config
ufw default deny incoming
ufw default allow outgoing
ufw allow 2456:2457/udp
ufw limit 22/tcp
ufw --force enable

# Service Enable
install -m 0644 valheim.service /etc/systemd/system/valheim.service
systemctl enable sshd ufw fail2ban docker valheim.service

# Condigure Data Source
install -m 0644 datasource.cfg /etc/cloud/cloud.cfg.d/99-datasource.cfg
rm -f /etc/systemd/network/*.network
install -m 0644 50-wired-dhcp.cfg /etc/systemd/network/
