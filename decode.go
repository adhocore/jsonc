package jsonc

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type cachedDecoder struct {
	jsonc *Jsonc
	ext   string
}

func CachedDecoder(ext ...string) *cachedDecoder {
	ext = append(ext, ".cached.json")
	return &cachedDecoder{New(), ext[0]}
}

func (fd *cachedDecoder) Decode(file string, v interface{}) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}

	cache := strings.TrimSuffix(file, filepath.Ext(file)) + fd.ext
	cstat, err := os.Stat(cache)
	exist := !os.IsNotExist(err)
	if err != nil && exist {
		return err
	}

	// Update if not exist, or source file modified
	update := !exist || stat.ModTime().UnixMilli() != cstat.ModTime().UnixMilli()
	if !update {
		jsonb, _ := ioutil.ReadFile(cache)
		return json.Unmarshal(jsonb, v)
	}

	jsonb, _ := ioutil.ReadFile(file)
	cfile, err := os.Create(cache)
	if err != nil {
		return err
	}

	jsonb = fd.jsonc.Strip(jsonb)
	cfile.Write(jsonb)
	os.Chtimes(cache, stat.ModTime(), stat.ModTime())
	return json.Unmarshal(jsonb, v)
}
