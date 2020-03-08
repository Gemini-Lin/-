package main

import (
	"fmt"
)

//构造学生结构体
type student struct {
	id string
	name string
	class string
	course string
	score int
}

//返回类型为student指针的构造函数
func newStudent(id ,name,class, course string,score int) *student {
	return &student{
		id:     id,
		name:   name,
		class:  class,
		course: course,
		score:  score,
	}
}

//学生信息指针数组
type studentMgr struct {
	allStudents []*student
}

//初始化学生信息
func newStudentMgr() *studentMgr{
	return &studentMgr{
		//构造切片初始化，元素类型为student指针数组，元素数量为0，容量为1000
		allStudents:make([]*student,0,1000)}
}

//1.添加学生信息
//限定了接受者为studentMgr的指针的方法
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents,newStu)
}

//2.修改学生信息
func (s *studentMgr) modifyStudent(newStu *student) {
	//i为索引，v为数组
	for i, v := range s.allStudents{
		if newStu.id == v.id{//根据学号查找学生信息
			s.allStudents[i] = newStu
			return
		}
	}
	//提示用户输入有误
	fmt.Printf("输入的学生信息有误，系统中没有学号为:%s的学生\n",newStu.id)
}

//3.展示学生信息
func (s *studentMgr)showStudent() {
	for _, v := range s.allStudents{
		fmt.Print("学号：%s 姓名：%s 班级：%s 课程：%s 分数：%d\n",v.id,v.name,v.class,v.course,v.course)
	}
}