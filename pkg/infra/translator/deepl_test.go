package translator

import (
	"fmt"
	"testing"
)

func TestDeepLClient_Do(t *testing.T) {
	s, err := DefaultDeepLClient.Do("he")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}