.. _configuration:

=============
Configuration
=============

Hal doesn't have any command line options. Instead we utilize
environment variables exclusively, allowing you to use hal in more
flexible ways.

::

    PORT=9000           # The port on which the HTTP server will listen.
                        # Default: 9000
    HAL_NAME=hal        # The name to which Hal will respond.
                        # Default: hal
    HAL_ADAPTER=shell   # The adapter name.
                        # Default: shell
                        # Options: shell, slack, irc
    HAL_LOG_LEVEL=info  # The level of logging desired.
                        # Default: info
                        # Options: info, debug, warn, error, critical

