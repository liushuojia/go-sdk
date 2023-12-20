package dw

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"os"
	"testing"
)

// 目录下 go test
func TestHelloWorld(t *testing.T) {
	t.Log("draw")

	srcFile, err := os.Create("./new_os.png")
	if err != nil {
		fmt.Println("open err", err)
		return
	}
	defer srcFile.Close()
	wr := bufio.NewWriter(srcFile)

	New(1024, 1024).
		Background("#ffffff").
		Image("./go.jpg", 500, 500, 300, 300, true).
		Text("ABCCC100ABCCC100ABCCC100ABCCC100ABCCC100", 100, 100, 200, gg.AlignCenter).
		Writer(wr)
	//Save("./new.png")
}
