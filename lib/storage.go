package lib

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"sync"
)

// Storage .
type Storage interface {
	Load(token interface{}) error
	Store(token interface{}) error
}

// FileStore .
type FileStore struct {
	sync.RWMutex
	FilePath string
}

// Load .
func (m *FileStore) Load(token interface{}) error {
	file, err := m.read()
	if err != nil {
		return err
	}
	return json.Unmarshal(file, token)
}

func (m *FileStore) read() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()
	return ioutil.ReadFile(m.FilePath)
}

// Store .
func (m *FileStore) Store(token interface{}) error {
	marshal, err := json.Marshal(token)
	if err != nil {
		return errors.WithStack(err)
	}
	return m.write(marshal)
}

func (m *FileStore) write(marshal []byte) error {
	m.Lock()
	defer m.Unlock()
	return ioutil.WriteFile(m.FilePath, marshal, os.ModePerm)
}
