package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// ಠ_ಠ
func TestGetUrl(t *testing.T) {
	test_case := `origin   git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte`
	assert.Equal(t, expected, GetUrl(test_case), "works for ssh remote")
}

func TestGetUrl3(t *testing.T) {
	test_case := `origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)
origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (push)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte`
	assert.Equal(t, expected, GetUrl(test_case), "works for ssh remote")
}

func TestGetUrl2(t *testing.T) {
	test_case := `origin  https://github.com/syzer/repo.git (fetch)`
	expected := `https://github.com/syzer/repo`
	assert.Equal(t, GetUrl(test_case), expected, "works for https remote")
}
