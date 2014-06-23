* command line configuration
* docs!
* example scripts (handlers)
* command not found handler
* change enter/leave to join/part ?
* figure out how to autoregister new commands
* change handler signature to support `ResponseWriter` and `Message`, i.e.:
```
func (rw hal.ResponseWriter, msg *hal.Message)
```

## Done
* rename command back to listener
