package pwned

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var url = "https://api.pwnedpasswords.com"

func CheckPassword(passwd string) (int, error) {
	shaPasswd := fmt.Sprintf("%x", sha1.Sum([]byte(passwd)))
	resp, err := http.Get(fmt.Sprintf("%s/range/%s", url, shaPasswd[:5]))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, strings.ToUpper(shaPasswd[5:])) {
			nbResult, err := strconv.Atoi(strings.Split(line, ":")[1])
			if err != nil {
				log.Fatal(err)
			}
			return nbResult, nil
		}
	}

	return 0, nil
}

func SelectPassword(args []string, reader io.Reader) string {
	if len(args) == 0 {
		reader := bufio.NewReader(reader)
		fmt.Print("Enter candidate password: ")
		input, _ := reader.ReadString('\n')
		return strings.Replace(input, "\n", "", -1)
	} else {
		return args[0]
	}
}
