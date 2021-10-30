[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F03-web-monitor&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Website Monitoring and Alerting using Golang

This application monitor websites with expected status code. If the response status code does not match with expected status code, connection refused from server, then it triggers an email and alerts user

### Prerequisites

- Go
- Configure a SMTP server - [SMTP settings for Gmail](hhttps://myaccount.google.com/u/4/security) - Enable Less secure app access
- set GMAIL_ID, GMAIL_PASSWORD as environment variables

### Golang packages

- net/http
- time
- net/smtp
- os

### Flow chart

![Web monitoring flow chart](https://github.com/akilans/golang-mini-projects/blob/main/images/web-monitor.png?raw=true)

### Demo

![Alt Web monitoring](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/golang-web-monitor.gif)
