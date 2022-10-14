package main

import "fmt"

func EmailInfo(email string) string {
	res := "Domain: dan TLD: com"
	res1 := ""
	res2 := ""
	bool1 := false
	bool2 := false
	
	for i := 0; i < len(email); i++{
		if bool1 == true && email[i] != '.'{
			res1 = res1 + string(email[i])
		} else if bool2 == true{
			res2 = res2 + string(email[i])
		}
		if email[i] == '@'{
			bool1 = true
		}
		if email[i] == '.' {
			bool2 = true
			bool1 = false
		}
	}

	res = "Domain: " + res1 + " dan TLD: " + res2

	return res // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
