package bloomfilter

// TODO: add logrus

import (
	"crypto/md5"
	"errors"
	"fmt"
)

const (
	// HashSize defines checksum length of the hash
	HashSize = 16
)

var (
	// ErrBitTooLong defines the error if given num of bits too long.
	ErrBitTooLong = errors.New("bits length too long.")
)

// BloomFilterImpl implements the BloomFilter interface.
type BloomFilterImpl struct {
	bits  int
	store KVStore
	hash  func(data []byte) [HashSize]byte
	// TODO: add mutex to handle read/write lock
}

// New instantiate the BloomFilter with given length of bits
func New(numBits int) (BloomFilter, error) {
	if numBits > 32 {
		return nil, ErrBitTooLong
	}
	// NOTE: current new func returns BloomFilterImpl which implements the
	// BloomFilter interface.
	return &BloomFilterImpl{
		bits: numBits,
		// TODO: input different or better hash, such as murmur hash
		hash:  md5.Sum,
		store: NewKVStore(),
	}, nil
}

// Add implements the add function for filter interface
func (f *BloomFilterImpl) Add(b []byte) (bool, error) {
	if err := f.store.Set(f.genHash(b), b); err != nil {
		if err != ErrKeyExist {
			return false, err
		}
		return true, err
		// NOTE: any log/handle the key exist case. Here we do not need to do
		// anything and it will still add successfully.
	}
	return true, nil
}

// Check implements the check function for filter interface
func (f *BloomFilterImpl) Check(b []byte) (bool, error) {
	_, err := f.store.Get(f.genHash(b))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *BloomFilterImpl) genHash(b []byte) string {
	ret := fmt.Sprintf("%x", f.hash(b))[:f.bits]
	return ret
}
