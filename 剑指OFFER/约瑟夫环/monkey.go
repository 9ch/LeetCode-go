package main

import "fmt"

/**
一群猴子排成一圈，按1，2，...，n依次编号。

然后从第1只开始数，数到第m只,把它踢出圈，

从它后面再开始数，再数到第m只，在把它踢出去...，

如此不停的进行下去，直到最后只剩下一只猴子为止，那只猴子就叫做大王。
*/

func main() {
	m := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	m = append(m[:1], m[2:]...)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(monkey(m, 3))
}

func monkey(monkeys []string, m int) []string {
	i := 1
	for len(monkeys) > 1 {
		if i%m == 0 {
			monkeys = append(monkeys[:i-1], monkeys[i:]...)
			i = 1
			continue
		} else {
			monkeys = append(monkeys, monkeys[i-1])
			monkeys = append(monkeys[:i-1], monkeys[i:]...)
			i++
		}
	}
	return monkeys
}
