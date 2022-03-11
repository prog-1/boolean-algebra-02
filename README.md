# Boolean Algebra

## Disjunctive Normal Form

Write a program that finds a Boolean function in a [disjunctive normal form](https://en.wikipedia.org/wiki/Disjunctive_normal_form) (DNF)
for a given truth table.

Implement a function `func DNF(table [][]int) string` that returns a string representing the Boolean function in DNF for the truth table.

The truth table can be defined for an arbitrary number of variables.
For `n` variables, the truth table must contain `n^2` rows (all permutations of `n` variables) and `n+1` columns.
The last column defines the function result for a given combination of `n` variables.

Write tests for the function.

### Example

The truth table is defined as

```
table := [][]int{
	{0, 0, 1},
	{0, 1, 0},
	{1, 0, 1},
	{1, 1, 1},
}
```

The truth table contains two variables. There are 4 rows, because there are `2^2` value combinations for two variables.
The last column defines the function result.

`DNF(table)` must return `(!a_0 | !a_1) & (a_0 | !a_1) & (a_0 | a_1)`, where `a_0` and `a_1` are the function variables.

## Conjunctive Normal Form (optional, +100%)

Similarly to the program that calculates a Boolean function in DNF, write a program that finds a Boolean function in [conjunctive normal form](https://en.wikipedia.org/wiki/Conjunctive_normal_form).

Implement a function `func CNF(table [][]int) string` that returns a string representing the Boolean function in CNF for the truth table.

Write tests for the function.

*Note: For this task, you need to figure out an algorithm that translates a truth table to CNF.*

## 4-Bit Full Adder

Design a 4-Bit full adder in [https://logic.ly](https://logic.ly/demo/samples). You may use "1-Bit Full Adder" design from samples.

Make a circuit screenshot and upload it to this repository.
