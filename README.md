# Portable SOCKS5 Proxy (No Root Required)

This project provides a lightweight, portable SOCKS5 proxy server written in Go. It is designed specifically for **VPS or servers where you do not have root (administrator) access**.

Since installing software packages on restricted servers can be tricky or impossible without root privileges, this solution compiles into a single binary file that you can just copy and run. It includes a helper script and cron job instructions to ensure the proxy stays running and restarts automatically if the server reboots.

## Features

- **No Root Access Needed**: runs entirely in user space.
- **Portable**: Compiled as a single static binary.
- **Persistent**: Includes a shell script to auto-restart the proxy if it crashes.
- **Auto-Start**: Instructions for `cron` to start the proxy on server reboot.
- **Hardcoded Security**: Username and password are compiled into the binary for simplicity (no config files to manage).

---

## 1. Configuration

Before compiling, you need to set your desired username and password directly in the source code.

1. Open `main.go` in a text editor.
2. Look for the Credentials section:
   ```go
   const requiredUser = "dmitry"
   const requiredPass = "<SECRET_PASSWORD>"
   ```
3. Change `"dmitry"` and `"<SECRET_PASSWORD>"` to your preferred username and password.
4. (Optional) The server listens on port `40048` by default. You can change `listenAddr` in `main.go` if you need a different port.

---

## 2. Compilation

You need to compile the Go program into a binary that can run on Linux. You can do this from your local machine (Windows, Mac, or Linux).

**Prerequisite**: Make sure you have [Go installed](https://go.dev/dl/).

Open your terminal or command prompt in the project directory and run:

### For Linux/Mac Users:
```bash
GOOS=linux GOARCH=amd64 go build -o proxy-linux main.go
```

### For Windows Users (PowerShell):
```powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o proxy-linux main.go
```

This will create a file named `proxy-linux`. This is your executable program.

---

## 3. Installation on the Server

Now you need to move the files to your remote server. We will assume you are using a directory named `proxy` in your home folder.

### Step 3.1: Create the directory
SSH into your server and create the folder:
```bash
ssh user@your-server-ip
mkdir -p ~/proxy
exit
```

### Step 3.2: Copy files using SCP
From your local machine, run the following commands to copy the binary and the script. Replace `user@your-server-ip` with your actual server login.

```bash
# Copy the compiled binary
scp proxy-linux user@your-server-ip:~/proxy/

# Copy the persistence script
scp linux-no-root-persistence/run_proxy.sh user@your-server-ip:~/proxy/
```

---

## 4. Setup Persistence

Now, go back to your server to finish the setup.

```bash
ssh user@your-server-ip
cd ~/proxy
```

### Step 4.1: Make files executable
```bash
chmod +x proxy-linux
chmod +x run_proxy.sh
```

### Step 4.2: Setup Cron (Auto-start)
We will use `cron` to make sure the proxy starts when the server reboots.

1. Find out your current path and username. Run `pwd`. It should look like `/home/yourusername/proxy`.
2. Open your crontab:
   ```bash
   crontab -e
   ```
3. Add the following line to the end of the file. **Important**: Replace `<username>` with your actual username found in step 1.

   ```cron
   @reboot /bin/sleep 300 && /usr/bin/nohup /home/<username>/proxy/run_proxy.sh > /dev/null 2>&1 &
   ```
   *(The `sleep 300` waits for 5 minutes after reboot to ensure the network is ready before starting.)*

4. Save and exit the editor.

### Step 4.3: Start it now
You don't need to reboot to start it the first time. Just run:
```bash
nohup ./run_proxy.sh > /dev/null 2>&1 &
```

---

## 5. Verification

To check if the proxy is running, you can look for the process:

```bash
ps aux | grep proxy-linux
```

You can also check the log file created in the directory:
```bash
tail -f ~/proxy/proxy.log
```

You can now configure your browser or other applications to use your SOCKS5 proxy:
- **IP**: Your server's public IP
- **Port**: `40048` (or whatever you set)
- **User**: (Set in Step 1)
- **Pass**: (Set in Step 1)
