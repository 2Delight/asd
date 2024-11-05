### Rationale for eliciting classes for problem B, method 1

I extracted class Queen to separate logic of checking availability of the place 
from class QueenProblemResolver which defines the logic for the solution and store solutions for the problem

| Problem                 | Method                                | Solver          | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|-------------------------|---------------------------------------|-----------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Key Word in Context (A) | Implicit invocation, event-driven (4) | Egor Timonin    | +                | -                          | +                 | +                     | -             |
| Key Word in Context (A) | Pipes-and-filters (3)                 | Anton Timonin   | +                | -                          | +                 | -                     | -             |
| Key Word in Context (A) | Abstract Data Types (1)               | Oleg Sidorenkov | -                | +                          | -                 | -                     | +             |

a) In which case it is easier to change the implementation algorithm in each of the modules?
It is easy to change algorithm in event-driven (ED) solution owing to separation of modules. They are less coupled, because communication happening through events.
In pipes-and-filters (PF) approach it is also easier to change modules, because they are independently solve their tasks.
Whereas Abstract Data Types (ADR) has more coupled modules, which means, if we want to change algorithm we have typically required to change both modules. 

b) In which solution it is easier (= seemingly less effort) to change data representation  
In ED solution if subscribers would be more than one, it is difficult to change representation for one subscriber without affecting other, 
so it is quite difficult to change data representation using this approach.
In PF method it would be required to change several modules if we changed one of them, because following component have to support updated representation.
While ADT approach uses types with separation of concerns which means that every entity has responsibility for their own data and
behaviour, so changing their representations could be less affective for other modules.

c) In which solution it is easier to add additional functions to the modules
Owing to ED design and separation of modules it is easier to add new functionality without affecting other modules as new subscribers or handlers.
Modules if PF approach also separated, so it would be easier to add new functionality without affecting other.
Whereas in ADT adding new functionality means that it is usually required to change related modules, because components more connected and have dependencies on each other.

d) Which solution is seemingly more performant?
Owing to ED design it is easier to make modules with high performance, because listeners process tasks asynchronous
which means that they could handle tasks without waiting when previous event be processed.
PF design seems less performance, because tasks are processed sequentially.
In that time ADT using only required classes to process their tasks without difficulties with ED and PF design approaches,
so this method could be performing if we required to solve task sequentially, while with asynchronous operations this method could be less performant than ED method

e) If you are asked to implement a similar program, which of the solutions would you reuse?
ADT method seems more useful for tasks which could have a good separation of modules with well-defined responsibilities and behaviours.

| Problem          | Method                                | Solver           | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|------------------|---------------------------------------|------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Eight Queens (B) | Abstract Data Types (1)               | Egor Timonin     | -                | +                          | -                 | -                     | -             |
| Eight Queens (B) | Shared Data (2)                       | Anton Timonin    | -                | -                          | -                 | -                     | -             |
| Eight Queens (B) | Implicit invocation, event-driven (4) | Ksenia Polyakova | +                | -                          | +                 | +                     | -             |

a) In which case it is easier to change the implementation algorithm in each of the modules?
Abstract Data Types (ADR) has more coupled modules, which means, if we want to change algorithm we have typically required to change both modules.
In Shared Data (SD) it would be harder to change algorithm, because several components work with shared source, so changes would affect several modules.
It is easy to change algorithm in event-driven (ED) solution owing to separation of modules. They are less coupled, because communication happening through events.

b) In which solution it is easier (= seemingly less effort) to change data representation  
ADT approach uses types with separation of concerns which means that every entity has responsibility for their own data and
behaviour, so changing their representations could be less affective for other modules.
In SD it would be challenging to change data representation without affection of other components, because all of them use equal data source.
In ED solution if subscribers would be more than one, it is difficult to change representation for one subscriber without affecting other,
so it is quite difficult to change data representation using this approach.

c) In which solution it is easier to add additional functions to the modules
In ADT adding new functionality means that it is usually required to change related modules, because components more connected and have dependencies on each other.
If additional functionality uses existing information, so it should be easy to add new functionality. If new functionality require to change data, it would affect other modules.
Owing to ED design and separation of modules it is easier to add new functionality without affecting other modules as new subscribers or handlers.

d) Which solution is seemingly more performant?
In ADT using only required classes to process their tasks, so this method could be performing if we required to solve task sequentially, 
while with asynchronous operations this method could be less performant than ED method
For SD if there are a lot of reading load, so this method could be performing, but if it is required to change shared data,
issues with synchronization appear, handling them leading to problems with performance.
Owing to ED design it is easier to make modules with high performance, because listeners process tasks asynchronous
which means that they could handle tasks without waiting when previous event be processed.

e) If you are asked to implement a similar program, which of the solutions would you reuse?
ADT method seems more useful for tasks which could have a good separation of modules with well-defined responsibilities and behaviours.