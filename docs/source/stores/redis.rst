=====
Redis
=====

Setup
~~~~~

The Redis store requires an available Redis server. Authentication and
custom databases are not supported at this time. Please `open an
issue <https://github.com/danryan/hal/issues>`__ if you need this support!

Configuration
~~~~~~~~~~~~~

``HAL_STORE``
^^^^^^^^^^^^^

Set to ``redis``

-  Default: ``memory``
-  Example:

   ::

       HAL_STORE=redis

``HAL_REDIS_URL``
^^^^^^^^^^^^^^^^^

The Redis server URL

-  Default: ``localhost:6367``
-  Example:

   ::

       HAL_REDIS_URL=redis.example.com:6379

``HAL_REDIS_NAMESPACE``
^^^^^^^^^^^^^^^^^^^^^^^

Set a namespace to prepend to all keys

-  Default: ``hal``
-  Example

   ::

       HAL_REDIS_NAMESPACE=foo
       # sets all keys to "foo:<key>"

Usage
~~~~~

.. code:: go

    // blank import to register adapter
    import _ "github.com/danryan/hal/store/redis"

