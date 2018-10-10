# DBFS

Dbfs starts a 9P2000 file server representing a db as a file system.

# 9P2000

A 9P2000 server is an agent that provides one or more hierarchical
file systems -- file trees -- that may be accessed by processes. A
server responds to requests by clients to navigate the hierarchy, and
to create, remove, read, and write files.

# Usage

    TODO

# Schema
//srv{nm, tp, lsn} Plan9 file server, serve 9P Rx->Tx. tp = ver
//node{nm, tp, fid, qid, ver, prn, chld} File tree
oradb{}
mssqldb{}
sqlitedb{}

# Links
A: Just an article with an example
https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

G: A well-style go project, not small, and not so big as perkeep
https://github.com/gopherjs/gopherjs/wiki/Developer-Guidelines

D: Effective go
https://golang.org/doc/effective_go.html

D: Go review
https://github.com/golang/go/wiki/CodeReviewComments
