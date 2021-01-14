# Letter Generator

This small tool helps you to generate form letters.


## Usage

### Install

~~~bash
go get github.com/feduxorg/letter-generator-go
~~~

### Initialize

This will look for `~/.local/share/letter-template/.git` and for a defined
remote repository.

* Setup current directory:

  ~~~
  lg init -V
  ~~~

* Setup given directory:

  ~~~
  lg init -V my/dir
  ~~~

### Build

This will build the letters. Make sure you've initialized the directory before.

~~~
lg build -V
~~~

After building letters, you will find a directory called `letters` in the
current directory.

## Copyright

(c) 2017, Dennis Günnewig

