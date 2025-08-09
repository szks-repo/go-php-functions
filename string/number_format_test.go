package string

import (
	"reflect"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		if !reflect.DeepEqual(NumberFormat[int](7777), "7,777") {
			t.Errorf("failed")
		}
	})

	t.Run("int64", func(t *testing.T) {
		if !reflect.DeepEqual(NumberFormat[int64](-7777), "-7,777") {
			t.Errorf("failed")
		}
	})

	t.Run("uint", func(t *testing.T) {
		if !reflect.DeepEqual(NumberFormat[uint](7777), "7,777") {
			t.Errorf("failed")
		}
	})

	t.Run("uint64", func(t *testing.T) {
		if !reflect.DeepEqual(NumberFormat[uint64](7777), "7,777") {
			t.Errorf("failed")
		}
	})

	t.Run("float64", func(t *testing.T) {
		if !reflect.DeepEqual(NumberFormat[float64](7777.7777), "7,777.778") {
			t.Errorf("failed: got=%s", NumberFormat[float64](7777.7777))
		}
	})

}
