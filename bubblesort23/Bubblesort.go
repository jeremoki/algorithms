package main

func main (){
	var x = [] int { 3,6,2,4}
	var  isDone   = false

	for !isDone {
		isDone =true
		var i = 0
		for i < len(x)-1{
			if  x[i]<x[i+1]{
				x[i], x[i+1] = x[i + 1], x[i]
				println(isDone)
			}
			i++
		}

	}
}






