# git-dailylog

commit log to dailylog

## Description

![ScreenShot](./sample/ss.png)

## Usage

### Initialize
> create `.dailylog` format file or copy from ~/.dailylog 
```console
$ git dailylog init
```
#### Log Format File .dailylog

Reference: [Git log format string cheatsheet](https://devhints.io/git-log-format)

```--pretty="format:[format]"```
```.dailylog
"%h: %ad %an: %s"
```

### Get
get commit log from:time.start to:today.end

```console
$ git dailylog get [time: default today]
```

#### TimeFormat

> for example Today = 2018-06-28

| format | from | to |
| :---: | :---: | :---: |
| today | 2018-06-28 00:00:00 | 2018-06-28 23:59:59 |
| yesterday | 2018-06-27 00:00:00 | 2018-06-28 23:59:59 |
| 2days | 2018-06-26 00:00:00 | 2018-06-28 23:59:59 |
| 1weeks | 2018-06-21 00:00:00 | 2018-06-28 23:59:59 |
| 1month | 2018-05-28 00:00:00 | 2018-06-28 23:59:59 |
| 1years | 2017-06-28 00:00:00 | 2018-06-28 23:59:59 |

#### Options
##### Reverse
> default git log Desc. Asc git log use when reverse option.  
```console
$ git dailylog get today --reverse
```

##### Author Filter
```console
$ git dailylog get today --author=aozora0000
```


## Install
To install, use `go get`:

```console
$ go get -d github.com/aozora0000/git-dailylog
```

## Contribution

1. Fork ([https://github.com/aozora0000/git-dailylog/fork](https://github.com/aozora0000/git-dailylog/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[aozora0000](https://github.com/aozora0000)
