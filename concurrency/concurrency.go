package main

import(
"fmt"
"time"
"sync"
"math/rand"
"strconv"
)

func main() {
	//runfunctionsconcurrently()
	uploadPicturesConcurrently()
}

func uploadPicturesConcurrently() {
	var wg sync.WaitGroup
	for i:=1; i<100; i++ {		
		wg.Add(1)
		go func(pic string) {			
			uploadPicture("Pic " + pic + ".jpg")
			wg.Done()
		}(strconv.Itoa(i))
	}
	wg.Wait()
}

func uploadPicture(path string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	processingTime := 200 * r1.Intn(15)
	time.Sleep(time.Millisecond * time.Duration(processingTime))	
	fmt.Println(path, "uploaded in ", strconv.Itoa(processingTime), " milli seconds" )
}

func runfunctionsconcurrently() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("sheep") 
		wg.Done()
	}()
	
	wg.Add(1)
	go func() {
		count("goat")
		wg.Done()
	}()
	
	wg.Add(1)
	go func() {
		count("fish")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i:=0; i < 5;i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}		
}

