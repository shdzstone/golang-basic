package main

import (
	"fmt"
	"os"
)

/*
1.广度优先搜索走迷宫
* 循环创建二维slice
* 使用slice来实现队列
* 用Fscanf读取文件
* 对Point的抽象
*/

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	//注：和其它语言的二维数组定义不同
	//go语言中的[][]int意思是[]int类型的一个slice
	//对二维数组的下一维声明需要使用for循环
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//i,j和坐标轴中的x,y不一样
//i是向下，j是向右
type point struct {
	i, j int
}

func (p point) String() string {
	return fmt.Sprintf("point:{i=%d,j=%d}", p.i, p.j)
}

//广度优先算法探索的四个方向：上、左、下、右
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//go中没有操作符重载：自己定义
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//返回maze在i,j处的值
func (p point) at(grid [][]int) (int, bool) {
	//第i行
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

//遍历二维数组
func walk(maze [][]int, start, end point) [][]int {
	fmt.Println(start)
	fmt.Println(end)

	//符合迷宫路径的点的坐标路径
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		//当前队列首元素
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			fmt.Println("break:", cur)
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			//maze at next is 0
			//and steps at next is 0
			//and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				fmt.Println("val=1:", next)
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				fmt.Println("value!=0", next)
				continue
			}

			if next == start {
				fmt.Println("==start", next)
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("lang/maze/maze.in")
	for _, row := range maze {
		for _, value := range row {
			fmt.Printf("%3d ", value)
		}
		fmt.Println()
	}

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d  ", val)
		}
		fmt.Println()
	}
}
