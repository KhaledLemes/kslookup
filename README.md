# kslookup

A simple CLI tool built in Go to look up information about websites, such as their public IP addresses and the servers they are running on.  
It uses the [`urfave/cli`](https://github.com/urfave/cli) package to provide a clean command-line interface.

---

## Features
- Lookup the **public IPs** of a website.  
- Find out the **name servers** a website is running on.  
- Easy to use, fast, and lightweight.

---

## Installation
Clone the repository and build the project:

```bash
git clone https://github.com/KhaledLemes/kslookup.git
cd kslookup
go build -o kslookup
