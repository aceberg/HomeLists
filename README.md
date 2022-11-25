<h1><a href="https://github.com/aceberg/HomeLists">
    <img src="https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/logo.png" width="20" />
</a>HomeLists</h1>
<br/>

[![Docker](https://github.com/aceberg/HomeLists/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/HomeLists/actions/workflows/main-docker-all.yml)
[![Binary-release](https://github.com/aceberg/HomeLists/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/HomeLists/actions/workflows/release.yml)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/homelists)

Count consumable supplies    
https://github.com/aceberg/HomeLists


![Screenshot1](https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/Screenshot%202022-11-07%20at%2022-26-22%20Home%20Lists%20-%20Dashboard.png)

![Screenshot2](https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/Screenshot%202022-11-07%20at%2022-24-17%20Home%20Lists%20-%2001-Fridge.png)

![Screenshot3](https://raw.githubusercontent.com/aceberg/HomeLists/main/assets/Screenshot%202022-11-07%20at%2022-27-15%20Home%20Lists%20-%2003-Supplies.png)

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

## Support

If you like this app, please support me
- BTC: bc1qj59rxmfvanvqqltq9t73qls4su3xrvwuv3sxhr
- ETH: 0x276124c218aa8110F96989AA1f6f2Bb960C234B7
- USDT(ETH Network): 0x276124c218aa8110F96989AA1f6f2Bb960C234B7


## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/go-homelists/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [List icons created by Freepik - Flaticon](https://www.flaticon.com/free-icon/list_3039383?related_id=3039383&origin=search)