package lib

import (
	"encoding/json"
	"log"
	"sync"
	"testing"
	"time"
)

var str string
var fileStore *FileStore

func init() {
	fileStore = &FileStore{FilePath: "./access_token.json"}
	str = `{"error":"","error_description":"ok","expires_in":123,"access_token":"24.789f34b30c2def51ce7095fd86b62295.2592000.1602746610.282335-15708295","expires_time_stamp":1602746610}`
}

func TestFileStore_Load(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1 := time.NewTicker(time.Millisecond)
			t2 := time.NewTicker(10 * time.Second)
			var end bool
			for {
				select {
				case <-t1.C:
					token := &AccessToken{}
					if err := fileStore.Load(token); err != nil {
						t.Error(err)
					}
					log.Println(token.ExpiresIn, token.ErrorDescription)

				case <-t2.C:
					log.Println("ending ...")
					end = true
				}
				if end {
					break
				}
			}
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1 := time.NewTicker(time.Millisecond)
			t2 := time.NewTicker(10 * time.Second)
			var end bool
			for {
				select {
				case <-t1.C:
					token := &AccessToken{}
					json.Unmarshal([]byte(str), token)
					token.ExpiresIn = time.Now().Unix()

					if err := fileStore.Store(token); err != nil {
						t.Error(err)
					}

				case <-t2.C:
					log.Println("ending ...")
					end = true
				}
				if end {
					break
				}
			}
		}()
	}

	wg.Wait()
}

func BenchmarkFileStore_Load(b *testing.B) {
	fileStore := &FileStore{FilePath: "./access_token.json"}
	token := &AccessToken{}

	for i := 0; i < b.N; i++ {
		err := fileStore.Load(token)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFileStore_Store(b *testing.B) {
	fileStore := &FileStore{FilePath: "./access_token.json"}
	token := &AccessToken{}
	json.Unmarshal([]byte(str), token)
	for i := 0; i < b.N; i++ {
		err := fileStore.Store(token)
		if err != nil {
			b.Error(err)
		}
	}
}
