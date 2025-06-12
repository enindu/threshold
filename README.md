# Threshold

Threshold is a simple application based on [systemd](https://github.com/systemd/systemd) that automatically brings down a network interface after a specified amount of data has passed through it. Internally, it reads from the `/proc/net/dev` file to monitor network usage and uses the [netlink](https://github.com/vishvananda/netlink) package to manage network interfaces.

## Install

Since this application requires root privileges, you’ll need to install Threshold as the root user and configure the environment accordingly. Assuming you are already logged in as root, add the following lines to the `/root/.bash_profile` file:

```bash
export PATH="$PATH:$HOME/.go/bin"
export GOPATH="$HOME/.go"
```

After updating the file, apply the changes by running:

```
source /root/.bash_profile
```

You can now install Threshold using the `go install` command:

```
go install github.com/enindu/threshold@latest
```

## Usage

Here is how to use Threshold:

```
threshold <command>:<subcommand> [arguments]
```

For example, to start the daemon, run:

```
threshold daemon:start eth0 1024 1
```

To view a list of available commands:

```
threshold help
```

To see the subcommands for a specific command:

```
threshold <command>:help
```

For instance, to view help for the `daemon` command:

```
threshold daemon:help
```

## License

This software is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). You can view the full license [here](https://github.com/enindu/threshold/blob/master/COPYING).
