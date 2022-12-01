package main
import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
type employee struct{
	ID,Salary int
	Name string
}
type NodeSL struct{
	info employee
	link *NodeSL
}
type ListLinked struct{
	First *NodeSL
}
//reate a node from a employee
func (l *ListLinked)Create( x employee ) *NodeSL{
	p := &NodeSL{
		info: x,
		link: nil,
	}
	return p
}
//insert a node at the first nodeSL of list
func ( l *ListLinked ) InsertFirst( x employee ){
	p := l.Create(x)
	if l.Search(p.info.ID) == nil { // not yet
		if l.First == nil {
			l.First = p
		} else {
			p.link = l.First
			l.First = p
		}
	}
}
//display a employee in hte screen
func (l *ListLinked)Export_1_Employee( x employee ){
	fmt.Printf("| %-5v| %-20v| %-4v$ |\n", x.ID, x.Name, x.Salary)
	fmt.Printf("--------------------------------------\n")
}
//display full list
func ( l *ListLinked ) Display(){
	p := l.First
	if p != nil{
		fmt.Printf("--------------------------------------\n")
		for ; p != nil; {
			l.Export_1_Employee(p.info)
			p = p.link
		}
	}
}
//enter a employee from keyboard
func (l *ListLinked)Enter_1_Employee( x *employee ) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\tEnter ID: ")
	fmt.Scan(&x.ID)
	reader.ReadString('\n')
	fmt.Print("\tEnter Name: ")
	x.Name, _ = reader.ReadString('\n')
	fmt.Print("\tEnter Salary: ")
	fmt.Scan(&x.Salary)
}
//add a node from keyboard
func ( l *ListLinked ) Add(){
	var x employee
	l.Enter_1_Employee(&x)
	x.Name = x.Name[:len(x.Name)-2]
	if l.Search(x.ID) != nil{
		fmt.Print("\nThis employee already exists!\n")
		return
	}
	l.InsertFirst(x)
	fmt.Print("\tDONE!\n")
}
//menu
func( l *ListLinked ) Menu(){
	fmt.Println("\n\t1. Add 1 employee")
	fmt.Println("\t2. Remove 1 empoyee")
	fmt.Println("\t3. Display list employee")
	fmt.Println("\t4. Sort by salary")
	fmt.Println("\t5. Search by id")
	fmt.Println("\t6. Read file")
	fmt.Println("\t7. Write file")
	fmt.Println("\t8. Exit")
	fmt.Print("Enter your option: ")
}
//search a employee from their id
func ( l * ListLinked ) Search( id int ) *NodeSL{
	p := l.First
	if p != nil{
		for p != nil{
			if p.info.ID == id{
				return p
			}
			p = p.link
		}
	}
	return nil
}
//remove a node
func ( l * ListLinked ) Remove(){
	var id int
	fmt.Print("\n\tEnter id of employee you want to remove from the list: ")
	fmt.Scan(&id)
	if l.Search(id) != nil{
		p := l.First
		var q *NodeSL
		for p != nil && p.info.ID != id{
			q = p
			p = p.link
		}
		//Delete first
		if p != nil{
			if p == l.First{
				l.First = l.First.link
			}else{
				q.link = p.link
			}
			fmt.Print("\nSUCCESS!")
		}
		return
	}else {
		fmt.Print("*Can't find that id!\n")
	}
}
//arrange list
func ( l *ListLinked ) SortBySalary(){
	var temp employee
	if l.First != nil{
		p := l.First
		var q *NodeSL
		for ; p.link != nil; p = p.link{
			for q = p.link;q != nil; q = q.link{
				if q.info.Salary > p.info.Salary{
					temp = q.info
					q.info = p.info
					p.info = temp
				}
			}
		}
		fmt.Print("\nDONE!\n")
	}
}
//check error
func (l *ListLinked)CheckNilError( err error ){
	if err != nil{
		panic(err)
	}
}
//read file
func ( l *ListLinked ) ReadFile( FileName string){
	var x employee
	file,err := os.Open(FileName)
	l.CheckNilError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		items := strings.Split(line, ",")
		x.ID,_ = strconv.Atoi(items[0])
		x.Salary,_ = strconv.Atoi(items[2])
		x.Name = items[1]
		l.InsertFirst(x)
	}
	fmt.Println("**Done!**")
}
//change struct employee to string
func (l *ListLinked)ToString( x employee ) string{
	return fmt.Sprintf("%v,%s,%v", x.ID, x.Name, x.Salary)
}
//put list to array of string
func (l *ListLinked ) ArrString() []string {
	p := l.First
	arr := []string{}
	for p != nil {
		arr = append( arr, l.ToString(p.info))
		p = p.link
	}
	return arr
}
//write list into file
func (l *ListLinked ) WriteFile( FileName string ){
	lines := l.ArrString()
	file, err := os.OpenFile(FileName, os.O_RDWR, 0644)
	l.CheckNilError(err)
	defer file.Close()
	for _,val := range lines {
		_, err := fmt.Fprintln(file, val)
		l.CheckNilError(err)
	}
	fmt.Println("\nI'm Done!")
}
// func main(){
// 	myList := ListLinked{}
// 	var op,idtemp int
// 	var NodeTemp *Node
// 	for op != 8{
// 		myList.Menu()
// 		fmt.Scan(&op)
// 		switch op{
// 		case 1:
// 			fmt.Print("\n\tAdd 1 employee\n")
// 			myList.Add()
// 		case 2:
// 			fmt.Print("\n\tRemove 1 employee\n")
// 			myList.Remove()
// 		case 3:
// 			fmt.Print("\n\tEmployee\n")
// 			myList.Display()
// 		case 4:
// 			fmt.Print("\n\tSort\n")
// 			myList.SortBySalary()
// 			myList.Display()
// 		case 5:
// 			fmt.Print("\n\tSearch\n")
// 			fmt.Print("Enter id: ")
// 			fmt.Scan(&idtemp)
// 			NodeTemp = myList.Search(idtemp)
// 			if NodeTemp != nil{
// 				fmt.Printf("---------------------------------\n")
// 				Export_1_Employee(NodeTemp.info)
// 			}else{
// 				fmt.Print("Not find!\n")
// 			}
// 		case 6:
// 			myList.ReadFile("test.txt")
// 		case 7:
// 			myList.WriteFile("test.txt")
// 		case 8:
// 			fmt.Print("EXIT!\n")
// 		default:
// 			fmt.Print("\nWrong! Please enter again!\n")
// 		}
// 	}
// }