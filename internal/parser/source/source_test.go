package source

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringSource_ReadRunes(t *testing.T) {
	content := "aaabbbccc"
	source := NewStringSource(content)
	compareStringSourceWithString(t, source, content)
}

func Test_FileSource_ReadRunes(t *testing.T) {
	content := "very long \n\t\n\ttest striiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiing\t\n"
	filename := createTmpFile(t, "tmp*.proto", content)
	defer os.Remove(filename)

	source, err := NewFileSource(filename)
	if err != nil {
		t.Fatalf("failed to open file source: %v", err)
	}
	defer source.Close()

	compareFileSourceWithString(t, source, content)
}

/*
 * #################
 * utils for testing
 * #################
 */

func createTmpFile(t *testing.T, filename string, data string) string {
	t.Helper()
	file, err := os.CreateTemp("", filename)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.WriteString(data); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
	return file.Name()
}

func compareStringSourceWithString(t *testing.T, src *stringSource, data string) {
	t.Helper()
	var actual []rune
	for {
		r, err := src.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		actual = append(actual, r)
	}

	assert.Equal(
		t,
		[]rune(data),
		actual,
	)

	r, err := src.Next()
	assert.Zero(t, r)
	assert.True(t, errors.Is(err, io.EOF))
}

func compareFileSourceWithString(t *testing.T, src *fileSource, data string) {
	t.Helper()
	var actual []rune
	for {
		r, err := src.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		actual = append(actual, r)
	}

	assert.Equal(
		t,
		[]rune(data),
		actual,
	)

	r, err := src.Next()
	assert.Zero(t, r)
	assert.True(t, errors.Is(err, io.EOF))
}
