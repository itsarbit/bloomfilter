package bloomfilter

// TODO: mock Store interface and test Store independently

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testConstNumBits = 5
)

var (
	wordmap         = map[string]struct{}{}
	testEmptyData   = []byte{}
	testStringData  = []byte("this is test")
	testStringData2 = []byte("this is test2")
	testRandomWord  = []byte("dfadkqrhkf")
)

type BloomFilterSuite struct {
	suite.Suite

	bf BloomFilter
}

func TestBloomFilterSuite(t *testing.T) {
	suite.Run(t, &BloomFilterSuite{})
}

func (s *BloomFilterSuite) SetupTest() {
	bf, err := New(testConstNumBits)
	s.bf = bf
	assert.NoError(s.T(), err)
}

func (s *BloomFilterSuite) TestAdd() {
	suc, err := s.bf.Add(testEmptyData)
	assert.True(s.T(), suc)
	assert.NoError(s.T(), err)
	suc, err = s.bf.Add(testEmptyData)
	assert.True(s.T(), suc)
	assert.Equal(s.T(), err, ErrKeyExist)
	suc, err = s.bf.Add(testStringData)
	assert.True(s.T(), suc)
	assert.NoError(s.T(), err)
}

func (s *BloomFilterSuite) TestGet() {
	suc, err := s.bf.Add(testStringData)
	assert.True(s.T(), suc)
	assert.NoError(s.T(), err)
	exist, err := s.bf.Check(testStringData)
	assert.NoError(s.T(), err)
	assert.True(s.T(), exist)
	assert.NoError(s.T(), err)
	exist, err = s.bf.Check(testStringData2)
	assert.False(s.T(), exist)
	assert.Equal(s.T(), err, ErrKeyNotFound)
	exist, err = s.bf.Check(testEmptyData)
	assert.False(s.T(), exist)
	assert.Equal(s.T(), err, ErrKeyNotFound)
}

type BloomFilterLongSuite struct {
	suite.Suite

	bf BloomFilter
}

func TestBloomFilterLongSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the long test")
		return
	}
	suite.Run(t, new(BloomFilterLongSuite))
}

func (s *BloomFilterLongSuite) SetupSuite() {
	file, err := os.Open("./wordlist.txt")
	defer file.Close()
	assert.NoError(s.T(), err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordmap[scanner.Text()] = struct{}{}
	}

}

func (s *BloomFilterLongSuite) SetupTest() {
	bf, err := New(testConstNumBits)
	s.bf = bf
	assert.NoError(s.T(), err)
	var suc bool
	for k, _ := range wordmap {
		suc, _ = s.bf.Add([]byte(k))
		assert.True(s.T(), suc)
	}
}

func (s *BloomFilterLongSuite) TestCheck() {
	var exist bool
	var err error
	for k, _ := range wordmap {
		exist, err = s.bf.Check([]byte(k))
		assert.True(s.T(), exist)
		assert.NoError(s.T(), err)
	}
	exist, err = s.bf.Check([]byte(testRandomWord))
	assert.False(s.T(), exist)
	assert.Error(s.T(), err)

}
