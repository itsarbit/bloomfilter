package bloomfilter

// TODO: add logrus

// BloomFilter defines the general interface bloom filter needed
// ref: http://codekata.com/kata/kata05-bloom-filters/
type BloomFilter interface {
	// Add add the object to the bloom filter
	Add(object []byte) (done bool, err error)
	// Check checks if an object is already in the storage:
	// if false: then it is sure it is not in the store
	// if true: then it is POSSIBLE that the object is in the store
	Check(object []byte) (done bool, err error)
}
