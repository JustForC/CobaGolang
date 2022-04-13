package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JustForC/CobaGolang/function"
	"github.com/JustForC/CobaGolang/model"
)

var questions []model.Question

func init() {
	questions = []model.Question{
		model.Question{"Siapakah Nama Presiden Indonesia?", "Joko Widodo"},
		model.Question{"Siapakah Nama Presiden Amerika Serikat?", "Bukan Joko Widodo"},
	}
}

func main() {

	var user model.User
	scanner := bufio.NewScanner(os.Stdin)
	var answer string

	// Question Initiate
	fmt.Println("Masukan Nama Depan Anda")
	fmt.Scanln(&user.FirstName)
	fmt.Println("Masukan Nama Belakang Anda")
	fmt.Scanln(&user.LastName)
	fmt.Println("Mari Kita Mulai Quiznya!")

	for i := 0; i < len(questions); i++ {
		fmt.Println(questions[i].Question)
		if scanner.Scan() {
			answer = scanner.Text()
		}
		function.CheckAnswer(questions[i], &user, answer)
	}

	fmt.Println(user.FirstName, user.LastName, "Point Kamu Adalah", user.Point)
}
