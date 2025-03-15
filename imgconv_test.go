package imgconv

import (
	"path/filepath"
	"testing"
)

func TestPattern(t *testing.T) {
	pattern := ""
	files, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range files {
		t.Logf("file path=%s", v)
	}

}
