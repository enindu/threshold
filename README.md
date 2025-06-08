# Threshold

Threshold is a simple application based on [systemd](https://github.com/systemd/systemd) that automatically brings down a network interface after a specified amount of data has passed through it. Internally, it reads from the `/proc/net/dev` file to monitor network usage and uses the [netlink](https://github.com/vishvananda/netlink) package to manage network interfaces.
