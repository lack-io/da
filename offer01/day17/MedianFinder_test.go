package day17

import "testing"

func TestMedianFinder(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	mid := obj.FindMedian()
	t.Log(mid)

	obj.AddNum(2)
	obj.AddNum(3)
	mid = obj.FindMedian()
	t.Log(mid)
}
