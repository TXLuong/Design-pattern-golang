## When to Use

- You have many different variants of an algorithm.
- You want to isolate the business logic of a class from the implementation details of algorithms that are not important to the invoker class.
- You have massive conditional statements that switch between different variants of the same algorithm.

## How to Implement

- Detect an algorithm that is prone to frequent changes.
- Declare an interface that is common to all variants of the algorithm.
- Implement the interface in all algorithm variants.
- Clients associate the desired variant to define the behavior they want.

## Pros

- Algorithms can be swapped inside the object at runtime.
- Variants are isolated from each other.
- Follows the Open/Closed Principle: you can create new strategies without modifying existing code.

## Cons

- Overkill when you only have a couple of algorithm variants.
- Clients must select the appropriate variant when using the algorithm.
- Many programming languages now support functional programming; passing functions as parameters can sometimes replace the need for this pattern.
