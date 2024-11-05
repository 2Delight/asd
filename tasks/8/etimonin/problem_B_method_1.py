"""
Problem B. Eight Queens (8Q)
For problem description see https://en.wikipedia.org/wiki/Eight_queens_puzzle

Method 1. Abstract Data Types
This is similar to object-oriented design method, use any of the approaches to identify classes and provide rationale for eliciting them.
"""


# Queen class that contains information about queen placement
class Queen:
    def __init__(self, x: int, y: int):
        self.x: int = x
        self.y: int = y

    def can_not_be_placed(self, x: int, y: int):
        # check that queen doesn't have the same x or y coordinates, check that coordinates lie on the diagonal
        return self.x == x or self.y == y or self.x - x == self.y - y or self.x - x == y - self.y


# QueenProblemResolver find solutions recursively
class QueenProblemResolver:
    def __init__(self, size: int):
        self.size: int = size
        self.queens: list[Queen] = []
        self.solutions: list[[Queen]] = []

    def is_available_position(self, x: int, y: int):
        for q in self.queens:
            if q.can_not_be_placed(x, y):
                return False

        return True

    def remove_last_queen(self):
        if self.queens:
            self.queens.pop()

    def solve_all(self, x=0):
        if x >= self.size:
            # if all queens are placed, we add solution
            self.solutions.append([Queen(q.x, q.y) for q in self.queens])
            return

        for row in range(self.size):
            if self.is_available_position(row, x):
                self.queens.append(Queen(row, x))
                self.solve_all(x + 1)

                # remove last added queen to find another solution
                self.remove_last_queen()

    def display_solution(self, solution: list[Queen]):
        board = [["_" for _ in range(self.size)] for _ in range(self.size)]
        for queen in solution:
            board[queen.x][queen.y] = "Q"
        for row in board:
            print(" ".join(row))
        print("\n")


resolver = QueenProblemResolver(8)
resolver.solve_all()

print("solutions count:", len(resolver.solutions))
resolver.display_solution(resolver.solutions[0])
