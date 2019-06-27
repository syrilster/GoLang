# Functions
* Go variable and function declarations work in reverse order from Java.  
* The sayHello declaration here begins with the keyword func rather than a return type.  
* The return type, string in this case, comes at the very end just before the opening curly-brace.  The same is true for function parameters.  Instead of String listener, the parameter declaration is listener string.
* Variables are declared with the keyword var, with the type likewise following the variable name. 
* When the colon-equals operator := is used rather than a plain equals sign, both the var keyword and the variable type can be omitted.  The compiler is smart enough to deduce the correct type from what’s on the right-hand side on the operator.

# Scope and Lifecycle
* Because the variable message is declared inside of main, its scope is limited to that function.  
* currentTime is declared at the package level, and therefore is visible to all functions in the package.  
* Notice the lack of access level modifiers on the variables and functions here.  
* While Java uses the keywords public, private, etc to control access from outside of a class or package.  Go does this through capitalization.  A variable or function beginning with a lower-case character is analogous to private in Java.  
* If the currentTime variable had been declared like this instead: then it would be analogous to Java’s public, and would be accessible from other packages as well.  The same convention applies for functions.  Because it starts with an upper-case “P”, the Println function may be used outside of the standard fmt package.

```
var CurrentTime time.Time
```

* Lastly, notice the init() function.  Although main() is the primary entry-point for a Go application, init() is a special (optional) function that is automatically called before anything else.  
* It is typically used to initialize variables, or other setup and validation tasks.  For now, you can think of init() is roughly comparable to a Java constructor, even though in Go the distinction between static and instance methods doesn’t quite apply.

# Multiple Return Types?!?
* Java methods can take any number of input parameters, but always return no more than ONE return type.  
* If you need a method to return more than one distinct piece of information, then you would either create a custom type to hold the multiple pieces or else give up and refactor the method to avoid that clutter.
* One of the surprises with Go functions is that they quite commonly return multiple values in a single call!  
* Consider the example class funcMultipleReturnTypes which trims whitespace from a string, and returns both the trimmed string and the number of characters removed in the process.
