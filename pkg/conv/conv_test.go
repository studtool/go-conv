package conv

import (
	"testing"

	"github.com/franela/goblin"

	"strings"
)

func TestBytesToString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("BytesToString", func() {
		g.It("converts []byte to string", func() {
			g.Assert(BytesToString([]byte{'a', 'b', 'c'})).Equal("abc")
		})
	})
}

func TestStringToBytes(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("StringToBytes", func() {
		g.It("converts string to []byte", func() {
			g.Assert(StringToBytes("abc")).Equal([]byte{'a', 'b', 'c'})
		})
	})
}

func BenchmarkBytesToString(b *testing.B) {
	const n = 64
	v := makeBytes(n)

	for i := 0; i < b.N; i++ {
		_ = BytesToString(v)
	}
}

func BenchmarkBytesToStringAlloc(b *testing.B) {
	const n = 64
	v := makeBytes(n)

	for i := 0; i < b.N; i++ {
		_ = string(v)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	const n = 64
	v := makeString(n)

	for i := 0; i < b.N; i++ {
		_ = StringToBytes(v)
	}
}

func BenchmarkStringToBytesAlloc(b *testing.B) {
	const n = 64
	v := makeString(n)

	for i := 0; i < b.N; i++ {
		_ = []byte(v)
	}
}

func makeBytes(n int) []byte {
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = 'a'
	}
	return v
}

func makeString(n int) string {
	v := strings.Builder{}
	v.Grow(n)
	for i := 0; i < n; i++ {
		v.WriteByte('a')
	}
	return v.String()
}
