package lesson1

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func init() {
	fmt.Println("Map is ", m)
	info = fmt.Sprintf("Info goos: %v, goarch : %v", runtime.GOOS, runtime.GOARCH)

}

var m = map[string]string{"name": "roger zou", "age": "41"}
var info string

func GetInfo() string {
	return info
}
func PrintProcess() {
	fmt.Printf("current process info is %v", os.Getppid())
	var str = `
				hello
				%v
				ok
				`
	fmt.Printf(str, "roger")
}

func Closure(i int) {
	defer func(last int) {
		last = last * last
		fmt.Printf("\nadd last %d", last)
	}(i)

	i += i

	fmt.Printf("result is %d", i)
}
func Sum(arr []int) (result int) {
	for size, i := len(arr), 0; i < size; i++ {
		result += arr[i]
	}

	//
	for i, v := range arr {
		result += v
		fmt.Printf("\nindex %d, value %d ", i, v)
	}

	return
}

func SwitchFunc() {
	var v interface{}

	switch i := v.(type) {
	default:
		fmt.Println("case default", i)
		break
	case string:
		fmt.Println("case", i)
		break
	}
}

func ReadStr() {
	inputReader := bufio.NewReader(os.Stdin)
	inputs, error := inputReader.ReadString('\n')
	if error != nil {
		fmt.Println("read error")
		return
	}

	inputs = inputs[:len(inputs)]
	fmt.Printf("\nInput string is '%v'", inputs)
}

type myTalk string

func (talk *myTalk) Hello(userName string) (result string) {
	result = "hello " + userName
	return
}
func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	saying = heard + " saying ..."
	end = true
	//确保被执行
	defer func() {
		//如果是要处理的恐慌，则处理，否则抛出。
		//recover == catch.
		if e := recover(); e != nil {
			if se, ok := e.(DivideError); ok {
				err = se.Err
			} else {
				panic(e)
			}
		}
	}()
	return
}

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

////////////////////error implement.....
type DivideError struct {
	Dividee int
	Divider int
	Err     error
}

func (de *DivideError) Error() string {
	return fmt.Sprintf(" Cannot proceed, the divider is zero.\ndividee: %d\ndivider: 0", de.Dividee)
}
