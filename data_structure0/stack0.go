package main

import "fmt"
//栈的结构体
type stack struct {
	str []string
}
//判断栈内是否为空
func (sta *stack)Empty() bool{
	return len(sta.str) == 0
}
//入栈
func (sta *stack)Push(s1 string){
	sta.str = append(sta.str, s1)
}
//出栈
func (sta *stack)Pop(){
	sta.str = sta.str[:len(sta.str) - 1]
}
//获取栈首元素
func (sta *stack)Top() string{
	return sta.str[len(sta.str) - 1]
}
//主函数
func main(){
	var sta stack		//声明结构体变量
	s := "{[]}"
	for _, v := range s{		//遍历循环
		s1 := string(v)			//将取出的元素转为字符串
		if sta.Empty(){			//判断栈内是否为空，若为空，则直接压入元素
			sta.Push(s1)
		}else{					//若不内空
			if (sta.Top() == "(" && s1 == ")") || (sta.Top() == "[" && s1 == "]") || (sta.Top() == "{" && s1 == "}"){
				sta.Pop()		//如果栈首元素与当前字符对称，则弹出栈首
			}else{
				sta.Push(s1)	//若不对称，则直接入栈
			}
		}
	}
	fmt.Println(sta.Empty())	//通过判断栈是否为空来判断是否符合要求
	return
}
