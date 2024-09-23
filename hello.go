package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	student_grades := make(map [string][]int)
	fmt.Println("### Read as reader ###")
	f, err := os.Open("students.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
 
	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
		res:= strings.Split(sc.Text(), " ")
		num, _ := strconv.Atoi(res[1])
		student_grades[res[0]] = append(student_grades[res[0]], num)

		
	}

	keys := make([]string, len(student_grades))
    i := 0
    for k := range student_grades {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
	
	for _, key := range keys{
		fmt.Println(key)
		fmt.Println(student_grades[key])
		fmt.Print("Scores: ")
		var sum_grades =0
		for  i :=0; i<len(student_grades[key]);i++{
			fmt.Print(student_grades[key][i]," ")
			sum_grades +=student_grades[key][i]
		}
		fmt.Println()

		var len_grades float32 = float32(len(student_grades[key]))
		
		
		fmt.Println("Average score: ", float32(sum_grades)/len_grades )
		fmt.Println()
	}
	
 
}