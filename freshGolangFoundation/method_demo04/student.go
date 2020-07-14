package main

import "fmt"

type student struct {
	id int		//学号是唯一的
	name string
	class string
}

// newStudent是 student类型的构造函数
func newStudent(id int, name,class string)  *student{
	return &student{
		id: id,
		name: name,
		class: class,
	}
}

// 学员管理的类型
type studentMar struct {
	allStudent []*student
}
//newStudentMar 是studentMar 的构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudent: make([]*student,0,100),
	}
}

// 1.添加学生
func (s *studentMar) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent,newStu)
}

//2、编辑学生
func (s *studentMar) modifyStudent(newStu *student) {
	for i,v := range s.allStudent{
		if newStu.id == v.id{		// 当学号相同时，就表示找到了要修改的学生
			s.allStudent[i] = newStu  // 根据切片的索引直接把学生赋值进来
			return
		}
	}
	fmt.Printf("输入的学生未找到:%d \n",newStu.id)
}

// 3。展示学生
func (s *studentMar) showStudent() {
	for _,v:= range s.allStudent{
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n",v.id,v.name,v.class)
	}
}