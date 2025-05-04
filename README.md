# 590-Final

# Elixir Problem 
### To run: cd into 'Elixir' folder and run: ```elixir animal.exs```
Program should run automatically to simulate object oriented programming as specified. 

I've written Elixir code that simulates the OO pillars, similar to Java, but in a functional-based language. One of the main downsides to not having the same base functionality and class-based programming style as Java is the way that the 'Main' function and class have to be set up. Because there are no classes in Elixir, we have to work with inheriting from a module system instead of 

# Go Problem
### To run: cd into 'Go' folder and run ```go run problem.go```

Program will print out a subsequence of numbers of the form 1^2, 2^2...30^2. 

(1): Because this program runs using buffered channels, there is no chance for deadlock. This is because the processes to add ints to the buffered channel are initialized and add odd/even integers sequentially on a process from 0-1500 milliseconds. Since the channels are buffered to have a maximum of 5 numbers, the consumer channel will wait until the next number is received before sending it to the final filter function. This ensures that numbers are consumed by the cons1 and cons2 functions sequentially in the order that they were pushed to the buffered channel. If a number is attempted to be pushed to inCh, but the channel is full, the program will just block push because of the buffered channel properties. This ensures no deadlock and that every number is processed.