package userdata

import (
	"fmt"
)

type Params struct {
	User string
	SSHKey string
	Volume string
	WorldName string
	WorldPswd string
}

func (p *Params) Gen() string {
	return fmt.Sprintf(USER_DATA,
		p.User,
		p.SSHKey,
		p.Volume,
		p.WorldName,
		p.WorldPswd,
		p.Volume,
	)
}

const USER_DATA string = `
#cloud-config
hostname: valheim-arch-vps

users:
  - name: %s
    groups: [wheel]
    shell: /bin/bash
    sudo: ALL=(ALL) NOPASSWD:ALL
    ssh_authorized_keys:
      - %s
mounts:
  - [ /dev/disk/by-id/scsi-0DO_Volume_%s, /mnt/gamedata, ext4, "defaults,nofail,discard", "0", "2" ]

write_files:
  - path: /etc/valheim-server.env
    permissions: '0600'
    content: |
      SERVER_NAME=Valheim Server
      SERVER_PORT=2456
      WORLD_NAME=%s
      SERVER_PASS=%s
      SERVER_PUBLIC=false

runcmd:
  - resize2fs /dev/disk/by-id/scsi-0DO_Volume_%s
  - systemctl enable --now sshd ufw fail2ban
  - mkdir -p /mnt/gamedata/config /mnt/gamedata/server
  - systemctl start valheim.service
`
