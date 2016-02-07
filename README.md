# Goller [![Build Status](https://travis-ci.org/antham/goller.svg)](https://travis-ci.org/antham/goller) [![codecov.io](https://codecov.io/github/antham/goller/coverage.svg?branch=master)](https://codecov.io/github/antham/goller?branch=master)

Agregate log fields, count field occurence

[![asciicast](https://asciinema.org/a/d5o27yx3kclnluwaan2p46fki.png)](https://asciinema.org/a/d5o27yx3kclnluwaan2p46fki)

## Install

From [release page](https://github.com/antham/goller/releases) download the binary according to your system architecture

## Usage

### Global

```bash
usage: goller [<flags>] <command> [<args> ...]

Aggregate log fields and count occurences

Flags:
  --help     Show help (also see --help-long and --help-man).
  --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  group [<flags>] <parser> <positions>
    Group occurence of field

  tokenize <parser>
    Show how first log line is tokenized
```

### Tokenize

*Show how first log line is parsed using given parsing strategy and display tokens with their positions*

```bash
usage: goller tokenize <parser>

Show how first log line is tokenized

Flags:
  --help     Show help (also see --help-long and --help-man).
  --version  Show application version.

Args:
  <parser>  Log line parser to use
```

### Group

*Group and count field occurences*

```bash
usage: goller group [<flags>] <parser> <positions>

Group occurence of field

Flags:
  --help           Show help (also see --help-long and --help-man).
  --version        Show application version.
  -d, --delimiter=" | "
                   Separator between results
  -i, --ignore     Ignore lines wrongly parsed
  -t, --transformer=TRANSFORMER
                   Transformers applied to every fields
  -s, --sort=SORT  Sort lines

Args:
  <parser>     Log line parser to use
  <positions>  Field positions
```

*A log line is splitted according to given parsing strategy, you can then refer every field using its position number. 0 position is a special position, it counts number of time a field occured*

If we want to parse thoses lines :

```bash
hello world
hello world
hi everybody
hello world
```

we will do :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller group whi 0,1,2
```

it produces :

```bash
3 | hello | world
1 | hi | everybody
```

we can reorganize fields as we want :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller group whi 2,1,0
```

it produces :

```bash
world | hello | 3
everybody | hi | 1
```

we can  keep only fields that matter :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller group whi 1
```

it produces :

```bash
hello
hi
```

## Parser argument

*Parsing strategy used to tokenize log line*

Available functions :
* [clf](#clf)
* [reg](#reg)
* [spl](#spl)
* [whi](#whi)

### clf

*Parse line following Common Log Format (NCSA Common log format)*

For instance :

```bash
echo '127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326'| goller group clf 1,2,3,4,5,6,7
```

produces :

```bash
127.0.0.1 | user-identifier | frank | 10/Oct/2000:13:55:36 -0700 | GET /apache_pb.gif HTTP/1.0 | 200 | 2326
```

### reg

*Parse line according to regexp*

For instance :

```bash
echo "helloworld\!" | goller group 'reg("(h.{4})(w.{4})(.)")' 1,2,3
```

produces :


```bash
hello | world | !

```

### whi

*Parse line following whitespaces*

For instance :

```bash
echo "test1 test2 test3" | goller group whi 1,2,3
```

produces :


```bash
test1 | test2 | test3
```

### spl

*Split lines using given string*

For instance :

```bash
echo "test1_test2_test3" | goller group 'spl("_")' 1,2,3
```

produces :

```bash
test1 | test2 | test3

```
## Ignore option (-i/--ignore)

*Ignore lines wrongly tokenized by parser*

For instance :

```bash
echo "hello world\nHi there\nHi everybody\nHi" | goller group whi 1,2
```

produces :

```bash
Wrong parsing strategy (based on first line tokenization), got 1 tokens instead of 2
Line : Hi
```

If we set the flag :

```bash
echo "hello world\nHi there\nHi everybody\nHi" | goller group -i whi 1,2
```

it produces :

```bash
hello | world
Hi | there
Hi | everybody
```

## Delimiter option (-d/--delimiter)

*Change separator between counted fields*

For instance :

```bash
echo "hello world \!" | goller group whi -d@ 1,2,3
```

produces :

```bash
hello@world@!
```

## Transformer option (-t/--transformer)

*Change a field before being counted, transformers could be chained*

For instance :

```bash
echo "1 2 3" | goller group whi -t '1:add("1").sub("1").add("10")' -t '2:add("2").sub("1").add("10")' -t '3:add("3").sub("1").add("10")' 1,2,3
```

produces :

```bash
11 | 13 | 15
```

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
echo "1 2 3" | goller group whi -t '1:add("1")' -t '2:add("2")' -t '3:add("3")' 1,2,3
```

produces :

```bash
2 | 4 | 6
```

### catl

*Concat a string on left side of field*

For instance :

```bash
echo "ello orld" | goller group whi -t '1:catl("h")' -t '2:catl("w")' 1,2
```

produces :

```bash
hello | world
```

### catr

*Concat a string on right side of field*

For instance :

```bash
echo "h w" | goller group whi -t '1:catr("ello")' -t '2:catr("orld")' 1,2
```

produces :

```bash
hello | world
```

### dell

*Delete n number of characters on left side of field*

For instance :

```bash
echo "123hello 12345world" | goller group whi -t '1:dell("3")' -t '2:dell("5")' 1,2
```

produces :

```bash
hello | world
```

### delr

*Delete n number of characters on right side of field*

For instance :

```bash
echo "hello123 world12345" | goller group whi -t '1:delr("3")' -t '2:delr("5")' 1,2
```

produces :

```bash
hello | world

```

### len

*Return number of characters in field*

For instance :

```bash
echo "hello world \!" | goller group whi -t '1:len' -t '2:len' -t '3:len' 1,2,3
```

produces :

```bash
5 | 5 | 2
```

### low

*Lowercase field*

For instance :

```bash
echo "HELLO WORLD" | goller group whi -t '1:low' -t '2:low' 1,2
```

produces :

```bash
hello | world
```

### match

*Return true if field match regexp, false otherwise*

For instance :

```bash
echo "hello world" | goller group whi -t '1:match("hi")' -t '2:match("w.{4}")' 1,2
```

produces :

```bash
false | true
```

### repl

*Replace pattern with string in field*

For instance :

```bash
echo "hello world" | goller group whi -t '1:repl("ello","i")' -t '2:repl("world","everybody")' 1,2
```

produces :

```bash
hi | everybody
```

### sub

*Substract given integer to field, field must be an integer*

For instance :

```bash
echo "1 2 3" | goller group whi -t '1:sub("1")' -t '2:sub("2")' -t '3:sub("3")' 1,2,3
```

produces :

```bash
0 | 0 | 0
```

### trim

*Trim all characters given as argument on right and left side of field*

For instance :

```bash
echo "@_@_@hello world\!*\!*" | goller group whi -t '1:trim("@_")' -t '2:trim("!*")' 1,2
```
produces :

```bash
hello | world
```

### triml

*Trim all characters given as argument on left side of field*

For instance :

```bash
echo "ooohello dddddworld" | goller group whi -t '1:triml("o")' -t '2:triml("d")' 1,2
```
produces :

```bash
hello | world
```

### trimr

*Trim all characters given as argument on right side of field*

For instance :

```bash
echo "hellohhhh worldwwww" | goller group whi -t '1:trimr("h")' -t '2:trimr("w")' 1,2
```
produces :

```bash
hello | world
```

### upp

*Uppercase field*

For instance :

```bash
echo "hello world" | goller group whi -t '1:upp' -t '2:upp' 1,2
```

produces :

```bash
HELLO | WORLD
```

## Sort option (-s/--sort)

*Sort a field according to given function, sorters could be used with several fields*

For instance :

```bash
echo "3 8 2\n4 9 3\n3 8 0\n3 1 10\n3 9 1\n1 9 1\n2 9 1"| go run main.go group -s "1:int,2:int,3:int" whi 1,2,3
```
produces

```bash
1 | 9 | 1
2 | 9 | 1
3 | 1 | 10
3 | 8 | 0
3 | 8 | 2
3 | 9 | 1
4 | 9 | 3
```

Available functions:
* [int](#int)
* [strl](#strl)
* [str](#str)

### int

*Sort integer fields*

For instance :

```bash
echo "5\n7\n9\n10\n6\n1\n5" | goller group whi -s "1:int" 1
```
produces :

```bash
1
5
6
7
9
10
```

### strl

*Sort using size string*

For instance :

```bash
echo "aaaaa\naaaa\naa\na\naaa" | goller group whi -s "1:strl" 1
```

produces :

```bash
a
aa
aaa
aaaa
aaaaa
```

### str

*Sort using lexicographic order*

For instance :

```bash
echo "e\nd\nb\nf\na\ng\nc" | goller group whi -s "1:str" 1
```

produces :

```bash
a
b
c
d
e
f
g
```
