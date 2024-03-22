# Teonet fortune single-page application-2 (SPA)

This is simple [Teonet](https://github.com/teonet-go/teonet) one-page application which get fortune message from [Teonet Fortune](https://github.com/teonet-go/teofortune) microservice using Google Cloud Function.

[![GoDoc](https://godoc.org/github.com/teonet-go/teofortune-spa-2?status.svg)](https://godoc.org/github.com/teonet-go/teofortune-spa-2/)
[![Go Report Card](https://goreportcard.com/badge/github.com/teonet-go/teofortune-spa-2)](https://goreportcard.com/report/github.com/teonet-go/teofortune-spa-2)

## Create default frontend

```shell
sudo npm i -g @vue/cli
vue create frontend
cd frontend
rm -rf .git
```

## Build frontend

```shell
cd frontend
npm run build
```

## Add bootstrap package

```shell
npm i bootstrap
```

## Add teoweb package

```shell
# npm config set registry https://npm.pkg.github.com
# npm i @teonet-go/teoweb@0.0.18
npm i teoweb
```

## How to use

To run this application local use next commands:

```shell
go run -tags=dev .
```

By default the teofortune-spa-2 site start at localhost:8080. You can publish
this site to Google [Cloud Run](https://console.cloud.google.com/run).

There is preinstalled teofortune-spa-2 web-site in Google Cloud:
<https://fortune-spa-2.teonet.app>

## Licence

[BSD](LICENSE)
