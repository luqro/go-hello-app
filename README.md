# go-hello-app

A simple "Hello, world!" web service written in Go.

This application is designed to run on AWS EC2 instances and is managed using Terraform and Ansible.

## 🚀 Features

- HTTP server on port 80
- Responds with "Hello, world from GitHub!"
- Suitable for deployment with systemd
- Minimal dependencies

## 📁 Project Structure

```
go-hello-app/
├── hello.go       # Main application file
└── go.mod         # Go module definition
```

## 🛠 Requirements

- Go 1.20+
- Linux (Ubuntu recommended)

## ▶️ Run Locally

```bash
go run hello.go
```

Then open:

```
http://localhost:80
```

You should see:

```
Hello, world from GitHub!
```

## ⚙️ Build

```bash
go build -o app
```

## 📦 Deployment

This repository is intended to be cloned remotely by Ansible and launched via `systemd`. The systemd unit file should look like this:

```ini
[Unit]
Description=Simple Go Web App
After=network.target

[Service]
ExecStart=/opt/go-app/app
WorkingDirectory=/opt/go-app
Restart=always
User=ubuntu

[Install]
WantedBy=multi-user.target
```

## 📄 License

MIT