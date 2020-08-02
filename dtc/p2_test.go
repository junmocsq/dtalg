package dtc

import "testing"

func TestMaxSubsequenceSum(t *testing.T) {
	arr := []int{-2, 11, -4, 13, -5, -2}
	want := 20
	_, _, actual := maxSubsequenceSumV1(arr)
	if want != actual {
		t.Fatalf("MaxSubsequenceSumV1 N^3 failed, want:%d actual:%d", want, actual)
	}

	_, _, actual = maxSubsequenceSumV2(arr)
	if want != actual {
		t.Fatalf("MaxSubsequenceSumV2 N^2 failed, want:%d actual:%d", want, actual)
	}

	actual = maxSubsequenceSumV3(arr, 0, len(arr))
	if want != actual {
		t.Fatalf("MaxSubsequenceSumV3 N*LogN failed, want:%d actual:%d", want, actual)
	}

	actual = maxSubsequenceSumV4(arr)
	if want != actual {
		t.Fatalf("MaxSubsequenceSumV4 N failed, want:%d actual:%d", want, actual)
	}
}
