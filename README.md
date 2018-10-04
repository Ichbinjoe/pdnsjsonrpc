# pdns-jsonrpc

This golang package implements the PowerDNS Remote JSON-RPC backend for
implementation. It avoids reducing flexibility with responses, while trying to
reduce the final implementation surface area, abstracting away redundant dns
features.

It also supports permission based support for methods, such as extracting DNSSEC
keys while still allowing basic public knowledge access.
