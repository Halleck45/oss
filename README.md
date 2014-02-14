# OSS

Tool for managing Open Source assets.

## Why ?

Managing Open Source assets is a hard task in any project ; we often download awesome librairies, images or medias, but 
where to keep trace of all theses files ? How can I manage Open Source licenses of medias in my project ? 

That's the problem oss try to resolve today.

![simple overview of oss](./doc/overview.gif)

## Installation

Download binary for your platform:

+ Linux Debian: [64bits](http://dl.bintray.com/halleck45/OSS/oss_linux_amd64), [32bits](http://dl.bintray.com/halleck45/OSS/oss_linux_386)
+ Windows: [64bits](http://dl.bintray.com/halleck45/OSS/oss_windows_amd64.exe), [32bits](http://dl.bintray.com/halleck45/OSS/oss_windows_386.exe)
+ [Others platforms](http://dl.bintray.com/halleck45/OSS/)

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


    clone git@github.com:Halleck45/OSS.git
    cd OSS
    make init && make install
    ( ... your development ...)
    make build
    
Remember to keep tests up to date (and, of course, to run them with `make test`)

Builds are hosted by [bintray](https://bintray.com). Today, only me (Halleck45) has permission to publish binaries.