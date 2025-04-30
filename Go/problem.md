
Write a program in Go that implements a 3-stage concurrent pipeline. There are 2 concurrent processes in state 1, two more concurrent processes in stage 2, and a single process in stage 3. All processes (in all stages) are executing comcurrently. They will use channels (not shared memory) to communicate and coordinate.

The detailed specs are:
```
  Uses a buffered channel "inCh" as a bounded data buffer (capacity = 5).
  Uses a second buffered channel "outCh" as a bounded data buffer (capacity = 5).

  Stage 1: Two Producers
    Producer 1 generates odd integers 1,3,5,7... up to 29
    Producer 2 generates even integers 2,4,6,... up to 30
    Both send into a shared buffered channel "inCh", (buffer size 5)
    control speed of the producers so that each rests a random time between values 
       After generating a value, make each process sleep for a random time between
       0 and 1500 milliseconds

  Stage 2: Two Consumers
    Each receives values from "inCh"
    Each squares the value and sends it to another buffered channel "outCh" 
    Simulate processing time via time.Sleep() ... 
       each process sleeps a random time between 0 and 3 seconds before generating
       another output

  Stage 3: Final Filter
    A single goroutine reads from "outCh"
    It prints the first value it receives.  
    It continues to read numbers from "outCh" but after printing the first,
       it only prints a value it receives if that value is strictly greater than 
       the previously printed value
    Discard any value that is not larger than the previous 
    So the final filter produces a strictly increasing sequence

  The program must:
    Print each processed result (from "outCh") exactly zero times, or once
    Exit cleanly when all data is processed
    Ensure no goroutine leaks or deadlocks
    The output should be some subsequence of the numbers 1^2, 2^2, 3^2, 4^2, ... 30^2
```

Questions for Discussion (readme file)
(1) What argument can you give that there will be no dealock among these processes?

(2) We get "easy" fan-out and fan-in here with Go channels, because a channel is a data structure that is not bound to any single process. How would we handle this situation in Elixir, where the mailbox semantics in actors mean that our "channel" is always attached to a single process? How do we have 2 producers creating and distributing values to 2 consumers?

