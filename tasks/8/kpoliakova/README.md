### Poliakova Ksenia

#### Problem A - Key Word in Context (KWIC)

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

#### Problem B - Eight Queens (8Q)

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
