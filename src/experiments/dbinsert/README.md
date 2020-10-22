# An experiment with MySQL and caching of prepared statements

This contains code for an experiment I wanted to do:
> Is it faster to do a single query with a bulk insert or is it faster to prepare a statement and do many small queries.

## Run the docker image with the DB
```
./run-db.sh
```

## Run the benchmark
```
$ bazel run :dbinsert_test -- -test.bench=. -test.run=50
...
BenchmarkBulk-4              159           8075570 ns/op
BenchmarkSingle-4            159           7695383 ns/op
```

This shows that we may be winning for this anecdotal case with 1 element in the slice.  This may give us our error.  When we run the same benchmark with 100 rows in the input slice:
```
BenchmarkBulk-4               78          13710611 ns/op
BenchmarkSingle-4             33          31782103 ns/op

```

And a 1000:
```
BenchmarkBulk-4               20          50391153 ns/op
BenchmarkSingle-4              5         235678415 ns/op
```

This result was expected, but I was not expecting such a large difference in this case.
