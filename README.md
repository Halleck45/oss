# OSS

Tool for managing Open Source assets.

## Why ?

Managing Open Source assets is a hard task in any project ; we often download awesome librairies, images or medias, but 
where to keep trace of all theses files ? How can I manage Open Source licenses of medias in my project ? 

That's the problem oss try to resolve today.

![simple overview of oss](./doc/overview.gif)

## Installation

Download the binary for your platform from the [latest release](https://github.com/Halleck45/oss/releases/latest).

## Usage

On first run, please execute
    
```bash
oss init
```
    
Then, you can use the following commands:

+ `oss `**`status`**: list assets of the project
+ `oss `**`add`**` <license> <file> [<description>]`: register new file
+ `oss `**`rm`**` <file>`: unregister file
+ `oss `**`show`**` <file>`: display information about the given file

Or to get information about SPDX licenses :

+ `oss `**`licenses`**: list available licenses (base on the [SPDX license list]((http://spdx.org/licenses/)))
+ `oss `**`search`**` <expression>`: Search licenses matching `expression`
+ `oss `**`update`**: update SPDX license list

Keep in mind that **the license identifier should be registered in [SPDX License list](http://spdx.org/licenses/)**

## Copyright

MIT License. Copyright (c) Jean-François Lépine. See LICENSE for details.

## Contributing


Cf. [CONTRIBUTING.md](./.github/CONTRIBUTING.md)
    
Remember to keep tests up to date (and, of course, to run them with `make test`)