##
 Poliakova Ksenia

# Problem A - Key Word in Context (KWIC)

I solved this problem with method 2 (Main/Subroutine with stepwise refinement or Shared Data).
Text is hardcoded in program in variable text, for example:

```
text = "For example, a search query including all of the words in an example definition (KWIC is an acronym for Key Word In Context," + \
    " the most common format for concordance lines) and the Wikipedia slogan in English (the free encyclopedia), searched against a Wikipedia " + \
    "page, might yield a KWIC index as follows. A KWIC index usually uses a wide layout to allow the display of maximum 'in context' " + \
    "information (not shown in the following example)."
```

The result of algorith will be written in command line. So how to launch it:
```
python main.py
```

### Table for comparison of different methods for Problem A

| Method                                       | Solver           | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|----------------------------------------------|------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Shared Data (2)                              | Ksenia Poliakova | +                | -                          | +                 | +                     | -             |
| Abstract Data Types (1)                      | Oleg Sidorenkov  | -                | +                          | -                 | -                     | +             |
| Implicit invocation, event-driven (4)        | Egor Timonin     | +                | -                          | +                 | +                     | -             |

### Answering questions

#### a) in which case it is easier to change the implementation algorithm in each of the modules? 
It is easier to change the implementation algorithm in two methods: in shared data method and event-driven. Shared data method can handle changes better as it offers a simpler formula. And in event-driven modules are isolated and they communicate through events.
#### b) in which solution it is easier (= seemingly less effort) to change data representation
Abstract data types uses types with separation of concerns which means that every entity has responsibility for their own data and behaviour, so changing their representations could be less affective for other modules.
#### c) in which solution it is easier to add additional functions to the modules
In event driven method it is is easier to add new functionality without affecting other modules as new subscribers or handlers. On the other hand, in abstract data types method adding new functionality means that it is usually required to change related modules, because components more connected and have dependencies on each other.
#### d) which solution is seemingly more performant?
Event-driven method seems to be with high performance, because tasks are driven asynchronously, so that can save some time. Abstract Data types doesn't seem to be high performance, as there we use a lot of classes, they consume much memory, that's why Shared Data is better. 
#### e) if you are asked to implement a similar program, which of the solutions would you reuse?
Abstract Data Types method contains several classes, that's why it is easier to reuse.

# Problem B - Eight Queens (8Q)

I solved this problem with method 4 (Implicit invocation or event-driven).
To get the result, please get additional file all_possible_placements.txt from [here](https://disk.yandex.com/d/kawmuTKwB37F3g).

How to launch program:
```
python main.py
```

Output examples:
```
...


['*', '*', '*', '*', '*', '*', '*', 'Q']
['*', 'Q', '*', '*', '*', '*', '*', '*']
['*', '*', '*', '*', 'Q', '*', '*', '*']
['*', '*', 'Q', '*', '*', '*', '*', '*']
['Q', '*', '*', '*', '*', '*', '*', '*']
['*', '*', '*', '*', '*', '*', 'Q', '*']
['*', '*', '*', 'Q', '*', '*', '*', '*']
['*', '*', '*', '*', '*', 'Q', '*', '*']


['*', '*', '*', '*', '*', '*', '*', 'Q']
['*', '*', 'Q', '*', '*', '*', '*', '*']
['Q', '*', '*', '*', '*', '*', '*', '*']
['*', '*', '*', '*', '*', 'Q', '*', '*']
['*', 'Q', '*', '*', '*', '*', '*', '*']
['*', '*', '*', '*', 'Q', '*', '*', '*']
['*', '*', '*', '*', '*', '*', 'Q', '*']
['*', '*', '*', 'Q', '*', '*', '*', '*']


['*', '*', '*', '*', '*', '*', '*', 'Q']
['*', '*', '*', 'Q', '*', '*', '*', '*']
['Q', '*', '*', '*', '*', '*', '*', '*']
['*', '*', 'Q', '*', '*', '*', '*', '*']
['*', '*', '*', '*', '*', 'Q', '*', '*']
['*', 'Q', '*', '*', '*', '*', '*', '*']
['*', '*', '*', '*', '*', '*', 'Q', '*']
['*', '*', '*', '*', 'Q', '*', '*', '*']

Amount of solutions: 92
```

### Table for comparison of different methods for Problem B

| Method                                       | Solver           | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|----------------------------------------------|------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Implicit invocation, event-driven (4)        | Ksenia Poliakova | +                | -                          | +                 | +                     | -             |
| Pipes-and-filters (3)                        | Oleg Sidorenkov  | +                | -                          | +                 | -                     | -             |
| Abstract Data Types (1)                      | Egor Timonin     | -                | +                          | -                 | -                     | +             |

### Answering questions

#### a) in which case it is easier to change the implementation algorithm in each of the modules? 
#### b) in which solution it is easier (= seemingly less effort) to change data representation

#### c) in which solution it is easier to add additional functions to the modules
In abstract data types adding new functionality means that it is usually required to change related modules, because components more connected and have dependencies on each other. But for other methods it's easier to add additional functions.
#### d) which solution is seemingly more performant?
Event driven method is more effective for tasks where you can parallelize some kind of processing, as in this task. This reduces downtime, and there is no need to wait for the last part of the task to be completed.
#### e) if you are asked to implement a similar program, which of the solutions would you reuse?
The same thing as in the prevoius problem, I would reuse abstract data types method.
