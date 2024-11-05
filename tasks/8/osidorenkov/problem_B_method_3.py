from itertools import permutations
from typing import List


def output(perm: List[List[str]]):
    size = len(perm)
    board = []
    for i in range(size):
        row = ['.'] * size
        row[perm[i]] = 'Q'
        board.append(' '.join(row))
    return '\n'.join(board)


def is_valid(perm: List[List[str]]):
    for i in range(len(perm)):
        for j in range(i + 1, len(perm)):
            if abs(perm[i] - perm[j]) == abs(i - j):
                return False
    return True


def solve_n_queens(n: int):
    columns = range(n)
    solutions = []
    for perm in permutations(columns):
        if is_valid(perm):
            solutions.append(perm)
    return solutions

def main():
    n = 8
    solutions = solve_n_queens(n)
    print(f"{len(solutions)} solutions for {n} queens\n")
    for index, solution in enumerate(solutions, start=1):
        print(f"Solution {index}:\n{output(solution)}\n")


if __name__ == '__main__':
    main()
