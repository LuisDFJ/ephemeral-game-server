



# An Over-Engineered Ephemeral Game Server

This repository contains some automation scripts and instructions I gathered recently while attempting to build a Valheim game server for me and my friends. The requirements were as follows: I wanted an ephemeral server on DigitalOcean that I would only be charged for while it was in use, plus, optionally, the persistent disk space. For this I designed the following approach. First, I create a custom Linux image based on Arch Linux (for no specific reason), preconfigured with my default SSH hardening, fail2ban, a ufw firewall and, of course, the game server (run as a [Docker container](https://hub.docker.com/r/lloesche/valheim-server) with its corresponding systemd service configured). If your game server also uses a Docker container, this setup might work for you too, provided you make the appropriate firewall and persistent drive adjustments.

After creating the Linux image, I uploaded it to DigitalOcean and created a persistent Volume (around 10 GB, $1.00/month) that is mounted at boot and that the game’s Docker service is configured to use at `/mnt/gamedata/`. After playing, the droplet is completely destroyed so no extra charges apply. This setup has some clear advantages, like being able to scale the droplet up according to the RAM or CPU requirements of the moment (for instance, when hosting many users), and of course not paying for unused time (considering we only play a few hours on the weekend). On the other hand, it requires a lot more preparation and software infrastructure, and many steps could be automated further (this is a work in progress). Lastly, keeping a persistent IPv4 address in DigitalOcean unassigned to a droplet is costly (around $5/month) due to IPv4 scarcity, so to connect to the server one must ask for its IP at the start of each game session. This could also be solved by automating a DNS update API call on server creation, for as little as $5/year for a domain.

## Baking the Custom Linux Image

The instructions for baking the custom Linux image are documented in this [README](./linux-image-prebake/README.md), along with the configuration files. I am using QEMU x86 and the [Arch Linux cloud image](https://wiki.archlinux.org/title/Arch_Linux_on_a_VPS).

## Server Automation

There are plenty of ways to automate droplet creation for each game session. One can distribute a script or executable containing the access tokens for creating the droplet using DigitalOcean’s API (along with your DNS provider’s API, if you want a persistent domain for accessing the server). In my case, I have a spare VPS, also on DigitalOcean, that I use for hosting a blog. Since I’m already going down the road of over-engineering this game server, I decided to build my own authenticated API for creating the server. Each user can check the server status there and send a deploy command (with the respective DigitalOcean and DNS automation behind it). This has the advantage of not giving away precious API tokens, gives me control over a creation log, and allows safeguards such as tearing down the server if it goes unused.
