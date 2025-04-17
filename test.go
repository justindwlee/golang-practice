package main

// func main(){
// 	c := make(chan string)
// 	people := []string{"nico", "haha", "justin", "timkerlake"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	for i:=0; i<len(people); i++{
// 		fmt.Print("waiting for", i)
// 		fmt.Println(<-c)
// 	}

// }

// func isSexy(person string, c chan string){
// 	time.Sleep(time.Second * 3)
// 	c <- person + " is sexy"
// }