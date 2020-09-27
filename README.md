# leb128utils
A command line tool that helps decode and encode LEB128 integers.

## Usage

```
## Encode a ULEB128 integer
$ leb128utils -i=0x1234
$ leb128utils -i=1234
$ leb128utils -i=0b1010101

## Encode a SLEB128 integer
$ leb128utils -s -i=-0x1234
$ leb128utils -s -i=-1234
$ leb128utils -s -i=-0b1010101
```
