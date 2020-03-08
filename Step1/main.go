package main

import (
	"fmt"
	"os"
)

//学习了Go语言的基本使用，创建了简单的学生信息管理系统
//需求：

//1.添加学生信息
//2.修改学生信息
//3.展示学生信息

//菜单栏
func showMenu(){
	fmt.Println("欢迎登陆学生信息管理系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.修改学生信息")
	fmt.Println("3.展示所有学生信息")
	fmt.Println("4.退出系统")
}

//获取用户输入的学生信息
func getInput() *student {
	var (
		id string
		name string
		class string
		course string
		score int
	)
	//获取到用户输入的学生信息
	fmt.Println("请按要求输入学生信息")
	fmt.Print("请输入学生的学号：")
	fmt.Scanf("%s\n",&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanf("%s\n",&name)
	fmt.Print("请输入学生的班级：")
	fmt.Scanf("%s\n",&class)
	fmt.Print("请输入学生的课程：")
	fmt.Scanf("%s\n",&course)
	fmt.Print("请输入学生的分数：")
	fmt.Scanf("%d\n",&score)
	fmt.Print("成功输入信息\n")
	//调用结构体的默认构造方法
	stu := newStudent(id,name,class,course,score)
	return stu
}

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
		fmt.Printf("学号：%s 姓名：%s 班级：%s 课程：%s 分数：%d\n",v.id,v.name,v.class,v.course,v.score)
	}
}
func main() {
	//结构体的默认构造方法
	sm := newStudentMgr()
	for{
	//1.打印系统菜单
	showMenu()
	//2.获取用户输入
	var input int
	fmt.Print("请输入您要操作的序号：")
	fmt.Scanf("%d\n",&input)
	fmt.Println("用户输入的是：",input)
	//3.执行用户操作
		switch input {
		case 1:
			//添加学生
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			//修改学生信息
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			//展示所有学生信息
			sm.showStudent()
		case 4:
			//退出系统
			os.Exit(0)
		}
	}
}
