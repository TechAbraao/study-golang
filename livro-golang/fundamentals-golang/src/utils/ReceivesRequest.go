package utils

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
)

func Request(profileName string) {
	var profile string = profileName
	url := fmt.Sprintf("https://api.github.com/users/%s", profile)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
