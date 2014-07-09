# HAL

A chat bot in Go, now with 100% less CoffeeScript!

## Getting started

"Good morning, Dr. Chandra. This is HAL. I'm ready for my first lesson."

Hal is Go all the way down, and uses standard packages wherever possible. For an idea of how you can use it, look at [a simple example](examples/simple/main.go), or [a more complex example](examples/complex/main.go). API Documentation is available [here](http://godoc.org/github.com/danryan/hal).

## Is it any good?

[Probably not.](http://news.ycombinator.com/item?id=3067434)

## Configuration

Hal doesn't have any command line options; instead we utilize environment variables exclusively. See [configuration options](https://github.com/danryan/hal/wiki/Configuration) for an exhaustive list of options.

## Adapters

Adapters are what hal uses to integrate with your chat services. Visit the following links for detailed information. 

* [Shell](https://github.com/danryan/hal/wiki/Shell-Adapter)
* [Campfire](https://github.com/danryan/hal/wiki/Campfire-Adapter)
* [Hipchat](https://github.com/danryan/hal/wiki/Hipchat-Adapter)
* [IRC](https://github.com/danryan/hal/wiki/IRC-Adapter)
* [Slack](https://github.com/danryan/hal/wiki/Slack-Adapter)

If you don't see your preferred chat service listed, [let me know!](https://github.com/danryan/hal/issues/new)

## Stores

Stores are the brains of hal, used to persist long-term data and retrieve when needed. Review the links below for your store of choice.

* [Memory](https://github.com/danryan/hal/wiki/Memory-Store)
* [Redis](https://github.com/danryan/hal/wiki/Redis-Store)

 If you think support for a particular store is needed, [just open a ticket :beers:](https://github.com/danryan/hal/issues/new)

## Bugs, features, rants, hate-mail, etc.

Please use [the issue tracker](https://github.com/danryan/hal/issues) for development progress tracking, feature requests, or bug reports. Thank you! :heart:

## License

Copyright 2014 Applied Awesome LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
