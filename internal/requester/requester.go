package requester

import (
	"fmt"
	"net/http"
	"regexp"
	"io"
	"strconv"
	"strings"
	"sync"
)

func FindExchangeRates() (map[string]float64, error) {
	urls := []string{
		"usd-cad",
		"usd-eur",
		"usd-mxn",
		"cad-eur",
		"cad-mxn",
		"eur-mxn",
	}

	rates := make(map[string]float64)

	var mu sync.Mutex
	var wg sync.WaitGroup

	errChan := make(chan error, len(urls))


	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			fullURL := fmt.Sprintf("https://www.exchange-rates.org/converter/%s", url)
			response, err := http.Get(fullURL)
			if err != nil {
				errChan <- err
				return
			}

			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				errChan <- err
				return
			}

			pattern := regexp.MustCompile(`<span class="rate-to">(.*?)</span>`)
			match := pattern.FindStringSubmatch(string(body))
			if len(match) < 2 {
				errChan <- fmt.Errorf("Rate for %s not found", u)
				return
			}
			number := strings.Split(match[1], " ")[0]

			newRate, err := strconv.ParseFloat(number, 64)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			rates[url] = newRate

			currencies := strings.Split(url, "-")
			newKey := fmt.Sprintf("%s-%s", currencies[1], currencies[0])
			inverse := 1.0 / newRate
			rates[newKey] = inverse

			mu.Unlock()
		}(url)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return rates, nil
}
