package greetings

import(
"errors"
"fmt"
"math/rand"
)

func Hello(name string) (string, error){
	if name == "" {
		return "", errors.New("No name provided")
	}
	message:= fmt.Sprintf(randomGreetings(), name)
	return message, nil;
} 

func randomGreetings() string{
	message:= []string{
		"Hey %v, how u doing",
		"Long time no see %v",
		"Hope you are doing well %v",
	}

	return message[rand.Intn(len(message))]
}
