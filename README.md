# OSS

Tool for managing Open Source assets.

## Why ?

Managing Open Source assets is a hard task in any project ; we often download awesome librairies, images or medias, but 
where to keep trace of all theses files ? How can I manage Open Source licenses of my project ? 

That's the problem oss try to resolve today.

## Usage

On first run, please execute
    
    oss init
    
Then, you can use the following commands:

+ **status**: list assets of the project
+ **add** `<license>` `<file>` `[<description>]`: register new file
+ **rm** `<file>`: unregister file
+ **show** `<file>`: display information about the given file

Or to get information about SPDX licenses :

+ **licenses**: list available licenses (base on the [SPDX license list]((http://spdx.org/licenses/)))
+ **search** `<expression>`: Search licenses matching `expression`
+ **update**: update SPDX license list

Keep in mind that **the license identifier should be registered in [SPDX License list](http://spdx.org/licenses/)**

## Copyright

MIT License. Copyright (c) Jean-François Lépine. See LICENSE for details.

## Contributing

You need to install:

+ [go](https://golang.org/)
+ [semver](https://github.com/flazz/semver/)

Please never run `go build` manually. You should run `make build`.
