## go-scrdt

This project aims to implement state base Conflict-free Replicated Data Types in Go.

All objects are serialized to JSON representation.

Currently implemented:

* GSet (Grow-only Set, you can only add elements to this set)
* 2PSet (Two Pahse Set, you can once add and once remove element from the set)
* LCounter (Lazy Counter, contains maximum value from network, not real CRDT)
* GCounter (Grow-only Counter, true CRDT coutner)