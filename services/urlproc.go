package services

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rickalon/GoWebScraper/data"
)

func OrDone(ctx context.Context, ch <-chan *data.UrlObj) <-chan *data.UrlObj {
	stream := make(chan *data.UrlObj)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case stream <- val:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return stream
}

func UrlProc(ctx context.Context, url *data.MockURL, ch chan<- *data.UrlObj) {
	defer close(ch)
	var wg sync.WaitGroup
	ctxTo, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	for _, val := range url.Data {
		wg.Add(1)
		go proc(&wg, val, ch, ctxTo)
	}

	wg.Wait()
}

func proc(wg *sync.WaitGroup, url string, ch chan<- *data.UrlObj, ctx context.Context) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		resp, err := http.Get(url)
		if err != nil {
			ch <- &data.UrlObj{
				Url:    url,
				Result: "Url not found",
			}
			return
		}
		defer resp.Body.Close()
		tls := resp.TLS
		if tls == nil {
			ch <- &data.UrlObj{
				Url:    url,
				Result: "Tls not found",
			}
			return
		}
		for _, cert := range tls.PeerCertificates {
			res := &data.UrlObj{
				Url:     url,
				Result:  "Found",
				To:      cert.Subject.CommonName,
				From:    cert.Issuer.CommonName,
				Country: cert.Issuer.Country,
				ExpDate: cert.NotAfter.String(),
				EmiDate: cert.NotBefore.String(),
				Alg:     cert.PublicKeyAlgorithm.String(),
				DNS:     cert.DNSNames,
				IsCA:    cert.IsCA,
			}
			ch <- res
		}
		return
	}

}
