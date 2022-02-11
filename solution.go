package solution

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.Sprint("Hello 🗺️!")
}

func main() {
	fmt.Println(GetMessage())
}
