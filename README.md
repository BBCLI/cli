# BBCLI

## About

A CLI for Bitbucket Cloud written in Go

## Installation

```sh
git clone https://github.com/BBCLI/cli.git
cd cli
go build -o bbc && sudo mv bbc /usr/local/bin/  
```

## Init

To configure your access, use the `init` command:

```sh
cli init
```

username: <your bitbucket username> (https://bitbucket.org/account/settings/)

password: <bitbucket app password> (https://bitbucket.org/account/settings/app-passwords/)

## Commands

### Pull Requests

- `bbc pr list` - List all pull requests
- `bbc pr approve {pr id}` - Approve a pull request
- `bbc pr create {origin} {destination}` - Create a pull request.
- `bbc pr create {destination}` - Create a pull request with the current branch as the origin.
- `bbc pr diff {pr id}` - Show the diff of a pull request
- `bbc pr merge {pr id}` - Merge a pull request
- `bbc pr open {pr id}` - Open a pull request in the browser
- `bbc pr comment create {pr id}` - Create a comment on a pull request without file context
- `bbc pr comment create {pr id} {file path} {line_number}` - Create a comment on a pull request with file context
- `bbc pr comment reply {pr id} {comment_id}` - Reply to a comment on a pull request
- `bbc pr comment resolve {pr id} {comment_id}` - Resolve a comment on a pull request

### MISC
- `bbc init` - Configure your access
- `bbc me` - Show your user information