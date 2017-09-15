Glock
=====

CLI to interact with `err-lock` Errbot plugin.

Configuration
-------------
There is a configuration file needed to set the endpoint. The configuration is found in `$HOME/.config/glock-config.json`.

An example config looks like:
```
{
  "endpoint": "http://10.10.100.10:8080/locker"
}
```

Usage
-----
The usage is very similar to the Errbot plugin

`$ glock -action lock -chest sample1`

Options:
* action: Used to lock/unlock or view locked items. Possible options are: lock, unlock, locked
* chest: The name of the chest to lock/unlock. Optional.
* username: The username that will lock/unlock the chest. Optional.


Building
--------
Create the binary and then copy it to a `bin` directory.

```
$ go build -o glock cmd/glock/main.go
$mv glock ~/bin/
```
