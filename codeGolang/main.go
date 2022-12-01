package main

import (
	"bufio"
	"fmt"
	"os"
)

func Menu() {
	fmt.Println("\n\t1. Single List")
	fmt.Println("\t2. Binary Search Tree")
	fmt.Println("\t3. Hash Table")
	fmt.Println("\t4. Exit")
	fmt.Print("Enter your option: ")
}
func main() {
	fmt.Print(string(Reset))
	var op int
	Single := &ListLinked{}
	BST := &Tree{}
	HT := &Dictionary{}
	for op != 4 {
		Menu()
		fmt.Scan(&op)
		switch op {
		case 1:
			var idtemp int
			for op != 8 {
				fmt.Println("\n\t\tSingle List")
				Single.Menu()
				fmt.Scan(&op)
				switch op {
				case 1:
					fmt.Print("\n\tAdd 1 employee\n")
					Single.Add()
				case 2:
					fmt.Print("\n\tRemove 1 employee\n")
					Single.Remove()
				case 3:
					fmt.Print("\n\tEmployee\n")
					Single.Display()
				case 4:
					fmt.Print("\n\tSort\n")
					Single.SortBySalary()
					Single.Display()
				case 5:
					fmt.Print("\n\tSearch\n")
					fmt.Print("Enter id: ")
					fmt.Scan(&idtemp)
					NodeTemp := Single.Search(idtemp)
					if NodeTemp != nil {
						fmt.Printf("---------------------------------\n")
						Single.Export_1_Employee(NodeTemp.info)
					} else {
						fmt.Print("Not find!\n")
					}
				case 6:
					Single.ReadFile("SingleList.txt")
				case 7:
					Single.WriteFile("SingleList.txt")
				case 8:
					fmt.Print("EXIT!\n")
				default:
					fmt.Print("\nWrong! Please enter again!\n")
				}
			}
		case 2:
			for op != 5 {
				BST.Menu()
				fmt.Scan(&op)
				switch op {
				case 1:
					fmt.Println("\n\t\tAdd Quizz")
					BST.InsertNodeBSTs()
				case 2:
					fmt.Println("\n\t\tAll Quizz")
					BST.LNR(BST.root)
				case 3:
					fmt.Println("\n\t\tTake quizz")
					BST.DelKey()
					a := BST.GetArr()
					BST.TakeQuizz(a)
				case 4:
					fmt.Println("\n\t\tRead file")
					BST.ReadFile("quizz.txt")
				case 5:
					break
				default:
					fmt.Println("Wrong! Please Enter Again!")
				}
			}
		case 3:
			reader := bufio.NewReader(os.Stdin)
			var str string
			for op != 0 {
				HT.Menu()
				fmt.Scan(&op)
				switch op {
				case 1:
					fmt.Print("\n\tAdd word \n")
					HT.Insert()
				case 2:
					fmt.Print("\n\tUpdate word\n")
					fmt.Print("\nEnter word you want to update: ")
					reader.ReadString('\n') //xóa bộ nhớ đệm
					str = EnterString()
					HT.Update(str)
				case 3:
					fmt.Print("\n\tLookup\n")
					fmt.Print("\nEnter word you want to search: ")
					reader.ReadString('\n') //xóa bộ nhớ đệm
					str = EnterString()
					temp := HT.Search(str)
					if temp == nil {
						fmt.Println("Not find!")
					} else {
						fmt.Printf("%s (%s)  :%-10s\n", temp.vocal.english, temp.vocal.kind, temp.vocal.mean)
					}
				case 4:
					fmt.Print("\n\tDelete word\n")
					fmt.Print("\nEnter word you want to delete: ")
					reader.ReadString('\n') //xóa bộ nhớ đệm
					str = EnterString()
					HT.Del(str)
				case 5:
					fmt.Print("\n\tDisplay\n")
					HT.Display()
				case 6:
					fmt.Println("\n\tRead file")
					HT.ReadFile("vocal.txt")
				case 7:
					fmt.Println("\n\tWrite file")
					HT.WriteFile("vocal.txt")
				case 8:
					fmt.Println("\n\tRun Quizz")
					reader.ReadString('\n') //xóa bộ nhớ đệm
					arr := HT.GetArr()
					RunQuizz(HT, arr)
				case 9:
					fmt.Print("EXIT!\n")
				default:
					fmt.Print("\nWrong! Please enter again!\n")
				}
			}
		case 4:
			break
		default:
			fmt.Print("\nWrong! Please enter again!\n")
		}
	}
	fmt.Println("Done!!!")
}
