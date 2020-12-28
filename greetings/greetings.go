package greetings 

import(
	"fmt"
	"math/rand"
	"time"
	"errors"
)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")		
	}
	
	message := fmt.Sprintf(getRandomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)		
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomGreeting() string {
	greets := []string {
		"Hello There %v!",
		"How are you doing %v!",
		"Hi, How do you do %v",
		"Hey, Whatsup %v!",
		"Hello Wonderful World %v",
		"Surprise! Here I am %v",
		"Merry Chrismas %v!",
		"Happy New Year %v!", 
	}
	return greets[rand.Intn(len(greets))]
}