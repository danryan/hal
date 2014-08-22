.. _custom_store:

=====================
Adding a Custom Store
=====================

Providing support for a new backend is fairly uncomplicated. Taking
advantage of Go's interface type, we simply need a new struct type that
implements the ``hal.Store`` interface (plus a handful of helper functions,
but we'll get to that). Let's take a look at the default memory adapter
to see how one works.

Start by declaring a new package and importing hal (and other packages
you may need).

.. code:: go

    package memory

    import (
        "fmt"
        "github.com/danryan/hal"
    )

Next, we need to define a hook that will tell hal about the store and
how to create a new one. ``hal.RegisterStore`` take two arguments: a
string to use as a identifying name, and a constructor function that
initializes and returns the store. This should go into the ``init()``
function so that it is called when the file is parsed. Doing so allows
us to ``import _`` the package for the side effect of registering our
store.

.. code:: go

    func init() {
        hal.RegisterStore("memory", New)
    }

We now need to define a ``store`` struct. Easy enough:

.. code:: go

    type store struct {
        hal.BasicStore
        data map[string][]byte
    }

Notice that we embed ``hal.BasicStore`` in our struct. This gives us a
number of extra things, namely the ability to interact with the robot.
The ``data`` field is a basic map of strings to byte-slices. We'll use
this to store and retrieve data. It wouldn't pass a `Jepsen
simulation <https://github.com/aphyr/jepsen>`__ but it's at least Web
Scale.

Time to define our constructor function. If you recall, this gets passed
to ``hal.RegisterStore`` so hal knows how to initialize our store. The
expected function signature is ``func(*hal.Robot) (hal.Store, error)``.

.. code:: go

    func New(robot *hal.Robot) (hal.Store, error) {
            // make a new store object and initialize the data field
        s := &store{
            data: map[string][]byte{},
        }

            // set the store's robot to the robot we passed as an argument.
        s.SetRobot(robot)

            // return the store object
            // if this were a more complex adapter, we would need to check for and return errors if applicable.
        return s, nil
    }

So far so good! At this point we've handled all of the setup functions
necessary for hal to register and initialize a new store, but we still
need our struct to conform to the ``hal.Store`` interface in order for
our program to compile. So let's do that now!

``Open()`` is called immediately after the adapter is initialized and
immediately before the ``robot.Run()`` function returns. This function
would generally be used to initialize a connection to an underlying
database (the [[Redis Store]], for example). We don't *use* it for our
little memory store, but it is *required*, otherwise our store won't
work as ``hal.Store``.

.. code:: go

    func (s *store) Open() error {
        return nil
    }

``Close()`` is called immediately before the adapter is shut down and
immediately after the ``robot.Stop()`` function begins. This function is
useful for closing connections to a database (much like the [[Redis
Store]] does). We have nothing to close so our function will be very
boring. Just like ``Open``, it is *required* in order to implement the
``hal.Store`` interface.

.. code:: go

    func (s *store) Close() error {
        return nil
    }

``Get`` is our way to retrieve a value from a store by a key (a
*key-value store*, if you will). It should take a string *key* and
return a byte-slice and/or an error if necessary.

.. code:: go

    func (s *store) Get(key string) ([]byte, error) {
        val, ok := s.data[key]
        if !ok {
            return nil, fmt.Errorf("key %s was not found", key) 
        }
        return val, nil 
    }

``Set`` pushes stores a value to a given key. It take a string *key*, a
byte-slice *data*, and may return an error if necessary.

.. code:: go

    func (s *store) Set(key string, data []byte) error {
        s.data[key] = data
        return nil
    }

``Delete`` removes the value referenced by a given key. It expects a
string *key*, and may return an error if necessary.

.. code:: go

    func (s *store) Delete(key string) error {
        if _, ok := s.data[key]; !ok {
            return fmt.Errorf("key %s was not found", key)
        }
        delete(s.data, key)
        return nil
    }

And we're done! Now go contribute a store for your favorite key-value backend :)
