package myshawdowsockets
import (
	"sort"
	"testing"
)

func (password *password) Len() int {
	return passwordLength
}

func (password *password) Less(i, j int) bool {
	return password[i] < password[j]
}

func (password *password) Swap(i, j int) {
	password[i], password[j] = password[j], password[i]
}

func TestRandPassword(t *testing.T) {
	password := RandPassword()
	t.Log(password)
	bsPassword, err := ParsePassword(password)
	if err != nil {
		t.Error(err)
	}
	sort.Sort(bsPassword)
	for i := 0; i < passwordLength; i++ {
		if bsPassword[i] != byte(i) {
			t.Error("password form failed")
		}
	}
}