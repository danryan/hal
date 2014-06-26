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

By default, Hal uses Slack's hubot integration. Currently Hal will listen in on all public channels. In the future, you'll be able to specify channels by either a whitelist or blacklist. Private groups require the IRC gateway to work around a current limitation of the Slack API. See [Using IRC](#irc-gateway). The IRC gateway is author's prefered method as your bot will automatically join all channels and groups it belongs to, and removing Hal from a room is as simple as a `/kick hal` command. Some advanced features like 
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

#### Using IRC Gateway<a name="irc-gateway"></a>

The default integration only works with public chats. If you want hal to listen in on private chats, you must utilize the IRC gateway. You'll need a real user for hal, so be mindful of the username you choose for it and make sure you configure your bot to use that name so it can login to the IRC gateway. When enabled, hal will only use the IRC gateway to listen for messages. Hal can be configured to either respond using the API or the IRC gateway.

1. Enable the IRC gateway in [the admin settings interface](https://revily.slack.com/admin/settings)
    * Choose "Enable IRC gateway (SSL only)". You don't want your private messages sent unencrypted.
2. [Register](https://my.slack.com/signup) a new user
3. Sign in as this new user
4. Capture your new [IRC credentials](https://my.slack.com/account/gateways)
5. Set the following environment variables

```
HAL_SLACK_IRC_ENABLED           # Enable the Slack IRC listener
                                # Default: 0
                                # Options: 0, 1  ; 0 is disabled, 1 is enabled
HAL_SLACK_IRC_PASSWORD          # The IRC gateway password
                                # Default: none (required)
HAL_SLACK_RESPONSE_METHOD       # The method by which hal will respond to a message.
                                # The irc option requires that the IRC gateway be configured
                                # Default: http
                                # Options: http, irc
```

For more information, please see the following link:
* [Connecting to Slack over IRC and XMPP](https://slack.zendesk.com/hc/en-us/articles/201727913-Connecting-to-Slack-over-IRC-and-XMPP)

### Shell

Hal comes with a default shell adapter, useful for testing your response handlers locally. It has no special configuration variables.

## Bugs, features, rants, hate-mail, etc.

This section has been converted to issues. Please use (the issue tracker)[https://github.com/danryan/hal/issues) for development progress tracking, feature requests, or bug reports. Thank you! :heart:
