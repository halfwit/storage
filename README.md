# Storage

## Overview

Storage is a Plumb destination that will fetch a remote resource
With it, you can fetch resources from the web via the Plan9 Plumber, and is meant to be used in conjunction with [store](https://github.com/halfwit/store); though it can be used with `plumb` proper, with the caveat that you set dst, described below.

## Usage

Storage relies on the filename attribute being set via plumb rules: you MUST set `attr filename=/path/I/want/to/download/file/to`.
Additionally, setting the destination (`dst`) field in your plumb message to `storage` allows you to issue normal plumbs, should you not want to simply view the resource your normal way:

```

type	is	text
data	matches	'$protocol/$urlchars'
data	matches	'https://gist.github.com/($urlchars)'
data	set	'https//gist.githubusercontent.com/$1/raw'
attr	add	'filename=/usr/halfwit/notes/gist/$1'
plumb	to	storage
plumb	client	storage

# This rule matches 
type	is	text
data	matches	'$protocol/$urlchars'
data	matches	'https://gist.github.com/($urlchars)'
plumb	to	web
plumb	client	window	$browser

```

## Using `store`

Store parses remote resources for an appropriate mime-type, and issues plumb messages with `type	is	image/png`, where image/png is an example of one such mime-type. It also sets dst to storage on your behalf. 
