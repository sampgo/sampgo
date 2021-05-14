<p align="center">
<img src="https://github.com/sampgo/sampgo/blob/master/assets/img/GTASAGOMP_Stylesheet_condensed@2x.png?raw=true"/>
</p>
<p align="center">sampgo is a SAMP gamemode SDK written in Go, based on Zeex's SAMPGDK.</p>
<p align="center"><strong>Both Linux and Windows supported!</strong></p>


## NOTE
The only working branch is `experimental` - for that to work, you must clone the the base repo, edit the `go get ...` line and add `@experimental` to the end of it.  This will force Go to download the experimental branch (over the master or main branch) and requires GO111MODULE env to be set to `on`.  The provided batch scripts in the base gamemode repository already does it for you and is highly recommended to use that for now!  If you have any issues and wish to update the sampgo your gamemode is built with, please delete the Go cache and edit your `go.mod` file.

## Quickstart
Clone the `https://github.com/sampgo/base.git` repo.

Or check the wiki for more information.

## Installation
```
go get -u github.com/sampgo/sampgo
```

## Credits
- AliLogic for quite a few things.

- dakyskye for helping out with some event handling logic, and a lot of motivation.

- 00face for his amazing graphic artistry! If you would like to support him in anyway, please check out [his Ko-fi](https://ko-fi.com/00face)!

- [JetBrains](https://www.jetbrains.com/) for providing us with access to the All Products Pack.

## API coverage
See [here!](https://github.com/sampgo/sampgo/wiki/API-coverage)

## Support
Join our [Discord](https://discord.gg/6ke4MEkJGB), we will be more than happy to aid you if possible.
