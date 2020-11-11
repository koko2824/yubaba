package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

func strStdin() (strInput string) {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	name := stdin.Text()
	return name
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func rename(name string) (newName string) {
	rand.Seed(time.Now().UnixNano())
	i := utf8.RuneCountInString(name)
	newName = string(getRuneAt(name, rand.Intn(i)))
	return newName
}

func main() {
	fmt.Println("契約書だよ。そこに名前を書きな。")
	name := strStdin()
	fmt.Printf("フン、「%s」というのかい。贅沢な名前だねぇ。\n", name)
	newName := rename(name)
	fmt.Println("今後お前の名前は「"+newName+"」だよ。いいかい、「"+newName+"」だよ。分かったら返事をするんだ！「"+newName+"」!!")
}
