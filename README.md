# HomeLists

[![Docker](https://github.com/aceberg/HomeLists/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/HomeLists/actions/workflows/main-docker-all.yml)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/homelists)

Count consumable supplies    
https://github.com/aceberg/HomeLists


![Screenshot1](https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/Screenshot%202022-10-06%20at%2015-13-20%20Home%20Lists%20-%20Dashboard.png)

![Screenshot2](https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/Screenshot%202022-10-06%20at%2015-16-48%20Home%20Lists%20-%2002-Test.png)

## Quick start

```sh
docker run --name homelists \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/homelists:/data/homelists \
-p 8842:8842 \
aceberg/homelists
```

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DBPATH    | Path to Database | /data/homelists/sqlite.db |
| GUIPORT   | Port for web GUI | 8842 |
| THEME | Any theme name from https://bootswatch.com in lowcase | superhero |
| TZ | Set your timezone for correct time | "" |

## Config file

Config file path is `/data/homelists/config`. All variables could be set there.

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/go-homelists/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)