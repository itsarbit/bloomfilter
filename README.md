Bloom Filter
----

# Introduction
A Bloom filter is a space-efficient probabilistic data structure.
that is used to test whether an element is a member of a set. 
False positive matches are possible, but false negatives are not.
In other words, a query returns either "possibly in set" or "definitely not in set".

(reference by Wikipedia)


# Usage

```
bf, _ := New(NumBits)
done, _ := bf.Add([]byte("some data"))
exist, _ := bf.Check([]byte("some data"))
```

# Test

Short test case:
`go test -test.short -cover`
```
PASS
coverage: 92.0% of statements
ok  	/bloomfilter	0.022s
```

Long test with given word list:
`go test -cover`
```
PASS
coverage: 91.7% of statements
ok  	/bloomfilter	0.841s

```

# Reference
http://codekata.com/kata/kata05-bloom-filters/

