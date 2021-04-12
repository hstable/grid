package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/config"
	"net/http"
	"sync"
	"testing"
)

func BenchmarkLogin(t *testing.B) {
	var wg sync.WaitGroup
	b, _ := json.Marshal(gin.H{"Username": "test", "Password": "test"})
	for i := 0; i < t.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := http.Post("http://127.0.0.1:"+config.Get().ListeningPort+"/api/user/login", "json", bytes.NewReader(b))
			if err != nil || res.StatusCode >= 300 {
				t.Fail()
				return
			}
			res.Body.Close()
		}()
	}
	wg.Wait()
}
