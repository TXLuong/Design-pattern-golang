When to use : 
 - You have many different variants of an algorithm
 - You want to isolate business logic of a class from the implementations details of algorithms that may not be important in the invoker class 
 - You have massive conditional statement that switches between differentvairants of the same algorithm
How to implement : 
 - Detect an algorithm that's prone to frequent changes.
 - Declare interface which is common for all variants fo the algorithm 
 - Implement the interface in all variants of algorithm 
 - Clients need to associate it as what they want to behave
Pros: 
 - Swap inside object at runtime 
 - Isolate variants 
 - Open/close principal.You create new strategies
Cons: 
 - You have just a couple of viriants of algorithm 
 - Clients need to select a proper vairant when using 
 - Lots of programing language now support functional programming, we can pass function as param then this strategy maybe not nessesary anymore 