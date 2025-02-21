package test_common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeepCopy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}

		src := TestStruct{
			Name: "test",
			Age:  10,
		}

		dst, err := DeepCopy(src)
		require.NoError(t, err)
		require.Equal(t, src, *dst)

		// 実体は新しくなることの確認
		src.Name = "test2"
		require.Equal(t, dst.Name, "test")
		src.Name = "test"
		dst.Name = "test2"
		require.Equal(t, src.Name, "test")
	})

	t.Run("success (primitive)", func(t *testing.T) {
		src := "test"
		dst, err := DeepCopy(src)
		require.NoError(t, err)
		require.Equal(t, src, *dst)

		srcPtr := &src
		*srcPtr = "test2"
		require.Equal(t, *dst, "test")
	})
}
