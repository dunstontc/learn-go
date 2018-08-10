package main

import "fmt"

func main() {
	var x [58]string

	// ascii alphabet
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)     // [A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \ ] ^ _ ` a b c d e f g h i j k l m n o p q r s t u v w x y z]
	fmt.Println(x[42]) // k
}
