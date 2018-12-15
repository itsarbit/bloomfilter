package bloomfilter

// TODO: add logrus

import "errors"

var (
	// ErrKeyNotFound defines a key is not found in the store.
	ErrKeyNotFound = errors.New("key not found.")
	// ErrKeyExist defines if key already exist in the store.
	ErrKeyExist = errors.New("key already exist.")
)

// KVStore defines the general key value store interface
type KVStore interface {
	// Get gets the value from store with given key
	Get(key string) (val []byte, err error)
	// Set sets the value to store
	// NOTE: if error returns ErrKeyExist, it is not necessarily be a real
	// error case. It is intended to be a special error to inform the caller
	// that the previous data is been overwritten instead of new key-value
	// pair.
	Set(key string, val []byte) error
}

// NewKVStore instantiate a KVStore
func NewKVStore() KVStore {
	return &localStore{
		m: map[string][]byte{},
	}
}

// localStore implements in memory storage with KVStore interface.
// TODO: kvstore can also be implemented by bit slice
type localStore struct {
	m map[string][]byte
	// TODO: add mutex to handle map read/write lock
}

// Get implements the KVStore Get interface
func (l *localStore) Get(key string) ([]byte, error) {
	b, ok := l.m[key]
	if !ok {
		return []byte{}, ErrKeyNotFound
	}
	return b, nil
}

// Set implements KVStore Set function
func (l *localStore) Set(key string, b []byte) error {
	_, ok := l.m[key]
	l.m[key] = b
	if ok {
		return ErrKeyExist
	}
	return nil
}
