package main

import (
	"testing"
	"fmt"
)

// ಠ_ಠ
func TestGetUrl(t *testing.T) {
	test_case := `origin   git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte`
	fmt.Println(`output:`, GetUrl(test_case))
	fmt.Print(expected)
	t.SkipNow()
}

func TestGetUrl2(t *testing.T) {
	test_case := `origin  https://github.com/syzer/repo.git (fetch)`
	expected := `https://github.com/syzer/repo`
	if (expected != GetUrl(test_case)) {
		t.Fail()
	}
}
