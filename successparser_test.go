package goparsec

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	parser := Success()
	result, err := parser.Parse("aaa")
	if err != nil {
		t.Errorf("Success should always success.")
		return
	}
	if result.Get() != nil {
		t.Errorf("result should be nil.")
		return
	}
}
