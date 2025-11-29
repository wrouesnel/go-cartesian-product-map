package cartesian_test

import (
	"testing"

	"github.com/wrouesnel/go-cartesian-product-map"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FunctionSuite struct{}

var _ = Suite(&FunctionSuite{})

func (s *FunctionSuite) TestInterfaceIter(c *C) {
	x := map[string][]interface{}{
		"some_key":    {1, 2, "c"},
		"another_key": {"ten", "nine", "eight"},
	}

	y := map[string][]interface{}{
		"b_key": {10, 11, 12},
	}

	z := map[string][]interface{}{
		"key-c": {"test"},
	}

	a := cartesian.Iter(x, y, z)

	// receive products through channel
	results := []map[string]interface{}{}
	for product := range a {
		results = append(results, product)
	}

	// Unordered Output:
	// map[another_key:eight b_key:12 key-c:test some_key:c]
	// map[another_key:eight b_key:12 key-c:test some_key:2]
	// map[another_key:eight b_key:12 key-c:test some_key:1]
	// map[another_key:ten b_key:12 key-c:test some_key:2]
	// map[another_key:nine b_key:12 key-c:test some_key:1]
	// map[another_key:eight b_key:10 key-c:test some_key:1]
	// map[another_key:eight b_key:11 key-c:test some_key:1]
	// map[another_key:ten b_key:10 key-c:test some_key:2]
	// map[another_key:ten b_key:11 key-c:test some_key:2]
	// map[another_key:nine b_key:10 key-c:test some_key:1]
	// map[another_key:nine b_key:11 key-c:test some_key:1]
	// map[another_key:nine b_key:12 key-c:test some_key:2]
	// map[another_key:nine b_key:10 key-c:test some_key:2]
	// map[another_key:nine b_key:11 key-c:test some_key:2]
	// map[another_key:ten b_key:12 key-c:test some_key:1]
	// map[another_key:eight b_key:10 key-c:test some_key:2]
	// map[another_key:ten b_key:12 key-c:test some_key:c]
	// map[another_key:ten b_key:10 key-c:test some_key:c]
	// map[another_key:eight b_key:11 key-c:test some_key:2]
	// map[another_key:ten b_key:10 key-c:test some_key:1]
	// map[another_key:ten b_key:11 key-c:test some_key:1]
	// map[another_key:ten b_key:11 key-c:test some_key:c]
	// map[another_key:nine b_key:12 key-c:test some_key:c]
	// map[another_key:nine b_key:10 key-c:test some_key:c]
	// map[another_key:nine b_key:11 key-c:test some_key:c]
	// map[another_key:eight b_key:10 key-c:test some_key:c]
	// map[another_key:eight b_key:11 key-c:test some_key:c]
}
