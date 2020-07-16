# makerlog

:smile: makerlog

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/makerlog)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/makerlog/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/makerlog.svg)](https://github.com/moul/makerlog/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/makerlog.svg)](https://microbadger.com/images/moul/makerlog)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/makerlog/workflows/Go/badge.svg)](https://github.com/moul/makerlog/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/makerlog/workflows/Release/badge.svg)](https://github.com/moul/makerlog/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/makerlog/workflows/PR/badge.svg)](https://github.com/moul/makerlog/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/makerlog.svg)](https://golangci.com/r/github.com/moul/makerlog)
[![codecov](https://codecov.io/gh/moul/makerlog/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/makerlog)
[![Go Report Card](https://goreportcard.com/badge/moul.io/makerlog)](https://goreportcard.com/report/moul.io/makerlog)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/makerlog/badge)](https://www.codefactor.io/repository/github/moul/makerlog)


## Usage

### Login

```console
# get a token by yourself or use the built-in login helper
$ makerlog login --username=YOUR_USERNAME --password=YOUR_PASS
cof2CuG3Aungiegi5udee9zaeBiemu1gohngiusa
```

```console
# then use this token as parameter
$ makerlog --token=YOUR_TOKEN COMMAND
# or env var
$ export MAKERLOG_TOKEN=YOUR_TOKEN
$ makerlog COMMAND
```

### Get raw tasks

```console
# 5865 is my user id (moul), without this option, you get everybody's tasks
$ makerlog raw tasks_list --user=5865 | jq 'del(.results[].user)'
{
  "count": 476,
  "next": "https://api.getmakerlog.com/tasks/?limit=20&offset=20&user=5865",
  "results": [
    {
      "id": 287829,
      "done": true,
      "content": "🎉 release a first version of my Golang's #makerlog API client + CLI",
      "created_at": "2020-07-16T13:33:11.556909+02:00",
      "updated_at": "2020-07-16T13:33:24.366567+02:00",
      "done_at": "2020-07-16T13:33:11.556575+02:00",
      "project_set": [
        {
          "id": 13807,
          "name": "makerlog",
          "user": 5865
        }
      ],
      "attachment": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/Screenshot_-_Manfred_2020-07-16_at_13.32.16.png",
      "og_image": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/8c7c9ae2-3c24-4efb-9975-65a0833615b0.jpg"
    },
    {
      "id": 287805,
      "done": true,
      "content": "🐙 yesterday on GitHub #oss #berty #sgtm",
      "created_at": "2020-07-16T11:09:26.386739+02:00",
      "updated_at": "2020-07-16T11:45:31.442381+02:00",
      "done_at": "2020-07-16T11:09:26.386547+02:00",
      "project_set": [
        {
          "id": 13767,
          "name": "sgtm",
          "user": 5865
        },
        {
          "id": 13394,
          "name": "berty",
          "user": 5865
        },
        {
          "id": 13405,
          "name": "oss",
          "user": 5865
        }
      ],
      "praise": 5,
      "attachment": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/320fc7e105e436af22eb1fa67c9cb415.png",
      "og_image": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/0c062ed4-32eb-487f-8b6b-32e9db41e607.jpg"
    },
    {
      "id": 287472,
      "done": true,
      ...
```

### Get raw notifications

```console
$ makerlog raw notifications_list | jq ". | length"
1

$ makerlog raw notifications_list | jq ".[0]"
{
  "id": 238616,
  "key": "received_praise",
  "verb": "praised you",
  "recipient": {
    "id": 5865,
    "username": "moul",
    "first_name": "Manfred",
    "last_name": "Touron",
    "description": "Coding every day since 2014",
    "avatar": "https://gravatar.com/avatar/da14d5cef42c8142d3d40286f28f29bd?s=150&d=mm&r=pg",
    "streak": 8,
    "timezone": "Europe/Paris",
    "week_tda": 1,
    "twitter_handle": "moul",
    "product_hunt_handle": "m42am",
    "github_handle": "moul",
    "nomadlist_handle": "moul",
    "bmc_handle": "moul",
    "shipstreams_handle": "moul42",
    "website": "https://manfred.life",
    "digest": true,
    "gold": true,
    "accent": "#ff00e8",
    "maker_score": 413,
    "hardcore_mode": true,
    "email_notifications": true,
    "og_image": "https://ik.imagekit.io/makerlog/media/uploads/og/2020/07/16/612312ff-d17f-44e3-9203-264bff6655ec.jpg",
    "date_joined": "2020-06-15T14:18:59.0813+02:00"
  },
  "actor": {
    "id": 1070,
    "username": "lori",
    "first_name": "Lori",
    "last_name": "Karikari",
    "description": "web dev and ops",
    "avatar": "https://ik.imagekit.io/makerlog/media/uploads/avatars/2020/02/29/IMG_20200210_131130_4055221726763142286.jpg",
    "streak": 127,
    "timezone": "Europe/Paris",
    "week_tda": 3,
    "twitter_handle": "LoriKarikari",
    "instagram_handle": "lorikarikari",
    "product_hunt_handle": "lorikarikari",
    "github_handle": "lorikarikari",
    "telegram_handle": "lorikarikari",
    "gold": true,
    "accent": "#9a35ce",
    "maker_score": 659,
    "dark_mode": true,
    "weekends_off": true,
    "og_image": "https://ik.imagekit.io/makerlog/media/uploads/og/2020/07/16/a8fd93d9-ab24-4425-beb0-c22cdab6c8d9.jpg",
    "date_joined": "2018-12-12T13:49:09.353972+01:00"
  },
  "target": {
    "id": 287805,
    "done": true,
    "content": "🐙 yesterday on GitHub #oss #berty #sgtm",
    "created_at": "2020-07-16T11:09:26.386739+02:00",
    "updated_at": "2020-07-16T11:45:31.442381+02:00",
    "done_at": "2020-07-16T11:09:26.386547+02:00",
    "user": {
      "id": 5865,
      "username": "moul",
      "first_name": "Manfred",
      "last_name": "Touron",
      "description": "Coding every day since 2014",
      "avatar": "https://gravatar.com/avatar/da14d5cef42c8142d3d40286f28f29bd?s=150&d=mm&r=pg",
      "streak": 8,
      "timezone": "Europe/Paris",
      "week_tda": 1,
      "twitter_handle": "moul",
      "product_hunt_handle": "m42am",
      "github_handle": "moul",
      "nomadlist_handle": "moul",
      "bmc_handle": "moul",
      "shipstreams_handle": "moul42",
      "website": "https://manfred.life",
      "digest": true,
      "gold": true,
      "accent": "#ff00e8",
      "maker_score": 413,
      "hardcore_mode": true,
      "email_notifications": true,
      "og_image": "https://ik.imagekit.io/makerlog/media/uploads/og/2020/07/16/612312ff-d17f-44e3-9203-264bff6655ec.jpg",
      "date_joined": "2020-06-15T14:18:59.0813+02:00"
    },
    "project_set": [
      {
        "id": 13767,
        "name": "sgtm",
        "user": 5865
      },
      {
        "id": 13394,
        "name": "berty",
        "user": 5865
      },
      {
        "id": 13405,
        "name": "oss",
        "user": 5865
      }
    ],
    "praise": 5,
    "attachment": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/320fc7e105e436af22eb1fa67c9cb415.png",
    "og_image": "https://ik.imagekit.io/makerlog/media/uploads/tasks/2020/07/16/0c062ed4-32eb-487f-8b6b-32e9db41e607.jpg"
  },
  "created": "2020-07-16T11:45:15.248305+02:00",
  "target_type": "task"
}
```

## Install

### Using go

```console
$ go get -u moul.io/makerlog/cmd/makerlog
```

### Releases

See https://github.com/moul/makerlog/releases

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/master/contribute.gif)

I really welcome contributions. Your input is the most precious material. I'm well aware of that and I thank you in advance. Everyone is encouraged to look at what they can do on their own scale; no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/makerlog/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/makerlog/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/makerlog/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/makerlog.svg)](https://starchart.cc/moul/makerlog)

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
