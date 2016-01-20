# Goller [![Build Status](https://travis-ci.org/antham/goller.svg)](https://travis-ci.org/antham/goller) [![Coverage Status](https://coveralls.io/repos/antham/goller/badge.svg?branch=master&service=github)](https://coveralls.io/github/antham/goller?branch=master) #

Agregate log fields to count occurence

[![asciicast](https://asciinema.org/a/edltgo8u72khflco1ex77g65l.png)](https://asciinema.org/a/edltgo8u72khflco1ex77g65l)

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

  group [<flags>] <positions>
    Group occurence of field


exit status 1
```

### Group

*Group and count field occurences*

```bash
usage: goller group [<flags>] <positions>

Group occurence of field

Flags:
  --help            Show help (also see --help-long and --help-man).
  --version         Show application version.
  -d, --delimiter=" | "
                    Separator between results
  -t, --transformer=TRANSFORMER
                    Transformers applied to every fields
  -p, --parser=whi  Log line parser to use
  -s, --sort=SORT   Sort lines

Args:
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
echo "hello world\nhello world\nhi everybody\nhello world" | goller group 0,1,2
```

it produces :

```bash
3 | hello | world
1 | hi | everybody
```

we can reorganize fields as we want :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller group 2,1,0
```

it produces :

```bash
world | hello | 3
everybody | hi | 1
```

we can  keep only fields that matter :

```bash
echo "hello world\nhello world\nhi everybody\nhello world" | goller group 1
```

it produces :

```bash
hello
hi
```

## Delimiter option (-d/--delimiter)

*Change separator between counted fields*

For instance :

```bash
echo "hello world \!" | goller group -d@ 1,2,3
```

produces :

```bash
hello@world@!
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
echo "helloworld\!" | goller group -p 'reg("(h.{4})(w.{4})(.)")' 1,2,3
```

produces :


```bash
hello | world | !

```

### whi

*Parse line following whitespaces*

For instance :

```bash
echo "test1 test2 test3" | goller group -p whi 1,2,3
```

produces :


```bash
test1 | test2 | test3

```

## Transformer option (-t/--transformer)

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
echo "1 2 3" | goller group -t '1:add("1")' -t '2:add("2")' -t '3:add("3")' 1,2,3
```

produces :

```bash
2 | 4 | 6
```

### catl

*Concat a string on left side of field*

For instance :

```bash
echo "ello orld" | goller group -t '1:catl("h")' -t '2:catl("w")' 1,2
```

produces :

```bash
hello | world
```

### catr

*Concat a string on right side of field*

For instance :

```bash
echo "h w" | goller group -t '1:catr("ello")' -t '2:catr("orld")' 1,2
```

produces :

```bash
hello | world
```

### dell

*Delete n number of characters on left side of field*

For instance :

```bash
echo "123hello 12345world" | goller group -t '1:dell("3")' -t '2:dell("5")' 1,2
```

produces :

```bash
hello | world
```

### delr

*Delete n number of characters on right side of field*

For instance :

```bash
echo "hello123 world12345" | goller group -t '1:delr("3")' -t '2:delr("5")' 1,2
```

produces :

```bash
hello | world

```

### len

*Return number of characters in field*

For instance :

```bash
echo "hello world \!" | goller group -t '1:len' -t '2:len' -t '3:len' 1,2,3
```

produces :

```bash
5 | 5 | 2
```

### low

*Lowercase field*

For instance :

```bash
echo "HELLO WORLD" | goller group -t '1:low' -t '2:low' 1,2
```

produces :

```bash
hello | world
```

### match

*Return true if field match regexp, false otherwise*

For instance :

```bash
echo "hello world" | goller group -t '1:match("hi")' -t '2:match("w.{4}")' 1,2
```

produces :

```bash
false | true
```

### repl

*Replace pattern with string in field*

For instance :

```bash
echo "hello world" | goller group -t '1:repl("ello","i")' -t '2:repl("world","everybody")' 1,2
```

produces :

```bash
hi | everybody
```

### sub

*Substract given integer to field, field must be an integer*

For instance :

```bash
echo "1 2 3" | goller group -t '1:sub("1")' -t '2:sub("2")' -t '3:sub("3")' 1,2,3
```

produces :

```bash
0 | 0 | 0
```

### trim

*Trim all characters given as argument on right and left side of field*

For instance :

```bash
echo "@_@_@hello world\!*\!*" | goller group -t '1:trim("@_")' -t '2:trim("!*")' 1,2
```
produces :

```bash
hello | world
```

### triml

*Trim all characters given as argument on left side of field*

For instance :

```bash
echo "ooohello dddddworld" | goller group -t '1:triml("o")' -t '2:triml("d")' 1,2
```
produces :

```bash
hello | world
```

### trimr

*Trim all characters given as argument on right side of field*

For instance :

```bash
echo "hellohhhh worldwwww" | goller group -t '1:trimr("h")' -t '2:trimr("w")' 1,2
```
produces :

```bash
hello | world
```

### upp

*Uppercase field*

For instance :

```bash
echo "hello world" | goller group -t '1:upp' -t '2:upp' 1,2
```

produces :

```bash
HELLO | WORLD
```

## Sort option (-s/--sort)

*Sort a field according to given function*

Available functions:
* [int](#int)
* [strl](#strl)
* [str](#str)

### int

*Sort integer fields*

For instance :

```bash
echo "5\n7\n9\n10\n6\n1\n5" | goller group -s "1:int" 1
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
echo "eeeee\ndddd\nbb\na\nccc" | goller group -s "1:strl" 1
```

produces :

```bash
a
bb
ccc
dddd
eeeee
```

### str

*Sort using lexicographic order*

For instance :

```bash
echo "e\nd\nb\nf\na\ng\nc" | goller group -s "1:str" 1
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
