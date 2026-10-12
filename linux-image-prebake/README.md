# Valheim Server VPS Image Baking

Create a seed file:

```sh
cloud-localds seed.img user_data.yaml
```

We derive an image from Arch Linux Cloud Image:

```sh
qemu-img create \
    -f qcow2 \
    -F qcow2 \
    -b ../Arch-Linux-x86_64-cloudimg.qcow2 \
    bake.qcow2 10G
```

And run the image:

```sh
qemu-system-x86_64 \
    -enable-kvm \
    -cpu host \
    -smp 4 \
    -m 4G \
    -drive file=bake.qcow2,format=qcow2 \
    -drive file=seed.img,format=raw \
    -nic user,hostfwd=tcp::2222-:22,hostfwd=udp::2456-:2456,hostfwd=udp::2457-:2457 \
    -nographic \
```


SCP baking files:

```sh
scp -P 2222 prepare.sh bake.sh datasource.cfg 50-wired-dhcp.network sshd-hardening.conf valheim.service <username>@localhost:~/
```

The SSH into the host and bake and run `prepare.sh`, reboot, and `bake.sh`.

Check VPS, `bake.sh` should already take care of most of the files:

```sh
cloud-init --version
# Check for data sources order, set datasource_list: [ ..., NoCloud, None ]
grep -rh datasource /etc/cloud/cloud.cfg /etc/cloud/cloud.cfg.d/

# Should be DHCP on en*/eth*
ls /etc/systemd/network; networkctl

# Virtio Modules should be in initramfs. If not add to MODULES=(virtio_blk,virtio_pci,virtio_scsi,virtio_net) and run mkinitcpio -P
lsinitcpio /boot/initramfs-linux.img | grep -E 'virtio_(blk|scsi|net|pci)'
```

Clean VPS Image:

```sh
cloud-init clean --logs --seed
rm -f /etc/ssh/ssh_host_*
truncate -s 0 /etc/machine-id
pacman -Scc --noconfirm
shutdown -h now
```

Finally

```sh
qemu-img convert -O qcow2 -c bake.qcow2 valheim-arch-do.qcow2
```
