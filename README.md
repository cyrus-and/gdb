gdb
===

Package `gdb` provides a convenient way to interact with the GDB/MI interface. The
methods offered by this module are very low level, the main goal is to avoid the
tedious parsing of the MI2 line based text interface.

The objects returned as a result of the commands or as asynchronous
notifications are generic Go maps suitable to be converted to JSON format with
`json.Marshal()`. The fields present in such objects are blindly added
according to the records returned from GDB (see the
[command syntax][mi2-syntax]): tuples are `map[string]interface{}` and lists are
`[]interface{}`. There are a couple of exceptions to this:

- the record class, where present, is represented by the `"class"` field;

- the record type is represented using the `"type"` field as follows:
    - `+`: `"status"`
    - `=`: `"notify"`
    - `~`: `"console"`
    - `@`: `"target"`
    - `&`: `"log"`

- the optional result list is stored into a tuple under the `"payload"` field.

Installation
------------

    go get github.com/cyrus-and/gdb

Documentation
-------------

[GoDoc][godoc]

Caveats
-------

I/O operations of the target program happens through a pseudoterminal obtained
using the [pty][pty] package which basically uses the `/dev/ptmx` on *nix
systems to request new terminal instances.

There are some unclear behaviors on Mac OS X. Calling `gdb.Write()` when the
target program is not running is a no-op, on Linux instead writes are somewhat
buffered and delivered later. Likewise, `gdb.Read()` may returns EOF even though
there is actually data to read, a solution may be keep trying.

Resources
---------

- [The `GDB/MI` Interface][gdb-mi]

[mi2-syntax]: https://sourceware.org/gdb/onlinedocs/gdb/GDB_002fMI-Output-Syntax.html
[godoc]: https://godoc.org/github.com/cyrus-and/gdb
[pty]: https://github.com/kr/pty
[gdb-mi]: https://sourceware.org/gdb/onlinedocs/gdb/GDB_002fMI.html
