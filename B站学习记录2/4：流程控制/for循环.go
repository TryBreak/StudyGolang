package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("for 循环")

	var sum int = 0
	for i := 1; i <= 5; i++ {
		fmt.Println("for 循环i", i)
		sum += i
		fmt.Println("for 循环sum", sum)
	}
	fmt.Println("-----------")

	// for i
	var str string = "hello中golang"

	var strbytes []byte = []byte(str)

	for i := 0; i < len(strbytes); i++ {
		fmt.Println(i, string(strbytes[i]))
	}

	fmt.Println("-----------")

	// for range
	var str2 string = "hello中golang"

	for i, v := range str2 {
		fmt.Println(i, string(v))
	}

	fmt.Println("-----------break--------------")

	// ====  break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			// 停止终结循环
			break
		}
	}
	fmt.Println("-----------双for循环---标签使用----------")

Label1:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 2 && j == 2 {
				break Label1
			}
			fmt.Printf("i:%v,j:%v \n", i, j)
		}
	}

	fmt.Println("------输出1-100中被6整除的数--------")

	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("------输出1-100中被6整除的数---continue-----")

Label3:
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue Label3 // 跳出本次循环
		}
		fmt.Println(i)
	}

	fmt.Println("-----99乘法表--------")

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				break
			}
			if j == i {
				fmt.Printf("%vx%v=%v \n", j, i, i*j)
			} else {
				fmt.Printf("%vx%v=%v ", j, i, i*j)
			}
		}
	}

	fmt.Println("-----99乘法表--一次性输出------")
	var NieTable string

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				break
			}
			var a = strconv.FormatInt(int64(i), 10)
			var b = strconv.FormatInt(int64(j), 10)
			var c = strconv.FormatInt(int64(j*i), 10)
			if j >= 2 && j*i < 10 {
				c += " "
			}
			NieTable += b + "x" + a + "=" + c + " "

			if j == i {
				NieTable += "\n"
			}

		}
	}

	fmt.Println(NieTable)

}
