# Get system metrics and Expose as REST API

It is golang REST based API to get the below system metrics using [gopsutil](https://github.com/shirou/gopsutil) package

- Hostname
- Total memory
- Free memory
- Memory usage in percentage
- System architecture
- Operating system
- Number of CPU cores
- Cpu usage in percentage

## Demo

![Alt System Metrics](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-system-metrics.gif)

## Usage

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 06-system-monitor dir
cd 06-system-monitor

# build
go build

# run

./monitor-agent

# Access localhost:8080 in browser

![Alt System Metrics](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/images/golang-system-metrics.png)

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
