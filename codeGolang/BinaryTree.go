package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	runCmd("cmd", "/c", "cls")
}

type Quizz struct {
	stt                          int
	question, a, b, c, d, answer string
}
type NodeBST struct {
	info        Quizz
	left, right *NodeBST
}
type Tree struct {
	root *NodeBST
}

var arr = []Quizz{
	{8, "\n\t\tvar int a = 20\n\t\tvar b = *a\n\t\tb = 30\n\t\tfmt.Println(a)\n", "20", "30", "địa chỉ của a", "lỗi biên dịch", "d"},
	{3, "\n\ta := []int{1,2,3,4,5,6}\n\ta = append(a[:3], a[4:]...)\n\tfmt.Println(a)\nchương trình sẽ xuất ra gì?", "1,2,3,4,5,6", "1,3,5", "1,2,3,5,6", "lỗi biên dịch", "c"},
	{5, "Đâu là kiểu khai báo mảng SAI?", "var a = [3]int", "a := [3]int{1,0,2}", "a := [...]int{1,2,3}", "var a int[]", "d"},
	{7, "\n\ta := 1\n\tif a >= 1{\n\t\tfmt.Println(a*10)\n\t} else{\n\t\tfmt.Println(0)\n\t}\nchương trình xuất ra a = ?", "10", "0", "lỗi", "1", "a"},
	{10, "\n\ta := 1\n\tswitch a{\n\t\tcase 1:\n\t\tcase 3:\n\t\t\tfmt.Print('có 31 ngày')\n\t\tcase 2:\n\t\t\tfmt.Print('có 28 hoặc 29 ngày')\n\t\t}\nChương trình sẽ xuất ra gì?", "Có 31 ngày", "có 29 ngày", "có 28 ngày", "có cái nịt", "d"},
}

// menu
func (Tr *Tree) Menu() {
	fmt.Println("\n\t1. add")
	fmt.Println("\t2. display quizz")
	fmt.Println("\t3. take quizz")
	fmt.Println("\t4. read file")
	fmt.Println("\t5. exit")
	fmt.Print("Enter your option: ")
}
func (Tr *Tree) Create(x Quizz) *NodeBST {
	p := &NodeBST{
		info:  x,
		left:  nil,
		right: nil,
	}
	return p
}

// enter 1 quizz fron keyboard
func (Tr *Tree) Enter_1_Quizz() *NodeBST {
	var x Quizz
	fmt.Print("\nEnter stt: ")
	fmt.Scan(&x.stt)
	reader := bufio.NewReader(os.Stdin) // initialize new reader
	reader.ReadString('\n')             // delete \n
	fmt.Print("Enter question: ")
	x.question, _ = reader.ReadString('\n')
	fmt.Print("Enter answer A: ")
	x.a, _ = reader.ReadString('\n')
	fmt.Print("Enter answer B: ")
	x.b, _ = reader.ReadString('\n')
	fmt.Print("Enter answer C: ")
	x.c, _ = reader.ReadString('\n')
	fmt.Print("Enter answer D: ")
	x.d, _ = reader.ReadString('\n')
	fmt.Print("Enter correct answer(a-d): ")
	x.answer, _ = reader.ReadString('\n')
	x.question = x.question[:len(x.question)-2] // remove the \n charactere
	x.a = x.a[:len(x.a)-2]
	x.b = x.b[:len(x.b)-2]
	x.c = x.c[:len(x.c)-2]
	x.d = x.d[:len(x.d)-2]
	x.answer = x.answer[:len(x.answer)-2]
	p := Tr.Create(x)
	return p
}

// display 1 node
func (Tr *Tree) Display_1_Quizz(x Quizz) {
	fmt.Printf("\n\t\tQuestion %v\n", x.stt)
	fmt.Printf("\tQuestion is: %v\n", x.question)
	fmt.Print("\n\n")
	fmt.Printf("\n\n\tAnswer A: %v\n", x.a)
	fmt.Printf("\tAnswer B: %v\n", x.b)
	fmt.Printf("\tAnswer C: %v\n", x.c)
	fmt.Printf("\tAnswer D: %v\n", x.d)
	fmt.Print("\n\n")
	fmt.Printf("\tCorrect answer: %v\n\n", x.answer)
}

// insert  node to the tree
func (Tr *Tree) Insert(quiz Quizz) {
	t := Tr.Create(quiz)
	if Tr.root == nil {
		Tr.root = t
	} else {
		p := Tr.root
		var prev *NodeBST
		for p != nil {
			prev = p
			if p.info.stt == quiz.stt {
				fmt.Println("Trùng!")
				return
			} else if p.info.stt > quiz.stt { // position of nodeBST on the left of tree
				p = p.left
			} else { // position of nodeBST on the right
				p = p.right
			}
		}
		if prev.info.stt > quiz.stt {
			prev.left = t
		} else {
			prev.right = t
		}
	}
}

// insert node from array
func (Tr *Tree) InsertNodeBSTs() {
	for i := 0; i < len(arr); i++ {
		Tr.Insert(arr[i])
	}
	fmt.Println("Done!")
}

// traversal LNR
func (Tr *Tree) LNR(p *NodeBST) {
	if Tr.root != nil {
		if p != nil {
			Tr.LNR(p.left)
			Tr.Display_1_Quizz(p.info)
			Tr.LNR(p.right)
		}
	} else {
		fmt.Println("Có cái nịt!!!")
	}
}

// take quizz
var arr2 []Quizz

func (Tr *Tree) AddToArr(p *NodeBST) {
	if p != nil {
		arr2 = append(arr2, p.info)
		Tr.AddToArr(p.left)
		Tr.AddToArr(p.right)
	}
}
func (Tr *Tree) GetArr() []Quizz {
	Tr.AddToArr(Tr.root)
	return arr2
}
func (Tr *Tree) DelKey() {
	reader := bufio.NewReader(os.Stdin) // initialize new reader
	reader.ReadString('\n')             // delete \n
}
func (Tr *Tree) TakeQuizz(arr []Quizz) {
	rand.Seed(time.Now().UnixNano())
	var str string
	var pos, count, n int
	n = len(arr)
	var b []Quizz
	reader := bufio.NewReader(os.Stdin)
	ClearTerminal()
	for i := 1; len(arr) != 0; i++ {
		pos = rand.Intn(len(arr)) //random a position in slice
		fmt.Println("***********************************************************")
		fmt.Print(string(Yellow))
		fmt.Printf("\t\t\tQuestion %v: %v\n", i, arr[pos].question)
		fmt.Print(string(Reset))
		fmt.Printf("\tA. %v\n", arr[pos].a)
		fmt.Printf("\tB. %v\n", arr[pos].b)
		fmt.Printf("\tC. %v\n", arr[pos].c)
		fmt.Printf("\tD. %v\n", arr[pos].d)
		fmt.Print(string(Cyan))
		fmt.Printf("Chọn --> ")
		str, _ = reader.ReadString('\n')
		fmt.Print(string(Reset))
		str = str[:len(str)-2]
		str = strings.ToLower(str)
		if str == arr[pos].answer {
			fmt.Println(string(Green), "\t\t\tCORRECT!", string(Reset))
			count++
		} else {
			fmt.Print(string(Cyan))
			fmt.Printf("Đáp án đúng là : -->   %s %s %v %s   <--\n", string(Reset), string(Yellow), strings.ToUpper(arr[pos].answer), string(Cyan))
			fmt.Print(string(Reset))
			fmt.Println(string(Red), "\t\t\tINCORRECT!", string(Reset))
			b = append(b, arr[pos]) // add the wrong question to the second array
		}
		fmt.Print("Continue..!")
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := consoleReader.ReadByte()
		ascii := input
		if ascii == 13 {
			ClearTerminal()
		}
		arr = append(arr[:pos], arr[pos+1:]...) // delete the question just answered
	}
	fmt.Printf("\tYou done %v/%v\n", count, n)
	if len(b) > 0 {
		fmt.Print("You have some questions wrong! Do you want do test again?(y/n): ")
		str, _ = reader.ReadString('\n')
		str = str[:len(str)-2]
		str = strings.ToLower(str)
		if str == "y" {
			Tr.TakeQuizz(b)
		} else {
			fmt.Println("\t\tEND GAME!!!")
		}
	}
}

// read file
// check error
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// read file
func (Tr *Tree) ReadFile(FileName string) {
	var x Quizz
	file, err := os.Open(FileName)
	CheckNilError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		x.stt, _ = strconv.Atoi(items[0])
		x.question = items[1]
		x.a = items[2]
		x.b = items[3]
		x.c = items[4]
		x.d = items[5]
		x.answer = items[6]
		Tr.Insert(x)
	}
	fmt.Println("**Done!**")
}

// func main(){
// 	Tr := Tree{}
// 	var op int;
// 	for op != 7{
// 		Tr.Menu();
// 		fmt.Scan(&op)
// 		switch op{
// 		case 1:
// 			fmt.Println("\n\t\tAdd Quizz")
// 			Tr.InsertNodeBSTs();
// 		case 2:
// 			fmt.Println("\n\t\tAll Quizz")
// 			Tr.LNR(Tr.root);
// 		case 3:
// 			fmt.Println("\n\t\tTake quizz")
// 			DelKey();
// 			a := Tr.GetArr();
// 			Tr.TakeQuizz(a);
// 		case 4:
// 			fmt.Println("\n\t\tRead file")
// 			Tr.ReadFile("quizz.txt")
// 		case 5:
// 			break;
// 		default:
// 			fmt.Println("Wrong! Please Enter Again!");
// 		}
// 	}
// 	fmt.Println("Done!")
// }
