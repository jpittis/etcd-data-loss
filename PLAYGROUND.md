## Dependencies

- This repository's `write-all-zeroes` script
- https://github.com/jpittis/etcd-playground
- Docker

## Scenario

1. Clone etcd-playground and start docker-compose.

```
$ docker-compose up
```

2. Write a normal record to etcd1.

```
$ etcdctl put foo bar
OK
```

3. Partition etcd2 from the rest of the cluster.

```
$ docker ps | grep manager
506860e40eeb   etcd-playground-manager
$ docker exec -it 506860e40eeb /bin/bash
$ curl -s -XPOST 'etcd2:3333/network?dev=etcd120&loss=100'
$ curl -s -XPOST 'etcd2:3333/network?dev=etcd230&loss=100'
```

4. Write a 1 KB zero record to etcd1.

```
$ cd write-all-zeroes
$ go run main.go
```

5. Partition etcd3 from the rest of the cluster.

```
$ curl -s -XPOST 'etcd3:3333/network?dev=etcd230&loss=100'
$ curl -s -XPOST 'etcd3:3333/network?dev=etcd310&loss=100'
```

6. Take down etcd1, corrupt it's WAL and restart it.

TODO

7. Unpartition etcd2 from etcd1.

TODO

8. Confirm both instances have quorum but are inconsistent.

TODO

9. Write a new record to etcd1, replacing the Raft index of the all-zeroes entry.

TODO

10. Unpartition etcd3, and confirm complete data-loss of the all-zeroes entry.

TODO
