//修改 echo 程序,输出参数的索引和值，每行一个

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

	for index, arg := range os.Args[1:] {

		fmt.Printf("%d\t%s\n", index+1, arg)
	}
}
