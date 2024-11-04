import sys

class QueensAlgorithm:
    def __init__(self, count):
        self.boardSize = count
        self.board = []
        self.queenPositions = []
        self.initializeBoard()

    def run(self):
        row = 0
        col = 0

        while row < self.boardSize:
            col = self.findSafePosition(row, col)

            if col < self.boardSize:
                self.placeQueen(row, col)
                row += 1
                col = 0
            else:
                if self.queenPositions:
                    lastPosition = self.queenPositions.pop()
                    self.removeQueen(lastPosition['row'], lastPosition['col'])
                    row = lastPosition['row']
                    col = lastPosition['col'] + 1
                else:
                    print("[Fail] No solutions")
                    return

        self.printSolution()

    def initializeBoard(self):
        self.board = [[0 for _ in range(self.boardSize)] for _ in range(self.boardSize)]

    def findSafePosition(self, row, startCol):
        for col in range(startCol, self.boardSize):
            if self.isSafe(row, col):
                return col
        return self.boardSize

    def placeQueen(self, row, col):
        self.board[row][col] = 1
        self.queenPositions.append({'row': row, 'col': col})

    def removeQueen(self, row, col):
        self.board[row][col] = 0

    def isSafe(self, row, col):
        # Столбцы
        for i in range(row):
            if self.board[i][col] == 1:
                return False

        # Диагонали
        i, j = row, col
        while i >= 0 and j >= 0:
            if self.board[i][j] == 1:
                return False
            i -= 1
            j -= 1

        i, j = row, col
        while i >= 0 and j < self.boardSize:
            if self.board[i][j] == 1:
                return False
            i -= 1
            j += 1

        return True

    def printSolution(self):
        for row in self.board:
            for col in row:
                print("q" if col == 1 else "·", end=" ")
            print()

numOfQueens = int(sys.argv[1])
queensAlgo = QueensAlgorithm(numOfQueens)
queensAlgo.run()

