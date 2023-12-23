This repository demonstrates a scenario in which etcd's repair on bootstrap behavior may
incorrectly identify the disk corruption of a committed entry as a torn write of an
uncommitted entry, leading to truncation and loss of committed data. The scenario is
unlikely to occur in the real world due to it requiring a record containing more than 512
zero bytes on the WAL to experience disk corruption. That being said, there may be users
writing records of this shape to etcd at scale who are at risk. It also demonstrates
general unsoundness in etcd's torn write detection and automated repair workflow which may
need to be inspected more carefully for other unintended consequences.

- [MVR.md](MVR.md) shows how to minimally reproduce the issue locally using a single etcd.
- [PLAYGROUND.md](PLAYGROUND.md) shows potential ramifiactions of the inconsistency using
  a 3 node etcd cluster and well timed network partitions.
