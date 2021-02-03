# SafeCLI &nbsp; [![Build status](https://ci.appveyor.com/api/projects/status/k4aio2mlgpgdsdv8/branch/master?svg=true)](https://ci.appveyor.com/project/zplusfour/safecli/branch/master)&nbsp; [![GitHub issues](https://img.shields.io/github/issues/zplusfour/safecli)](https://github.com/zplusfour/safecli/issues)&nbsp; [![GitHub forks](https://img.shields.io/github/forks/zplusfour/safecli)](https://github.com/zplusfour/safecli/network)&nbsp; [![GitHub stars](https://img.shields.io/github/stars/zplusfour/safecli)](https://github.com/zplusfour/safecli/stargazers)&nbsp; [![Go Reporter](https://goreportcard.com/badge/github.com/zplusfour/safecli)](https://goreportcard.com/report/github.com/zplusfour/safecli)&nbsp; [![codebeat badge](https://codebeat.co/badges/e3cf4a6c-db18-473f-8fdd-ce9e59aeb914)](https://codebeat.co/projects/github-com-zplusfour-safecli-master)

This is a simple CLI written in Go, to help you get the title of a page before you get into the page.

## Usage

Download `safe` from the releases page
### Open a page redirectly from the console
```sh
$ safe open <url>
```

for example:

```sh
$ safe open https://youtube.com # Opens YouTube on your browser
```

### Search, with duckduckgo, bing or google!

```sh
$ safe search <search engine> <search query>
```

for example:

```sh
$ safe search duckduckgo "Bill Gates" # Use quotes because of os.Args things, Opens DuckDuckGo search engine and searchs for 'Bill Gates', on your default browser
```

### Get the title of a site

```sh
$ safe title <sitename>
```

for example:

```sh
$ safe title google.com # Google
```

***

feel free to fork/clone and contribute to the project, have a great day!

- zplusfour