package images

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/seaweedfs/seaweedfs/weed/util"
)

func TestResizing(t *testing.T) {
	fname := "sample2.webp"

	dat, _ := os.ReadFile(fname)

	resized, _, _ := Resized(".webp", bytes.NewReader(dat), 100, 30, "")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resized)

	util.WriteFile("resized1.png", buf.Bytes(), 0644)

	os.Remove("resized1.png")

}

func TestResizedResetsReaderOnDecodeError(t *testing.T) {
	data := []byte("not an image")
	resized, _, _ := Resized(".jpg", bytes.NewReader(data), 10, 10, "")

	got, err := io.ReadAll(resized)
	if err != nil {
		t.Fatalf("ReadAll resized reader: %v", err)
	}
	if !bytes.Equal(got, data) {
		t.Fatalf("resized reader content = %q, want %q", got, data)
	}
}
