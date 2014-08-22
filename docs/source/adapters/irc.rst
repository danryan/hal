===
IRC
===

A simple adapter for use with any IRC server.

Setup
~~~~~

You will need an IRC user and preferred server. If your server requires
a password, be sure to provide it using the environment variable below.

Usage
~~~~~

.. code:: go

    // blank import to register adapter
    import _ "github.com/danryan/hal/adapter/irc"

Configuration
~~~~~~~~~~~~~

::

    HAL_ADAPTER=irc                 # The adapter
                                    # Default: shell
    HAL_IRC_USER=blah               # IRC username
                                    # Default: none (required)
    HAL_IRC_PASSWORD=sekret         # IRC password if required
                                    # Default: none (optional)
    HAL_IRC_NICK=hal                # IRC nick
                                    # Default: HAL_IRC_USER (optional)
    HAL_IRC_SERVER=irc.freenode.net # IRC server
                                    # Default: none (required)
    HAL_IRC_PORT=6667               # IRC server port
                                    # Default: 6667
    HAL_IRC_CHANNELS="#foo,#bar"    # Comma-separate list of channels to join after connecting
                                    # Default: none (required)
    HAL_IRC_USE_TLS=false           # Use an encrypted connection

