package go_try

import (
	"go/token"
	"testing"
)

var parseStr = "orders > 10000 && driving_years > 5 "

func TestParseStr(t *testing.T) {
	PrintStrExpr(parseStr)
}

func TestOperateCompare(t *testing.T) {
	if flag, err := OperateCompare(1, 2, token.NEQ); err != nil || !flag {
		t.Error()
	}

}
