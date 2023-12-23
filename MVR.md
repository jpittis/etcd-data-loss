## Dependencies

- This repository's `write-all-zeroes` script
- etcd
- etcd-dump-logs
- perl (for simulating data corruption)

## Scenario

1. Start a fresh etcd.

```
$ etcd
```

2. Write a key value pair to etcd containing 1 KB of zero bytes to etcd.

```
$ cd write-all-zeroes
$ go run main.go
```

3. Kill etcd.

```
$ kill -9 $(pidof etcd)
```

4. Simulate disk corruption of the record you previously wrote.

```
$ cd default.etcd/member/wal
$ perl -pi -e 's/all-zeroes/all-zeroez/g' 0000000000000000-0000000000000000.wal
```

5. Restarting etcd will lead to an automated repair that truncates the WAL.

```
$ etcd
...
{"level":"info","ts":"2023-12-23T13:58:47.196297-0500","caller":"wal/repair.go:41","msg":"repairing"
,"path":"default.etcd/member/wal/0000000000000000-0000000000000000.wal"}
{"level":"info","ts":"2023-12-23T13:58:47.22622-0500","caller":"wal/repair.go:97","msg":"repaired","
path":"default.etcd/member/wal/0000000000000000-0000000000000000.wal","error":"unexpected EOF"}
{"level":"info","ts":"2023-12-23T13:58:47.226469-0500","caller":"etcdserver/storage.go:110","msg":"r
epaired WAL","error":"unexpected EOF"}
{"level":"warn","ts":"2023-12-23T13:58:47.226572-0500","caller":"wal/util.go:90","msg":"ignored file
 in WAL directory","path":"0000000000000000-0000000000000000.wal.broken"}
...
```

6. You can see the lack of the record you wrote on the WAL.

```
$ etcd-dump-logs default.etcd
Snapshot:
empty
Start dumping log entries from snapshot.
{"level":"warn","msg":"ignored file in WAL directory","path":"0000000000000000-0000000000000000.wal.broken"}
WAL metadata:
nodeID=8e9e05c52164694d clusterID=cdf818194e3a8c32 term=3 commitIndex=6 vote=8e9e05c52164694d
WAL entries: 6
lastIndex=6
term	     index	type	data
   1	         1	conf	method=ConfChangeAddNode id=8e9e05c52164694d
   2	         2	norm	
   2	         3	norm	method=PUT path="/0/members/8e9e05c52164694d/attributes" val="{\"name\":\"default\",\"clientURLs\":[\"http://localhost:2379\"]}"
   2	         4	norm	method=PUT path="/0/version" val="3.5.0"
   3	         5	norm	
   3	         6	norm	method=PUT path="/0/members/8e9e05c52164694d/attributes" val="{\"name\":\"default\",\"clientURLs\":[\"http://localhost:2379\"]}"

Entry types (Normal,ConfigChange) count is : 6
```

7. But if you query etcd, the key value pair will still be present.

```
$ etcdctl get all-zeroes
all-zeroes
```
