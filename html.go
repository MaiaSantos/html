package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo obtem o título de uma pagina html
func Titulo(urls ...string) <-chan string { //... são parâmetros variáveis.
	c := make(chan string)
	for _, url := range urls { // _ ignora o índice.
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // transformando html em string.
		}(url)
	}
	return c
}
