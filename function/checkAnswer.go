package function

import (
	"strings"

	"github.com/JustForC/CobaGolang/model"
)

func CheckAnswer(question model.Question, user *model.User, answer string) {
	if strings.Compare(question.Answer, answer) == 0 {
		user.Point = user.Point + 1
	}
}
