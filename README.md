# go-ini

Package **ini** provides tools for working with the INI data format, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ini

[![GoDoc](https://godoc.org/github.com/reiver/go-ini?status.svg)](https://godoc.org/github.com/reiver/go-ini)



## Classic INI Format

The classic INI data format looks like this:
```ini
; This is a comment.

[section]

key=value

apple=one
banana=two
cherry=three


[person]
name=Joe Blow
organization=Acme Inc.

```

## De Facto INI Format

Beyond the classic INI format, there is a de facto INI format, that includes some widely used extensions.

The de facto INI format also includes an additional syntax for comments, and an additional syntax for key-value pairs.

For example:
```ini
; This is a comment in the classic INI format.

[section]
key=value

apple=one
banana=two
cherry=three


[person]
name=Joe Blow
organization=Acme Inc.


# This is also a comment in the de facto INI format.

[de facto]
length: 5cm
count: ۵
color: red
```
