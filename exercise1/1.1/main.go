//修改 echo 程序输出os.Args[0],即命令的名字

/**
echo1:
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
**/
/**
echo2:
package main
import(
	"fmt"
	"os"
)
func main(){
	s,sep:="",""
	for_.arg:=range os.Args[1:]{
		s+=sep+arg
		sep=" "
	}
	fmt.Println(s)
}
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""

	for _, arg := range os.Args[:1] {
		s += arg
	}

	fmt.Println(s)
}
