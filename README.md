# `act`

A little script that mananges todos. It is hugely inspired by [Steve Losh's
t](http://stevelosh.com/projects/t/).

## Installation

Just download the `act` file from
[here](https://github.com/math2001/act/releases/latest), alias it in your
`.bashrc` to whatever you want (see the example in the help message).

Just download

    $ act
    =====

    A very simple CLI todo manager by math2001
    Hugely inspired by Steve Losh's t (Python)

    Usage
    -----

    -e int
            Action id you want to edit (default -1)
    -f int
            Action id you have finished (default -1)
    -file string
            A path the file to store tasks (default "./acts")

    Example
    -------

    [~] $ alias act="path/to/act -file=~/acts"
    [~] $ act Fix #12
    [~] $ act Improve help message
    [~] $ act
    1 Fix #12
    2 Improve help message
    [~] $ act -e=2 [CLI] Improve help message
    [~] $ act
    1 Fix #12
    2 [CLI] Improve help message
    [~] $ act -f=2
    [~] $ act
    1 Fix #12

