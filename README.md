# HAL

A chat bot in Go, now with 100% less CoffeeScript!

## Getting started

"Good morning, Dr. Chandra. This is HAL. I'm ready for my first lesson."

Hal is Go all the way down, and uses standard packages wherever possible. It's a bit rough around the edges right now. For an idea of how you can use it, look at [a simple example](examples/simple/main.go), or [a more complex example](examples/complex/main.go). API Documentation is available [here](http://godoc.org/github.com/danryan/hal).

## Is it any good?

[Probably not.](http://news.ycombinator.com/item?id=3067434)

## Configuration

Hal doesn't have any command line options, instead opting to use environment variables exclusively.

```
PORT=9000           # The port on which the HTTP server will listen.
                    # Default: 9000
HAL_NAME=hal        # The name to which Hal will respond.
                    # Default: hal
HAL_ADAPTER=shell   # The adapter name.
                    # Default: shell
                    # Options: shell, slack
HAL_LOG_LEVEL=info  # The level of logging desired.
                    # Default: info
                    # Options: info, debug, warn, error, critical
```

## Adapters

### Slack

Hal uses Slack's hubot integration. Currently Hal will listen in on all channels. In the future, you'll be able to specify channels by either a whitelist or blacklist. In addition, private groups do not work. This is a limitation of Slack's API which may change in the future. Support for using their IRC gateway will be implemented shortly. 

Start by adding the Hubot integration for your team (if you haven't done so). Then, set the following environment variables when starting up your bot:

```
HAL_ADAPTER=slack               # The adapter
HAL_SLACK_TOKEN=blah            # Your integration token
                                # Default: none
HAL_SLACK_TEAM=acmeinc          # Your Slack subdomain (<team>.slack.com)
                                # Default: none
HAL_SLACK_BOTNAME=HAL           # The username Hal will send replies as
                                # Default: HAL_NAME
HAL_SLACK_ICON_EMOJI=":poop:"   # The emoji shortcut used as the response icon
                                # Default: none
HAL_SLACK_CHANNELS=""           # not yet implemented
HAL_SLACK_CHANNELMODE=""        # not yet implemented
HAL_SLACK_LINK_NAMES=""         # not yet implemented
```

No additional configuration should be required. 

### Shell

Hal comes with a default shell adapter, useful for testing your response handlers locally. It has no special configuration variables.

## TODO

### Definitely

* tests :O
* help (and help parsing)
* documentation

### Maybe

* command not found handler
* change enter/leave to join/part
* autoregister new commands
