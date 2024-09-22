package services

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rickalon/GoWebScraper/data"
)

func OrDone(ctx context.Context, ch <-chan *data.URL) <-chan *data.URL {
	stream := make(chan *data.URL)
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

func UrlProc(ctx context.Context, url *data.MockURL, ch chan<- *data.URL) {
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

func proc(wg *sync.WaitGroup, url string, ch chan<- *data.URL, ctx context.Context) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		resp, err := http.Get(url)
		if err != nil {
			ch <- &data.URL{UrlName: url}
			return
		}
		defer resp.Body.Close()
		tls := resp.TLS
		if tls == nil {
			ch <- &data.URL{UrlName: url}
			return
		}
		resContainer := &data.URL{UrlName: url}
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
			resContainer.Data = append(resContainer.Data, res)
		}
		ch <- resContainer
		return
	}

}
