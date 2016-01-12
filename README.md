# Goller [![Build Status](https://travis-ci.org/antham/goller.svg)](https://travis-ci.org/antham/goller) [![Coverage Status](https://coveralls.io/repos/antham/goller/badge.svg?branch=master&service=github)](https://coveralls.io/github/antham/goller?branch=master) #

Agregate log fields to count occurence

[![asciicast](https://asciinema.org/a/edltgo8u72khflco1ex77g65l.png)](https://asciinema.org/a/edltgo8u72khflco1ex77g65l)

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

## Parser option

Change parsing strategy with *-p* option :
* [reg](#reg)
* [whi](#whi)

### reg

*Parse line according to regexp*

For instance :

```bash
echo "helloworld\!"|goller counter -p 'reg("(h.{4})(w.{4})(.)")' 0,1,2
```

produces :


```bash
1 | hello | world | !

```

### whi

*Parse line following whitespaces*

For instance :

```bash
echo "test1 test2 test3"|goller counter -p whi 0,1,2
```

produces :


```bash
1 | test1 | test2 | test3

```
