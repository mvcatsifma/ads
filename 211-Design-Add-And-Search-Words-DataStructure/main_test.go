package p211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	t.Run("LeetCode example", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")

		// Expected outputs: [null,null,null,null,false,true,true,true]
		//assert.False(t, wd.Search("pad")) // false
		assert.True(t, wd.Search("bad")) // true
		assert.True(t, wd.Search(".ad")) // true - matches bad, dad, mad
		assert.True(t, wd.Search("b..")) // true - matches bad
	})

	t.Run("add and search exact word", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")

		assert.True(t, wd.Search("bad"))
		assert.True(t, wd.Search("dad"))
		assert.True(t, wd.Search("mad"))
		assert.False(t, wd.Search("pad"))
	})

	t.Run("search with single wildcard dot", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")

		assert.True(t, wd.Search(".ad"))   // matches bad, dad, mad
		assert.True(t, wd.Search("b.."))   // matches bad
		assert.True(t, wd.Search("d.d"))   // matches dad
		assert.False(t, wd.Search(".ad.")) // no 4-letter words
	})

	t.Run("search with multiple wildcards", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("hello")

		assert.True(t, wd.Search("h...."))
		assert.True(t, wd.Search(".e..."))
		assert.True(t, wd.Search("h.l.o"))
		assert.True(t, wd.Search("....."))
		assert.False(t, wd.Search("h.l")) // wrong length
	})

	t.Run("all wildcards", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("cat")

		assert.True(t, wd.Search("..."))
		assert.False(t, wd.Search("...."))
	})

	t.Run("empty dictionary", func(t *testing.T) {
		wd := Constructor()

		assert.False(t, wd.Search("cat"))
		assert.False(t, wd.Search("..."))
	})

	t.Run("single character words", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("a")
		wd.AddWord("b")
		wd.AddWord("c")

		assert.True(t, wd.Search("a"))
		assert.True(t, wd.Search("."))
		assert.False(t, wd.Search("d"))
	})

	t.Run("duplicate words", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("test")
		wd.AddWord("test") // duplicate

		assert.True(t, wd.Search("test"))
	})

	t.Run("words with same prefix", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("cat")
		wd.AddWord("cats")
		wd.AddWord("car")

		assert.True(t, wd.Search("cat"))
		assert.True(t, wd.Search("cats"))
		assert.True(t, wd.Search("car"))
		assert.True(t, wd.Search("ca."))
		assert.True(t, wd.Search("ca.."))
		assert.False(t, wd.Search("ca"))
	})

	t.Run("wildcard matching complexity", func(t *testing.T) {
		wd := Constructor()

		wd.AddWord("aab")
		wd.AddWord("aac")

		assert.True(t, wd.Search("a.b"))  // matches aab
		assert.True(t, wd.Search("a.c"))  // matches aac
		assert.True(t, wd.Search("a.."))  // matches both
		assert.True(t, wd.Search(".ab"))  // matches aab
		assert.True(t, wd.Search(".ac"))  // matches aac
		assert.False(t, wd.Search("a."))  // wrong length (2 vs 3)
		assert.False(t, wd.Search("b..")) // no 3-letter words starting with b
	})
}
