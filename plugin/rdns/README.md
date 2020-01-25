# rdns

## Name

*rdns* - collects A and AAAA answers and provides an interface to query IP-to-host mappings

## Description

The *rdns* plugin observes DNS answers and maps IP addresses to hostnames. You can query the plugin to find hostnames that IP addresses have resolved to.

## Syntax

~~~
rdns LISTEN_HOST LISTEN_PORT [ZONES...]
~~~

* **LISTEN_HOST** host interface on which to listen
* **LISTEN_PORT** is the port on which to listen

## Examples

Listen on 127.0.0.53:8853 for any domain

~~~ corefile
. {
    rdns 127.0.0.53 8853 .
}
~~~

Listen on a UNIX domain socket

~~~ corefile
. {
    rdns /path/to/socket.sock
}
~~~

## Notes
- If you use TCP for the interface server, ensure that you have a rule allowing the connection or that you are able to allow the connection
  - Using a UNIX stream socket will avoid this issue