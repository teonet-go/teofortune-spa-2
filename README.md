# Teonet fortune single-page application-2 (SPA)

This is simple [Teonet](https://github.com/teonet-go/teonet) one-page application which get fortune message from [Teonet Fortune](https://github.com/teonet-go/teofortune) microservice using Google Cloud Function.

[![GoDoc](https://godoc.org/github.com/teonet-go/teofortune-spa-2?status.svg)](https://godoc.org/github.com/teonet-go/teofortune-spa-2/)
[![Go Report Card](https://goreportcard.com/badge/github.com/teonet-go/teofortune-spa-2)](https://goreportcard.com/report/github.com/teonet-go/teofortune-spa-2)

## Create default frontend

    sudo npm i -g @vue/cli
    vue create frontend
    cd frontend
    rm -rf .git

## Build frontend

    cd frontend
    npm run build

## How to use

To run this application local use next commands:

    go run .

or (run without server)

    cd frontend
    npm run serve

By default the teofortune-spa-2 site start at localhost:8080. You can publish
this site to Google [Cloud Run](https://console.cloud.google.com/run).

There is preinstalled teofortune-spa-2 web-site in Google Cloud:
<https://fortune-spa-2.teonet.app>

## Licence

[BSD](LICENSE)
