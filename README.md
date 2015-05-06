# text-reader
Golang program to print text files for reading at normal human speeds

## Purpose
To print files with adjustable speed and line wrap so text files can be read from the command line.

While this can be used for "speed reading", the object is just to make the reading experience well-formatted and
well-paced. Having shorter line wrap than typical screens makes it easy to eyetrack from one line to the next. And
having a steady pace is engaging and keeps you reading.

## Status
Ready to use

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/text-reader

## Usage
Usage:

    text-reader [options] FILENAME

Options:

    -wpm=WPM    Words per minute [default: 300]
    -wrap=WRAP  Line wrap in number of characters [default: 72]
