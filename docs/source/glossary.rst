.. _glossary:

========
Glossary
========

.. glossary::
  :sorted:

  Hal
      Hal is a chat bot framework written in the Go programming language.

  Adapter
      chat adapter

  Store
      data storage

  Robot
      robot

  Handler
      A handler is the part of hal that evaluates incoming messages

  User
      A chat user

  Message
      A message is an incoming request that is processed by handlers

  Response
      A response is the return object that is processed by the adapter

  Envelope
      An envelope contains metadata about the message and response, used for additional processing by both handlers and adapters.

  Router
      A router is an HTTP server used for handling requests (Messages) that did not come through the chat adapter.
