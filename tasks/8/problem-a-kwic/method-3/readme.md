# Task

**Problem A**. Key Word in Context (KWIC)

**Method 3**. Pipes-and-filters

# How to run

1. Input your text in file input.txt

2. Run next command in your console: `python3 main.py`



# Questions

| Problem                 | Method                                       | Solver           | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|-------------------------|----------------------------------------------|------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Key Word in Context (A) | Abstract Data Types (1)                      | Oleg Sidorenkov  | -                | +                          | -                 | -                     | +             |
| Key Word in Context (A) | Pipes-and-filters (3)                        | Anton Timonin    | +                | -                          | +                 | -                     | +             |
| Key Word in Context (A) | Main/Subroutine with stepwise refinement (2) | Ksenya Polyakova | +                | -                          | +                 | +                     | -             |

## Q1
**In which case it is easier to change the implementation 
algorithm in each of the modules?** 

In case of pipes and filters and main/subroutine with stepwise refinement.
In pipes and filters it is easy to change, since module more isolated. And they could be changes without affecting entire system. 
In subroutine module easier to change concerete subtasks in the hole algorithm.

## Q2 
**In which solution it is easier (= seemingly less effort) to change data 
representation?**

In case of abstract data types is easier to change data types. Since types are more isolated in this method.

## Q3 
**In which solution it is easier to add additional functions to the modules?**

Pipes-and-Filters and Main/Subroutine make it easiest to add new functionality.
Pipes-and-Filters solution is organized like independent modules interact to each other and it is easy to change implementation of concerete module
Main/Subroutine. In this case it is easier, because all solution is divided to a small tasks, then it is easier to change subtusk without affection another task.

## Q3
**Which solution is seemingly more performant?**

In case of main/subroutine it the most performance efficient solution. 
Since it doesn't require additional types conversion and decomposition of processing steps. 

## Q4
**If you are asked to implement a similar program, 
which of the solutions would you reuse?**

I'd use pipes and filters, since for this task, 
when we should iteratively perform some work, it easier to implement/change 
and maintain solution in this way.