package main

import (
    "crypto/tls"
    "net/http"
    "sync"
    "sync/atomic"
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_LimiterByIp(t *testing.T) {
    var nNumeroDeChamadas int = 100

    wg := sync.WaitGroup{}
    wg.Add(nNumeroDeChamadas)

    var Http200 atomic.Int64
    var Http429 atomic.Int64
    var HttpOutros atomic.Int64

    for nKey := 1; nKey <= nNumeroDeChamadas; nKey++ {
        //wg.Add(1)
        go func() {
            defer wg.Done()
            http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

            //r, _ := http.NewRequest("GET", "http://localhost:8080", nil)
            //r.Header.Add("CF-Connecting-IP", "172.30.2.5")

            //res, Err := http.DefaultClient.Do(r)
            res, Err := http.Get("http://127.0.0.1:8080")

            if Err == nil {
                switch res.StatusCode {
                case http.StatusOK:
                    Http200.Add(1)
                case http.StatusTooManyRequests:
                    Http429.Add(1)
                default:
                    HttpOutros.Add(1)
                }
            } else {
                println(Err.Error())
            }
        }()
    }
    wg.Wait()

    assert.Less(t, Http200.Load(), int64(200))
    assert.Greater(t, Http429.Load(), int64(1))
    assert.Equal(t, int64(0), HttpOutros.Load())
}

func Test_LimiterByApiKy(t *testing.T) {
    var nNumeroDeChamadas int = 200

    wg := sync.WaitGroup{}
    wg.Add(nNumeroDeChamadas)

    var Http200 atomic.Int64
    var Http429 atomic.Int64
    var HttpOutros atomic.Int64

    for nKey := 1; nKey <= nNumeroDeChamadas; nKey++ {
        go func() {
            defer wg.Done()
            http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

            r, _ := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
            r.Header.Add("API_KEY", "TOKENA")

            res, Err := http.DefaultClient.Do(r)

            if Err == nil {
                switch res.StatusCode {
                case http.StatusOK:
                    Http200.Add(1)
                case http.StatusTooManyRequests:
                    Http429.Add(1)
                default:
                    HttpOutros.Add(1)
                }
            } else {
                println(Err.Error())
            }
        }()
    }

    wg.Wait()

    assert.Less(t, Http200.Load(), int64(200))
    assert.Greater(t, Http429.Load(), int64(1))
    assert.Equal(t, HttpOutros.Load(), int64(0))
}
