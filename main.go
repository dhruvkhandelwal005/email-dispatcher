package main

import "sync"

type Recipient struct {
	// Capital First letter means var is public(can be accessed by other functions as well) and if starting with small letter then it private to the function only
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipient("./emails.csv", recipientChannel)
	}()
	var wg sync.WaitGroup

	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel,&wg)

	}
	wg.Wait()

}
