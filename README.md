# Goller [![Build Status](https://travis-ci.org/antham/goller.svg)](https://travis-ci.org/antham/goller) [![Coverage Status](https://coveralls.io/repos/antham/goller/badge.svg?branch=master&service=github)](https://coveralls.io/github/antham/goller?branch=master) #

Agregate log fields to count occurence

![](https://raw.githubusercontent.com/antham/goller/gh-pages/example.gif)

## Usage

```bash
usage: goller [<flags>] <command> [<args> ...]

Logger parser

Flags:
  --help  Show help (also see --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  counter [<flags>] <positions>
    Count occurence of field
```

## Example

From this kind of log entry :

```bash
172.17.0.1 - - [19/Dec/2015:11:18:19 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:29 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:30 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:31 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:32 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:32 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:33 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:35 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:18:36 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:20:24 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:20:25 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:20:26 +0000] "GET / HTTP/1.1" 502 537 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:22:06 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:22:24 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:24:36 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:24:50 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:24:51 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:25:51 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:25:52 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:26:48 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:11:28:04 +0000] "GET / HTTP/1.1" 502 1569 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:26:11 +0000] "GET / HTTP/1.1" 502 1569 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:27:04 +0000] "GET / HTTP/1.1" 502 172 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:27:55 +0000] "GET / HTTP/1.1" 502 219 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:28:05 +0000] "GET / HTTP/1.1" 502 219 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:29:29 +0000] "GET / HTTP/1.1" 502 1599 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:29:48 +0000] "GET / HTTP/1.1" 502 5427 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:14:30:24 +0000] "GET / HTTP/1.1" 502 19171 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:15:51:27 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:15:52:50 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:15:53:21 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [19/Dec/2015:15:53:22 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [27/Dec/2015:01:32:07 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
172.17.0.1 - - [28/Dec/2015:11:47:36 +0000] "GET / HTTP/1.1" 502 158828 "-" "curl/7.45.0" "-"
172.17.0.1 - - [28/Dec/2015:12:21:09 +0000] "GET / HTTP/1.1" 499 0 "-" "curl/7.45.0" "-"
```

We want to count number of http error message related to an ip, based on field positions we write :

```bash
goller counter 0,13
```
it produces :

```bash
28 | 172.17.0.1 | 502
7 | 172.17.0.1 | 499
```

We keep the same fields but we want to count as well occurence of size response, we write :

```bash
goller counter 0,13,14
```

it produces :

```bash
12 | 172.17.0.1 | 502 | 537
8 | 172.17.0.1 | 502 | 172
2 | 172.17.0.1 | 502 | 1569
2 | 172.17.0.1 | 502 | 219
1 | 172.17.0.1 | 502 | 1599
1 | 172.17.0.1 | 502 | 5427
1 | 172.17.0.1 | 502 | 19171
7 | 172.17.0.1 | 499 | 0
1 | 172.17.0.1 | 502 | 158828
```
