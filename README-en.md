# utterances-oauth.go

Unofficial [utterances](https://github.com/utterance) API in Go

[![Build Status](https://cloud.drone.io/api/badges/GizmoOAO/utterances-oauth.go/status.svg)](https://cloud.drone.io/GizmoOAO/utterances-oauth.go)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/GizmoOAO/utterances-oauth.go/main)
[![GitHub](https://img.shields.io/github/license/GizmoOAO/utterances-oauth.go)](./LICENSE)

[中文](./README.md) | English

# Getting Started

## Config

Create a file named `.env` at the root. File should have the following values: [example](./.env.example)

- BOT_TOKEN: a personal access token that will be used when creating GitHub issues. Generate [here](https://github.com/settings/tokens/new?scopes=public_repo).
- CLIENT_ID: The client id to be used in the [GitHub OAuth web application flow](https://developer.github.com/v3/oauth/#web-application-flow)
- CLIENT_SECRET: The client secret for the OAuth web application flow
- STATE_PASSWORD: 32 character password for encrypting state in request headers/cookies. Generate [here](https://lastpass.com/generatepassword.php).
- ORIGINS: comma delimited list of permitted origins. For CORS.

## Deploy on your own Vercel instance

Click on the deploy button to get started!

[![Deploy to Vercel](https://vercel.com/button)](https://vercel.com/import/project?template=https://github.com/GizmoOAO/utterances-oauth.go)

## Basic

Building is quite easy, just make sure you have [Go](https://golang.org/) installed, and run `go build` You should be able to run the compiled executable after making required changes to `.env`

```bash
git clone https://github.com/GizmoOAO/utterances-oauth.go.git
cd utterances-oauth.go
go build
```

## Docker

By far the easiest way to get up and running. Refer to the example `docker-compose.yaml` example file, put it on your Docker host and run:

```bash
docker-compose up -d
```

# Thanks

- [utterance](https://github.com/utterance)

# License

MIT
