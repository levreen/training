# Test-driven development

## (1) Write the test
    You write a test for the feature you want, even though it doesn't exit yet. Then you run
    the test to ensure that it fails.
## (2) Make it pass
    You implement the feature in your main code. Don't worry about whether the code you're writing is sloppy or inefficient; your only goal is to get it working. Then you run the
    test to ensure that it passes.
## (3) Refactor your code
    Now, you're free to refactor your coe, to change and improve it, however you please. 

## This save me of incessant ask for logins. 
## Only work once. You need to run the cmd again. 
```git
$ git config --global credential.helper cache.
$ git config --global credential.helper 'store --file ~/.my-credentials'
```
## First-class functions
    Functions in Go are treated as "first-class citizens." In a programming language with first-class functions,
    functions can be assigned to variables, and then called from those variables.

    
