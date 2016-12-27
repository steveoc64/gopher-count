# gopher-count
Analyse the output from a GopherJS compile, and see whats taking up the most space.

## What and Why ?

[GO Language](http://golang.org) is a modern, compiled language for writing systems software. Amongst other things, it is particularly well suited to writing backend code for web based systems.

[GopherJS](http://gopherjs.org) is another tool which uses the Go compiler to 
generate workable Javascript for web apps from Go source code. There are situations where it makes good sense to write the front end in Go, especially for apps where the backend is also written in Go. 

One of the downsides of GopherJS, in its current state (end of 2016), is that the generated payload of the compiled Javascript can be quite large. Everytime your front end app imports another Go package for general use, the entire front end payload can grow ... very quickly ... by a lot.

This tool (gopher-count) is a simple tool that is used post-compilation on the front end, to show exactly how the size of the compiled JS file is distributed across the various packages.




## Installation

From go :
`go get github.com/steveoc64/gopher-count`

From git:
```
$ git clone git@github.com:steveoc64/gopher-count.git
$ cd gopher-count
$ make install
```

## Usage


```
$ cd <my-big-gopherjs-project>
$ gopherjs build *.go -o <my-project-name>.js -m
$ gopher-count <my-project-name>.js
```

This generates output to stdout, in the form :

<size-in-bytes> TAB <name-of-GO-package>

With 1 line per package.

Packages are output in the order in which they are presented in the compiled Javascript, and in standard UNIX fashion, the output can be piped through whatever filters you want.

Example, with output sorted in order of size of package :

```
$ gopherjs build *.go -o output-file.js -m
$ gopher-count output-file.js | sort -n
```



