# `randuser`

`randuser` is a command-line application that generates a random
username using the same algorithm as [diceware
passwords](https://en.wikipedia.org/wiki/Diceware). By default, the
generated name follows the `{{ adjective }}{{ noun }}` format.

The main use case for the tool is to generate usernames for online
services such as social media platforms. Words were sourced from the [23
thousand word
list](https://git.sr.ht/~jamesponddotco/acopw-go/blob/trunk/words/word-list.txt)
I use for [acopw-go](https://sr.ht/~jamesponddotco/acopw-go/) and
classified using [WordNet](https://wordnet.princeton.edu/).

See _randuser(1)_ after installing for more information.

## Installation

### From source

First install the dependencies:

- Go 1.22 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ randuser --help
NAME:
   randuser - generate a random username

USAGE:
   randuser [global options]

VERSION:
   1.0.0

GLOBAL OPTIONS:
   --title-case, -t          specifies whether the generated name should be title cased
   --format value, -f value  specifies the format of the generated name
   --help, -h                show help
   --version, -v             print the version

$ randuser -t -f '{{ noun }}{{ adjective }}'
ShapedPessimist
```

See _randuser(1)_ for more usage information.

## Contributing

Anyone can help make `randuser` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/randuser-devel) and report
bugs on the [issue
tracker](https://todo.sr.ht/~jamesponddotco/randuser).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/randuser-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/randuser-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/randuser).

---

Released under the [GPL-2.0 license](LICENSE.md).
