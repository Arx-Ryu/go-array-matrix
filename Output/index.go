package main

import (
	"fmt"
)

func main() {
	//Input	
	matrix1 := [][]int{  
		{1,2,4},   
		{3,7,2},  
		{5,8,4},
	}
	matrix2 := [][]int{  
		{1,2,2},   
		{2,2,1},  
		{1,1,1},
	}
	matrix3 := [][]int{  
		{11,14,12},   
		{12,45,38},  
		{23,45,12},
	}

	fmt.Println(matrix1)
	fmt.Println(matrix2)
	fmt.Println(matrix3)

	//Process
	matrixFinal := matrixCompiler(matrix1, matrix2, matrix3)
	//Output
	fmt.Print(matrixFinal)
}


func matrixCompiler(m1, m2, m3 [][]int) [][]int {
	var mFinal [][]int
	var mTemp1, mTemp2, mTemp3 []int
	for x := 0; x < len(m1); x++ {
		for y:= 0; y < len(m1[x]); y++{
			mTemp1 = append(mTemp1, m1[x][y])	
			mTemp2 = append(mTemp2, m2[x][y])
			mTemp3 = append(mTemp3, m3[x][y])			
		}					
	} 	
	mFinal = append(mFinal, mTemp1)
	mFinal = append(mFinal, mTemp2)
	mFinal = append(mFinal, mTemp3)

	return mFinal
}
