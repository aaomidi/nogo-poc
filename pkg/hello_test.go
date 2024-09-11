package pkg

import "testing"

func TestHello(t *testing.T) {
	t.Run("someTest", func(t *testing.T) {
		if 1 != 1 {
			t.Error("1 != 1")
		}
	})
}
