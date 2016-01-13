# Goller [![Build Status](https://travis-ci.org/antham/goller.svg)](https://travis-ci.org/antham/goller) [![Coverage Status](https://coveralls.io/repos/antham/goller/badge.svg?branch=master&service=github)](https://coveralls.io/github/antham/goller?branch=master) #

Agregate log fields to count occurence

[![asciicast](https://asciinema.org/a/edltgo8u72khflco1ex77g65l.png)](https://asciinema.org/a/edltgo8u72khflco1ex77g65l)

## Install

From [release page](https://github.com/antham/goller/releases) download the binary according to your system architecture

## Usage

### Global

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

### Counter

*Count number of time fields occured*

```bash
usage: goller counter [<flags>] <positions>

Count occurence of field

Flags:
  --help               Show help (also see --help-long and --help-man).
  -d, --delimiter=" | "
                       Separator bewteen results
  -t, --trans=TRANS    Transformers applied to every fields
  -p, --parser=PARSER  Log line parser to use

Args:
  <positions>  Field positions
```

For instance :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller counter 0,1
```

produces :

```bash
1 | hi | everybody
3 | hello | world
```

## Delimiter option (-d/--delimiter)

*Change separator between counted fields*

For instance :

```bash
echo "hello world !" | goller counter -d@ 0,1,2
```

produces :

```bash
1@hello@world@!
```

## Parser option (-p/--parser)

*Change parsing strategy to tokenize log line*

Available functions :
* [reg](#reg)
* [whi](#whi)

### reg

*Parse line according to regexp*

For instance :

```bash
echo "helloworld\!" | goller counter -p 'reg("(h.{4})(w.{4})(.)")' 0,1,2
```

produces :


```bash
1 | hello | world | !

```

### whi

*Parse line following whitespaces*

For instance :

```bash
echo "test1 test2 test3" | goller counter -p whi 0,1,2
```

produces :


```bash
1 | test1 | test2 | test3

```

## Transformer option (-t/--trans)

*Change a field before being counted, transformers could be chained*

Available functions:
* [add](#add)
* [catl](#catl)
* [catr](#catr)
* [dell](#dell)
* [delr](#delr)
* [len](#len)
* [low](#low)
* [match](#match)
* [repl](#repl)
* [sub](#sub)
* [trim](#trim)
* [triml](#triml)
* [trimr](#trimr)
* [upp](#upp)

### add

*Add given integer to field, field must be an integer*

For instance :

```bash
echo "1 2 3" | goller counter -t '0:add("1")' -t '1:add("2")' -t '2:add("3")' 0,1,2
```

produces :

```bash
1 | 2 | 4 | 6
```

### catl

*Concat a string on left side of field*

For instance :

```bash
echo "ello orld" | goller counter -t '0:catl("h")' -t '1:catl("w")' 0,1
```

produces :

```bash
1 | hello | world
```

### catr

*Concat a string on right side of field*

For instance :

```bash
echo "h w" | goller counter -t '0:catr("ello")' -t '1:catr("orld")' 0,1
```

produces :

```bash
1 | hello | world
```

### dell

*Delete n number of characters on left side of field*

For instance :

```bash
echo "123hello 12345world" | goller counter -t '0:dell("3")' -t '1:dell("5")' 0,1
```

produces :

```bash
1 | hello | world
```

### delr

*Delete n number of characters on right side of field*

For instance :

```bash
echo "hello123 world12345" | goller counter -t '0:delr("3")' -t '1:delr("5")' 0,1
```

produces :

```bash
1 | hello | world
```

### len

*Return number of characters in field*

For instance :

```bash
echo "hello world \!" | goller counter -t '0:len' -t '1:len' -t '2:len' 0,1,2
```

produces :

```bash
1 | 5 | 5 | 1
```

### low

*Lowercase field*

For instance :

```bash
echo "HELLO WORLD" | goller counter -t '0:low' -t '1:low' 0,1
```

produces :

```bash
1 | hello | world
```

### match

*Return true if field match regexp, false otherwise*

For instance :

```bash
echo "hello world" | goller counter -t '0:match("hi")' -t '1:match("w.{4}")' 0,1
```

produces :

```bash
1 | false | true
```

### repl

*Replace pattern with string in field*

For instance :

```bash
echo "hello world" | goller counter -t '0:repl("ello","i")' -t '1:repl("world","everybody")' 0,1
```

produces :

```bash
1 | hi | everybody
```

### sub

*Substract given integer to field, field must be an integer*

For instance :

```bash
echo "1 2 3" | goller counter -t '0:sub("1")' -t '1:sub("2")' -t '2:sub("3")' 0,1,2
```

produces :

```bash
1 | 0 | 0 | 0
```

### trim

*Trim all characters given as argument on right and left side of field*

For instance :

```bash
echo "@_@_@hello world\!*\!*" | goller counter -t '0:trim("@_")' -t '1:trim("!*")' 0,1
```
produces :

```bash
1 | hello | world
```

### triml

*Trim all characters given as argument on left side of field*

For instance :

```bash
echo "ooohello dddddworld" | goller counter -t '0:triml("o")' -t '1:triml("d")' 0,1
```
produces :

```bash
1 | hello | world
```

### trimr

*Trim all characters given as argument on right side of field*

For instance :

```bash
echo "hellohhhh worldwwww" | goller counter -t '0:trimr("h")' -t '1:trimr("w")' 0,1
```
produces :

```bash
1 | hello | world
```

### upp

*Uppercase field*

For instance :

```bash
echo "hello world" | goller counter -t '0:upp' -t '1:upp' 0,1
```

produces :

```bash
1 | HELLO | WORLD
```
