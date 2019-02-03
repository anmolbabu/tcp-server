package utils

import (
	"fmt"
	"strings"
	"sync"
)

type JsonStore struct {
	data []string
	sync.RWMutex
}

type SizeStringSlice struct {
	Data []string
}

func NewSizeStringSlice() *SizeStringSlice {
	s := SizeStringSlice{}
	return &s
}

func (js *JsonStore) Write(p string) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	js.Lock()
	js.data = append(js.data, p)
	js.Unlock()
	return len(p), nil
}

func (js *JsonStore) WriteMany(p []string) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	js.Lock()
	js.data = append(js.data, p...)
	js.Unlock()
	return len(p), nil
}

func (js *JsonStore) Read(p *SizeStringSlice, isRemoveRead bool) (n int, err error) {
	js.RLock()
	defer js.RUnlock()

	// Do we have data to send?
	if len(js.data) == 0 {
		return 0, nil
	}

	p.Data = make([]string, len(js.data))
	n = copy(p.Data, js.data)
	if isRemoveRead {
		js.data = js.data[n:]
	}
	return n, nil
}

func NewJsonStore() *JsonStore {
	return &JsonStore{}
}

func (js *JsonStore) Search(p *SizeStringSlice, searchStr string) {
	js.RLock()
	defer js.RUnlock()
	searchStr = strings.TrimSpace(searchStr)

	// Do we have data to send?
	if len(js.data) == 0 {
		return
	}

	for _, val := range js.data {
		if strings.Contains(val, searchStr) {
			fmt.Println("Match found")
			p.Data = append(p.Data, val)
		}
	}

	return
}
