package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

const m = 101

type word struct {
	english, kind, mean string
}
type NodeTH struct {
	vocal word
	link  *NodeTH
}
type Dictionary struct {
	heads [m]*NodeTH
}

func (Dic *Dictionary) Menu() {
	fmt.Println("\n\t\tDICTIONARY")
	fmt.Println("\n\t\t1 .Add word ")
	fmt.Println("\t\t2. Update word ")
	fmt.Println("\t\t3. Search ")
	fmt.Println("\t\t4. Delele word ")
	fmt.Println("\t\t5. Display ")
	fmt.Println("\t\t6. Read file ")
	fmt.Println("\t\t7. Write file ")
	fmt.Println("\t\t8. Take quizz ")
	fmt.Println("\t\t9. Exit ")
	fmt.Print("\nEnter your option: ")
}
func (Dic *Dictionary) Standard(s string) string {
	return strings.ToLower(s)
}
func (Dic *Dictionary) Init(){
	for i := 0; i < m; i++ {
		Dic.heads[i] = nil;
	}
}
func (Dic *Dictionary) Create(x word) *NodeTH {
	x.english = Dic.Standard(x.english)
	p := &NodeTH{
		vocal: x,
		link:  nil,
	}
	return p
}
func (Dic *Dictionary) HashFunc(vocal string) int {
	key := 0
	vocal = strings.ToLower(vocal)
	for i := 0; i < len(vocal); i++ {
		key += int(vocal[i]) * (i + 1)
	}
	return key % m
}
func (Dic *Dictionary) InsertElement(w word) {
	k := Dic.HashFunc(w.english)
	p := Dic.Create(w)
	t := Dic.heads[k] //băm từ tiếng anh
	if t == nil {
		Dic.heads[k] = p
		fmt.Printf("\tAdded %s\n", Dic.heads[k].vocal.english)
	} else {
		//z := Dic.HashFunc(t.word.english);
		for t != nil {
			if t.vocal.english == p.vocal.english && t.vocal.kind == p.vocal.kind {
				fmt.Println("Trung Tu")
				break
			}
			t = t.link
		}
		if t.vocal.english != p.vocal.english && t.vocal.kind != p.vocal.kind {
			t.link = p
			fmt.Printf("\tAdded %s\n", Dic.heads[k].vocal.english)
		}
	}
}
func (Dic *Dictionary) Insert() {
	var w word
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n');
	fmt.Print("\tEnter word: ")
	w.english, _ = reader.ReadString('\n')
	fmt.Print("\tEnter type: ")
	w.kind, _ = reader.ReadString('\n')
	fmt.Print("\tEnter mean: ")
	w.mean, _ = reader.ReadString('\n')
	//cắt kí tự \n
	w.english = w.english[:len(w.english)-2] 
	w.kind = w.kind[:len(w.kind)-2]
	w.mean = w.mean[:len(w.mean)-2]
	Dic.InsertElement(w)
}
func (Dic *Dictionary) DisplayWord(p *NodeTH) {
	for p != nil {
		fmt.Printf("%s (%s)  :%-10s\n", p.vocal.english, p.vocal.kind, p.vocal.mean)
		p = p.link
	}
}
func (Dic *Dictionary) Display() {
	count := 0
	for i := 0; i < m; i++ {
		if Dic.heads[i] != nil {
			fmt.Printf("\n\t===BUCKET %v ===\n", i)
			Dic.DisplayWord(Dic.heads[i])
			count++
		}
	}
	fmt.Printf("co %v word\n", count)
}
func (Dic *Dictionary) Search(w string) *NodeTH {
	w = Dic.Standard(w)
	p := Dic.heads[Dic.HashFunc(w)];
	if p != nil {
		for p != nil {
			if p.vocal.english == w {
				return p
			}
			p = p.link
		}
	}
	return p
}
func (Dic *Dictionary) Del(w string) {
	w = Dic.Standard(w)
	p := Dic.Search(w)
	if p != nil {
		t := Dic.heads[Dic.HashFunc(w)]
		var q *NodeTH
		for t.vocal.english != w {
			q = t
			t = t.link
		}
		//Delete first
		if q == nil {
			Dic.heads[Dic.HashFunc(w)] = nil
		} else {
			q.link = t.link
		}
		fmt.Println("\t\tDone! ")
	} else {
		fmt.Println("\tNot found! ")
	}
}
func (Dic *Dictionary) Update(w string){
	w = Dic.Standard(w);
	p := Dic.Search(w);
	if p != nil{
		fmt.Print("\n\tEnter new word ")
		fmt.Print("\nWord: ")
		p.vocal.english = EnterString();
		fmt.Print("Type: ")
		p.vocal.kind = EnterString();
		fmt.Print("Mean: ")
		p.vocal.mean = EnterString();
		fmt.Println("Done!")
	}else{
		fmt.Println("Not find!")
	}
}
//use for enter 1 value has type of string
func EnterString() string{
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-2];
	return str;
}

//Read and write file

//check error
func CheckNilErrorTH( err error ){
	if err != nil{
		panic(err)
	}
}
//read file
func (Dic *Dictionary ) ReadFile( FileName string){
	var x word;
	file,err := os.Open(FileName)
	CheckNilErrorTH(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		items := strings.Split(line, "|")
		x.english = items[0];
		x.kind = items[1];
		x.mean = items[2];
		Dic.InsertElement(x);
	}
	fmt.Println("**Done!**")
}
//change struct employee to string
func ToString( x word ) string{
	return fmt.Sprintf("%s|%s|%s", x.english, x.kind, x.mean)
}
//put list to array of string
func (Dic *Dictionary) ArrString() []string {
	arr := []string{}
	for i := 0; i < m; i++ {
		if Dic.heads[i] != nil{
			p := Dic.heads[i];
			for p != nil{
				arr = append( arr, ToString(p.vocal));
				p = p.link;
			}
		}
	}
	return arr
}
//write list into file
func (Dic *Dictionary) WriteFile( FileName string ){
	lines := Dic.ArrString()
	file, err := os.OpenFile(FileName, os.O_RDWR, 0644)
	CheckNilErrorTH(err)
	defer file.Close()
	for _,val := range lines {
		_, err := fmt.Fprintln(file, val)
		CheckNilErrorTH(err)
	}
	fmt.Println("\nI'm Done!")
}

//take quizz
func (Dic *Dictionary) GetArr() []int {
	var arr []int;
	for i := 0; i < m; i++ {
		if Dic.heads[i] != nil{
			arr = append(arr,i);
		}
	}
	return arr
}
func RunQuizz(a *Dictionary, arr []int){
	rand.Seed(time.Now().UnixNano());
	var pos,n,count int;
	var str string;
	n = len(arr);
	var b []int;
	var p *NodeTH;
	for len(arr) != 0{
		pos = rand.Intn(len(arr)); // pos là vị trí phần tử trong mảng
		p = a.heads[arr[pos]]; // arr[pos] lưu giá trị của các bucket
		for p != nil{
			fmt.Printf("\tQuestion: %v\n", p.vocal.mean);
			str = EnterString();
			if str == p.vocal.english{
				fmt.Println("Correct!");
				count++;
			}else{
				fmt.Println("Incorrect!");
				b = append( b,arr[pos]);// nếu sai thì thêm vào mảng phụ để tí làm lại
			}
			fmt.Printf("%s (%s)  :%-10s\n", p.vocal.english, p.vocal.kind, p.vocal.mean)
			p = p.link;
		}
		arr = append(arr[:pos], arr[pos+1:]...);//delete a element of slice
	}
	fmt.Printf("\nYou Done %v/%v\n",count,n);
	if len(b) > 0 {
		fmt.Print("You have some questions wrong! Do you want do test again?(y/n): ")
		k := EnterString();
		if k == "Y" || k == "y"{
			RunQuizz(a,b)
		}else{
			fmt.Println("\n\tEND GAME")
		}
	}
}
// func main() {
// 	Dic := Dictionary{};
// 	Dic.Init();
// 	op := 1
// 	var str string
// 	var temp *NodeTH;
// 	reader := bufio.NewReader(os.Stdin)// biến để đọc
// 	for op != 0 {
// 		Dic.Menu()
// 		fmt.Scan(&op)
// 		switch op {
// 		case 1:
// 			fmt.Print("\n\tAdd word \n");
// 			Dic.Insert();
// 		case 2:
// 			fmt.Print("\n\tUpdate word\n");
// 			fmt.Print("\nEnter word you want to update: ");
// 			reader.ReadString('\n') //xóa bộ nhớ đệm
// 			str = EnterString();
// 			Dic.Update(str);
// 		case 3:
// 			fmt.Print("\n\tLookup\n")
// 			fmt.Print("\nEnter word you want to search: ")
// 			reader.ReadString('\n') //xóa bộ nhớ đệm
// 			str = EnterString();
// 			temp = Dic.Search(str)
// 			if temp == nil{
// 				fmt.Println("Not find!")
// 			}else{
// 				fmt.Printf("%s (%s)  :%-10s\n", temp.vocal.english, temp.vocal.kind, temp.vocal.mean);
// 			}
// 		case 4:
// 			fmt.Print("\n\tDelete word\n")
// 			fmt.Print("\nEnter word you want to delete: ")
// 			reader.ReadString('\n') //xóa bộ nhớ đệm
// 			str = EnterString();
// 			Dic.Del(str);
// 		case 5:
// 			fmt.Print("\n\tDisplay\n");
// 			Dic.Display();
// 		case 6:
// 			fmt.Println("\n\tRead file");
// 			Dic.ReadFile("vocal.txt");
// 		case 7:
// 			fmt.Println("\n\tWrite file");
// 			Dic.WriteFile("vocal.txt");
// 		case 8:
// 			fmt.Println("\n\tRun Quizz");
// 			reader.ReadString('\n') //xóa bộ nhớ đệm
// 			arr := Dic.GetArr();
// 			RunQuizz(&Dic, arr);
// 		case 0:
// 			fmt.Print("EXIT!\n");
// 		default:
// 			fmt.Print("\nWrong! Please enter again!\n");
// 		}
// 	}
// }
