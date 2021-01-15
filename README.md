# The Letter Generator in Go

This tool helps you to generate form letters, based on a template for
multiple recipients.


## Setup

* Install the letter generator

  ~~~bash
  go get -u github.com/feduxorg/letter-generator
  ~~~

* Install LaTeX

  Make sure, you installed LaTeX on your system. This tools uses
  `pdflatex` in the background to generate the PDF from your letter template.

## Usage

### Initialize directory for your letter

* (optional) Install examples

  ~~~bash
  # Clone repository
  git clone https://github.com/feduxorg/letter-generator /tmp/letter-generator

  # Either create your own templates or use our examples

  ## Move examples to the correct place
  mv /tmp/letter-generator/docs/examples/en/ ~/.local/share/letter-template/
  ## or
  mv /tmp/letter-generator/docs/examples/de/ ~/.local/share/letter-template/

  # setup git repository
  cd ~/.local/share/letter-template/
  git init
  git add .
  git commit -m "Init"
  ~~~

* Create directory for your letter

  ~~~bash
  mkdir my-letter
  cd my-letter
  ~~~

* Setup directory

  * &ndash; with local repository `~/.local/share/letter-template/.git`

    ~~~bash
    # current directory
    lg init -V

    # given directory
    lg init -V my/dir
    ~~~

  * &ndash; with remote repository `https://github.com/xxxxx/xxxx.git`

    ~~~bash
    lg init -V --template-source `https://github.com/xxxxx/xxxx.git`
    ~~~

### Build letters

This will build the letters. Make sure you've initialized the directory before.

~~~bash
lg build -V
~~~

After building letters, you will find a directory called `letters` in the
current directory.

## Copyright

(c) 2021, Dennis Günnewig
