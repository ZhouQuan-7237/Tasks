package main

import (
	"fmt"       // 格式化输入输出（1）读取输入：fmt.Scan(&input) → 从键盘获取用户输入（2）打印日志：fmt.Println("Hello") → 输出到屏幕（3）格式化输出：fmt.Printf("姓名：%s，年龄：%d", name, age) → 类似填空模板
	"math/rand" // 生成随机数。需要先“播种”随机种子（如用当前时间），否则每次生成的随机数相同。rand.Seed(time.Now().UnixNano()) 播种
	"time"      // 处理时间相关的操作，如获取当前时间time.Now()、计算时间间隔end.Sub(start)
	"strconv"   // 处理字符串和数值之间的类型转换，把字符串转成数字，或把数字转成字符串
	"os"        // 处理操作系统的交互，如读取标准输入（os.Stdin）、文件操作
	"strings"   // 处理字符串，如去除空格strings.TrimSpace(" hello ") → "hello"、分割字符串strings.Split("a,b,c", ",") → ["a", "b", "c"]
	"bufio"     // 缓冲输入，更高效读取用户输入
)

func main() {

	const (														 // 通过设置最小值和最大值限定随机数的范围
		min = 1     											 // 最小值为1
		max = 1000  											 // 最大值为1000
	)

	rand.Seed(time.Now().UnixNano())							 // 初始化随机数生成器种子（随机数的起点），确保每次运行生成的随机数不同	rand.Seed设置随机数生成器的种子, time.Now()返回当前的时间 ,	UnixNano() 表示当前时间的精确时刻 , time.Now().UnixNano()获取当前时间的纳秒级时间戳 , rand.Seed(time.Now().UnixNano())将当前时间的纳秒级时间戳作为种子传递给随机数生成器,由于每次运行程序时，当前时间都会有所不同，因此这样可以确保每次运行程序时生成的随机数序列都是不同的					
	target := rand.Intn(max-min+1) + min						 // 定义变量target为目标数字，生成一个 1~1000 之间的随机整数赋值给target, rand.Intn(n) 生成一个0到n-1之间的随机整数, rand.Intn(max-min+1) 生成一个0到999之间的随机数，然后加上 min（即1），确保生成的数字在1到1000之间	
	fmt.Println("目标数字在1-1000之间")							  // 提示用户目标数字的范围					
															
	var ciShu int											 	 // 变量ciShu记录猜测次数															
	startTime := time.Now()										 // 变量startTime记录游戏开始的时间
	scanner := bufio.NewScanner(os.Stdin)						 // 创建 bufio.Scanner 对象，用于读取用户输入。bufio.NewScanner创建一个扫描器（Scanner）, os.Stdin表示标准输入（Standard Input）,通常是指通过键盘输入的内容 ,bufio.NewScanner(os.Stdin)用标准输入作为扫描器的输入源,可以逐行读取用户输入的内容, scanner := bufio.NewScanner(os.Stdin)创建一个逐行读取标准输入的扫描器,可以读取用户通过键盘输入的多行文本
	
	for {
		fmt.Print("请输入你猜的数字: ") 						  // 提示用户输入数字
		
		if !scanner.Scan() {									// 检查是否成功读取了一行用户输入。scanner.Scan() 是 bufio.Scanner 的一个方法，它从输入源（如 os.Stdin）中读取一行内容。如果 scanner.Scan() 成功读取了一行输入（即用户输入了有效内容并按下了回车键），它返回 true;如果读取失败，scanner.Scan() 会返回 false,常见的失败原因包括:输入流结束（比如，用户直接关闭终端窗口）
			fmt.Println("输入错误，请重新输入！")
			continue											// 继续下一次循环，让用户重新输入
		}

		input := strings.TrimSpace(scanner.Text())				 // 变量input记录去除首尾空格的输入内容。scanner.Text() 获取用户输入的原始字符串,strings.TrimSpace() 去除输入字符串的首尾空格，避免因多余空格导致的错误
		number, err := strconv.Atoi(input)   					 // strconv.Atoi() 将输入的字符串input转换为整数,并将结果存储在变量number中,如果转换过程中出现错误,将错误信息存储在变量 err 中。number 是第一个返回值，存储转换后的整数值,err 是第二个返回值，判断转换过程中是否有错误发生
		
		if err != nil { 										  // 检查输入内容是否可转换为整数。如果转换失败，说明输入不是有效整数：（1）输入内容包含非数字字符（如 abc、12.3）（2）输入内容为空字符串（直接按回车）。如果 err 为 nil,表示操作成功，没有发生错误;如果 err 不为 nil，表示操作失败，发生了错误。
			fmt.Println("输入无效,请重新输入！")				 
			continue 											 // 继续下一次循环，让用户重新输入
		}
		
		if number < min || number > max {						 // 判断输入是否在 1~1000 之间
			fmt.Println("输入超范围,请重新输入！")
			continue 											 // 继续下一次循环，让用户重新输入
		}									
		
		ciShu++												     // 猜测次数加一
		
		if number > target {									
			fmt.Println("猜大了，再试试！") 					  // 提示用户猜大了
		} else if number < target {
			fmt.Println("猜小了，再试试！") 					  // 提示用户猜小了
		} else {
			duration := time.Since(startTime)                    //计算从 startTime 到现在的时间间隔并赋值给duration。time.Since() 计算从指定的时间 startTime 到当前时间的持续时间
			fmt.Printf("恭喜！你用了%d次猜对了答案%d,共用时%.2f秒。\n", ciShu, target, duration.Seconds())// 输出最终结果，包括猜测次数和游戏用时,duration.Seconds() 将 time.Duration 类型的时间差转换为浮点数
			break												// 退出循环，结束游戏
		}
	}
}