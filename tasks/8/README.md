
### Comparsion

| Method                                | Change algorithm | Change data representation | Add functionality | Seem more performance | Ease of reuse |
|---------------------------------------|------------------|----------------------------|-------------------|-----------------------|---------------|
| Abstract Data Types (1)               | -                | +                          | -                 | -                     | +             |
| Shared Data (2)                       | -                | -                          | -                 | +                     | -             |
| Pipes-and-filters (3)                 | +                | -                          | +                 | -                     | -             |
| Implicit invocation, event-driven (4) | +                | -                          | +                 | +                     | -             |

**a) In which case it is easier to change the implementation algorithm in each of the modules?**

Abstract Data Types (ADT) and Shared Data (SD) methods have more interconnected components, so changing algorithm would affect several modules.

Pipes-and-filters (PF) and Implicit invocation, event-driven (ED) methods have less interconnected components, using modularity, so changes in one component would not affect other.

**b) In which solution it is easier (= seemingly less effort) to change data representation**

ADT method has separation of concerns, so changing in one module would not affect other.

SD has one source of data, so changing representation for one component would affect other.

PF and ED methods use specific contracts for communication, so changing data representation would affect filters and subscribers.


**c) In which solution it is easier to add additional functions to the modules**

ADT and SD methods adding new functionality would affect several classes or subroutines, because components are more interconnected, using modularity.

PF and ED have less interconnected modules, so it would less challenging to add new functionality.


**d) Which solution is seemingly more performant?**

ED would have high performance for parallel independent tasks.

SD would have high performance for sequence and parallel tasks with reading load, if tasks required to modify shared data, performance would be lower.

ADT and PF would have high performance if sequence execution required, for parallel execution these methods would not have high performance.


**e) If you are asked to implement a similar program, which of the solutions would you reuse?**

ADT method typically has atomic modules, so it would be easy to reuse code which used this method. 

PF and ED methods code could be reusable, but it is required to adapt several modules.

SD method is less modular, so it would be challenging to reuse this code.