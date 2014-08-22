=======
Hipchat
=======

Setup
~~~~~

Hal uses Hipchat’s XMPP gateway and so requires a user account to
integrate with Hipchat. Be sure to create one before configuring the
adapter. You will need the XMPP credentials, which can be found at
https://my.hipchat.com/account/xmpp.

Usage
~~~~~

.. code:: go

    // blank import to register adapter
    import _ "github.com/danryan/hal/adapter/hipchat"

Configuration
~~~~~~~~~~~~~

Set the following environment variables according to your needs.

``HAL_ADAPTER``
^^^^^^^^^^^^^^^

To use the Hipchat adapter, set ``HAL_ADAPTER`` to ``hipchat``.

``HAL_HIPCHAT_USER``
^^^^^^^^^^^^^^^^^^^^

The username is the first part of your XMPP JID before the ``@`` sign.
E.g., if your JID is ``134273_971874@chat.hipchat.com``, then
``HAL_HIPCHAT_USER`` should be ``134273_971874``.

-  **Default:** none
-  **Required:** false
-  **Example:** ``HAL_HIPCHAT_USER=134273_971874``

``HAL_HIPCHAT_PASSWORD``
^^^^^^^^^^^^^^^^^^^^^^^^

The password is the same as the Hipchat user’s password.

-  **Default:** none
-  **Required:** true
-  **Example:** ``HAL_HIPCHAT_PASSWORD=supersekretpassword``

``HAL_HIPCHAT_ROOMS``
^^^^^^^^^^^^^^^^^^^^^

This is a comma-separated list of rooms to join. Note that Hipchat has
two ways of specifying rooms: a human-readable format (ex. ``general``);
and an XMPP format (ex. ``134273_general``). Hal expects the former
human-readable format at this time due to a limitation of the
third-party Hipchat package presently used. The rooms are case sensitive
as well.

Hal will not fail if no rooms are specified, though hal will also not
join any rooms if this is left blank.

-  **Default:** none
-  **Required:** false
-  **Example:** ``HAL_HIPCHAT_ROOMS="general,room with spaces,random"``

``HAL_HIPCHAT_RESOURCE``
^^^^^^^^^^^^^^^^^^^^^^^^

This is an optional setting. The default, ``bot``, prevents the channel
history from being sent and thus prevents hal from parsing possibly
already handled messages. If changed from the default, ``bot``, channel
history will be sent. It is recommended that the default be left unless
you need channel history.

-  **Default:** bot
-  **Required:** false
-  **Example:** ``HAL_HIPCHAT_RESOURCE=something-other-than-bot``
