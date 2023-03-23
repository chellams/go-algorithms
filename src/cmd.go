package main

func main() {
	// var students []*Student
	//
	// for i := 5; i > 0; i-- {
	// 	v := rand.Intn(100-15) + 15
	// 	student := &Student{
	// 		id:    i,
	// 		name:  fmt.Sprintf("Name %d", i),
	// 		score: v,
	// 	}
	// 	fmt.Print(student)
	// 	students = append(students, student)
	// }
	//
	// isPass := func(student Student) bool {
	// 	return student.score >= 35
	// }
	// passedStudents := passesStudent(students, isPass)
	// fmt.Println("Passed students", passedStudents)
	// var nums = []int{1, 0, 3, 4, 5, 0, 7, 8, 0, 10}
	// var errorNums []int
	// // errorNums = Test(0, len(nums), nums, errorNums)
	// fmt.Println(Test(0, len(nums), nums, errorNums))

}

func Test(l int, r int, nums []int, errorNums []int) []int {
	if l <= r {
		mid := l + (r-l)/2
		left := nums[l:mid]
		right := nums[mid:r]

		res := FindZero(left...)
		if res == -1 {
			if len(left) == 1 {
				return append(errorNums, left[0])
			}
			errorNums = Test(l, mid, nums, errorNums)
		}

		res = FindZero(right...)
		if res == -1 {
			if len(right) == 1 {
				return append(errorNums, right[0])
			}
			errorNums = Test(mid, r, nums, errorNums)
		}
	}

	return errorNums
}

func FindZero(nums ...int) int {
	for _, num := range nums {
		if num == 0 {
			return -1
		}
	}
	return 0
}
