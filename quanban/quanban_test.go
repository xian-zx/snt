package quanban

import "testing"

func TestGoQuanBan(t *testing.T) {
	if Q2B(QuanStr) != BanStr {
		t.Error("Q2B failed")
	}
	if B2Q(BanStr) != QuanStr {
		t.Error("B2Q failed")
	}
}
