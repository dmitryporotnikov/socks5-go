# Portable SOCKS5 proxy, no root

[![Go 1.26](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![No root](https://img.shields.io/badge/root-not%20required-green)](#)
[![Linux](https://img.shields.io/badge/platform-linux-lightgrey)](linux-no-root-persistence)

A SOCKS5 proxy written in one Go file, for the VPS where you have a login and no sudo. Build the binary on your own machine, copy it across, and it runs out of your home directory. No packages to install on the server, no service manager to fight.

Credentials are compiled into the binary. That's the whole configuration story, so treat this as a personal tunnel and not as shared infrastructure.

## Build

Open `main.go`, set the two constants near the top:

```go
const requiredUser = "dmitry"
const requiredPass = "<SECRET_PASSWORD>"
```

The listener is `0.0.0.0:40048`; change `listenAddr` if that port is taken. Build for the server from any OS:

```bash
GOOS=linux GOARCH=amd64 go build -o proxy-linux main.go
```

## Put it on the server

```bash
ssh user@your-server-ip "mkdir -p ~/proxy"
scp proxy-linux user@your-server-ip:~/proxy/
scp linux-no-root-persistence/run_proxy.sh user@your-server-ip:~/proxy/
```

Then, on the server:

```bash
cd ~/proxy
chmod +x proxy-linux run_proxy.sh
nohup ./run_proxy.sh > /dev/null 2>&1 &
```

`run_proxy.sh` keeps the proxy alive: if the binary exits, it logs the code and starts it again after 10 seconds.

## Start it on boot

`crontab -e`, then add a line with the path from `pwd`:

```cron
@reboot /bin/sleep 300 && /usr/bin/nohup /home/<username>/proxy/run_proxy.sh > /dev/null 2>&1 &
```

The 300 second sleep gives the network time to come up before the proxy starts.

## Check it

```bash
ps aux | grep proxy-linux
tail -f ~/proxy/proxy.log
```

Point the client at the server's public IP, port `40048`, with the username and password from `main.go`.

## Two things to know

The password sits in the binary as a string constant, so anyone who gets a copy of the file can pull it out. The proxy also adds no encryption of its own: whatever the client puts through it is protected only by the protocol being carried.

There's no license file in this repo yet, so treat the code as all rights reserved.
