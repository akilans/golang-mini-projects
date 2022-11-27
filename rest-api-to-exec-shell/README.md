[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2Frest-api-to-exec-shell&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# REST API to Execute Shell Commands

- It is a golang REST API application to execute shell command from payload
- It is not recommended - Just for learning purpose

## Demo

![Alt Execute Shell Command](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/rest-api-to-exec-shell.gif)

## Usage

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go rest-api-to-exec-shell dir
cd rest-api-to-exec-shell

# run

go run main.go


# sample payload to list files and folder under /home/akilan/Desktop dir
{
    "command": "ls",
    "arguments": ["-l","/home/akilan/Desktop"]
}

# Sample response
{
    "Msg": "total 4\ndrwxrwxr-x 35 akilan akilan 4096 Nov 18 19:55 devops-work\n"
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
