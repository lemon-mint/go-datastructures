/*
Package datastructures exists solely to aid consumers of the go-datastructures
library when using dependency managers.  Depman, for instance, will work
correctly with any datastructure by simply importing this package instead of
each subpackage individually.

For more information about the datastructures package, see the README at

	http://github.com/lemon-mint/go-datastructures

*/
package datastructures

import (
	_ "github.com/lemon-mint/go-datastructures/augmentedtree"
	_ "github.com/lemon-mint/go-datastructures/bitarray"
	_ "github.com/lemon-mint/go-datastructures/btree/palm"
	_ "github.com/lemon-mint/go-datastructures/btree/plus"
	_ "github.com/lemon-mint/go-datastructures/fibheap"
	_ "github.com/lemon-mint/go-datastructures/futures"
	_ "github.com/lemon-mint/go-datastructures/hashmap/fastinteger"
	_ "github.com/lemon-mint/go-datastructures/numerics/optimization"
	_ "github.com/lemon-mint/go-datastructures/queue"
	_ "github.com/lemon-mint/go-datastructures/rangetree"
	_ "github.com/lemon-mint/go-datastructures/rangetree/skiplist"
	_ "github.com/lemon-mint/go-datastructures/set"
	_ "github.com/lemon-mint/go-datastructures/slice"
	_ "github.com/lemon-mint/go-datastructures/slice/skip"
	_ "github.com/lemon-mint/go-datastructures/sort"
	_ "github.com/lemon-mint/go-datastructures/threadsafe/err"
	_ "github.com/lemon-mint/go-datastructures/tree/avl"
	_ "github.com/lemon-mint/go-datastructures/trie/xfast"
	_ "github.com/lemon-mint/go-datastructures/trie/yfast"
)
