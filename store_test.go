package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransfromFunc(t *testing.T) {
	key := "momsbestpicture"
	pathkey := CASPathTransformFunc(key)
	expectedFileName := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathkey.PathName != expectedPathName {
		t.Errorf("have %s %s", pathkey.PathName, expectedPathName)
	}
	if pathkey.FileName != expectedFileName {
		t.Errorf("have %s %s", pathkey.FileName, expectedFileName)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)

		fmt.Println(string(b))
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("expected to NOT key %s", key)
		}
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
