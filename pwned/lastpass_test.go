package pwned

import (
	"errors"
	"testing"
	"io"
)

func TestListAllPasswords(t *testing.T) {

	ls := func() (map[string]string, error) {
/*		passwords := `(none)
    192.168.1.13 [id: 1]
    Betaseries [id: 2]
    Coursera [id: 3]
Business
    wordpress.com [id: 4]
Outils de productivité
    dropbox.com [id: 5]`*/
		passwords := map[string]string {
			"1": "192.168.1.13",
			"2": "Betaseries",
			"3": "Coursera",
			"4": "wordpress.com",
			"5": "dropbox.com",
		}

		return passwords, nil
	}

	show := func(id string, exec executor) (string, error) {
		switch id {
		case "1":
			return "password-ip", nil
		case "2":
			return "password-betaseries", nil
		case "3":
			return "password-coursera", nil
		case "4":
			return "password-wordpress", nil
		case "5":
			return "password-dropbox", nil
		default:
			return "", errors.New("you missed a test case")
		}
	}

	result := listAllPasswords(ls, show)

	if len(result) != 5 {
		t.Error("Expecting to list 5 passwords, got; ", len(result))
	}

	ip := result["1"]
	if ip.name != "192.168.1.13" && ip.value != "password-ip" {
		t.Errorf("Expecting ID [1] to have name [192.168.1.13] and password [password-1], got: [%s] and [%s]", ip.name, ip.value)
	}
	betaseries := result["2"]
	if betaseries.name != "Betaseries" && betaseries.value != "password-betaseries" {
		t.Errorf("Expecting ID [2] to have name [Betaseries] and password [password-betaseries], got: [%s] and [%s]", betaseries.name, betaseries.value)
	}
	coursera := result["3"]
	if coursera.name != "Coursera" && coursera.value != "password-coursera" {
		t.Errorf("Expecting ID [3] to have name [Coursera] and password [password-coursera], got: [%s] and [%s]", coursera.name, coursera.value)
	}
	wordpress := result["4"]
	if wordpress.name != "wordpress.com" && wordpress.value != "password-wordpress" {
		t.Errorf("Expecting ID [4] to have name [wordpress.com] and password [password-wordpress], got: [%s] and [%s]", wordpress.name, wordpress.value)
	}
	dropbox := result["5"]
	if dropbox.name != "dropbox.com" && dropbox.value != "password-dropbox" {
		t.Errorf("Expecting ID [5] to have name [dropbox.com] and password [password-dropbox], got: [%s] and [%s]", dropbox.name, dropbox.value)
	}
}

func TestParseShowCommand(t *testing.T) {
	output := func(output io.Writer, command ...string) error {
		output.Write([]byte(`[
  {
    "id": "3",
    "name": "sso.garmin.com",
    "fullname": "Santés/sso.garmin.com",
    "username": "toto@gmail.com",
    "password": "s3cr3t",
    "group": "Santés",
    "url": "https://sso.garmin.com/sso/login",
    "note": ""
  }
]`))
		return nil
	}

	result, err := show("3", output)

	if result != "s3cr3t" {
		t.Error("Expecting password to be s3cr3t, got: ", result)
	}
	if err != nil {
		t.Error("Expecting a password, not an error...", err)
	}
}

