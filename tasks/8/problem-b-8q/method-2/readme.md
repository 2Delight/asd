# Task

**Problem B**. Eight Queens (8Q)

**Method 2**. Main/Subroutine with stepwise refinement

# How to run

1. Input your text in file input.txt

2. Run next command in your console: `python3 main.py <arg>`

arg - is size of field with queens

e.g  `python3 main.py 8`


# Questions


| Problem        | Method                                       | Solver           | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|----------------|----------------------------------------------|------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| 8 Queens (B)   | Main/Subroutine with stepwise refinement (2) | Anton Timonin    | -                | -                          | -                 | -                     | -             |
| 8 Queens (B)   | Pipes-and-filters (3)                        | Oleg Sidorenkov  | +                | -                          | +                 | -                     | -             |
| 8 Queens (B)   | Implicit invocation (4)                      | Ksenya Polyakova | +                | -                          | +                 | +                     | -             |


## Q1 
**In which case it is easier to change the implementation algorithm in each of the modules?** 

Pipes-and-filters. Easier to change independent parts of the system
Impicit invocation. Easier to enchange system capabilities by adding new listener. More isolated solution.

## Q2 
**In which solution it is easier (= seemingly less effort) to change data representation** 

Main/Subroutine. Hard to change data reprsentation, since all steps are dependent to each other.
Pipes-and-filters and Implicit invocation. Hard to change data representation, since all stages cooperate with each other.

## Q3 
**In which solution it is easier to add additional functions to the modules?**

Main/Subroutine with stepwise refinement. It is hard to add additional functions, because every subtask is depend on previous task.
Pipes-and-filters and Implicit invocation. Modules functionality more independent, so it is easy to change particular module without affection other modules.

## Q4
**Which solution is seemingly more performant?**

If you solve the problem not through recursion, but through interactivity, it will still not be a very productive solution through Main/Subroutine. But it is not best solution.
Implicit invocation (event-driven). This method is more effective for tasks where you can parallelize some kind of processing, as in this task. This reduces downtime, and there is no need to wait for the last part of the task to be completed.


## Q5
**If you are asked to implement a similar program, which of the solutions would you reuse?**

It is really difficult to choose among the listed methods. If I choose among them, I would choose main/subroutine. The code turns out to be very related, but it's easy to read

