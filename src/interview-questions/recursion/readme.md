## Recursion

<hr>

In computer science, recursion is a method of solving a problem where the solution depends on solutions to smaller instances of the same problem.
Such problems can generally be solved by iteration, but this needs to identify and index the smaller instances at programming time.
Recursion solves such recursive problems by using functions that call themselves from within their own code.
The approach can be applied to many types of problems, and recursion is one of the central ideas of computer science.

<h1> EXAMPLE</h1>

````
void recursiveFunction(int num) {
    printf("%d\n", num);
    if (num < 4)
        recursiveFunction(num + 1);
}
````

<img src="https://upload.wikimedia.org/wikipedia/commons/a/a7/Recursive1.svg">

<h3>Useful Links</h3>
<a src="https://en.wikipedia.org/wiki/Recursion_(computer_science)">Wikipedia</a>,
<a src="https://www.sparknotes.com/cs/recursion/whatisrecursion/section1/">sparknotes.com</a>
