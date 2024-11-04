from collections import deque

# Приём входных данных
class Input:
    def __init__(self, filePath):
        try:
            with open(filePath, 'r') as file:
                self.lines = [line.strip() for line in file.readlines() if line.strip()]

            if not self.lines:
                raise ValueError("[DBG] Empty file")
        except FileNotFoundError:
            print(f"[ERROR] Not found {filePath} file")
            self.lines = []
        except ValueError as e:
            self.lines = []

    def getLines(self):
        return self.lines

# Сущность для сдвигов
class Shifter:
    def circularShift(self, lines):
        shifted_lines = []
        for line in lines:
            words = line.split(" ")
            dec = deque(words)
            for _ in range(len(dec)):
                dec.rotate()
                shifted_lines.append(" ".join(dec))
        return shifted_lines

# Сортировка строк
class Sort:
    def sort(self, lines):
        return sorted(lines)

# Вывод данных
class Output:
    def __init__(self, outputFilePath):
        self.outputFilePath = outputFilePath

    def display(self, lines):
        with open(self.outputFilePath, 'w') as file:
            for line in lines:
                file.write(line + "\n")

        print(f"[DEBUG] Results written to {self.outputFilePath} file")

class KWICAlgorithm:
    def __init__(self, filePath, outputFilePath):
        self.inputFilter = Input(filePath)
        self.circularShiftFilter = Shifter()
        self.sortFilter = Sort()
        self.outputFilter = Output(outputFilePath)

    def execute(self):
        lines = self.inputFilter.getLines()

        if not lines:
            print("[ERROR] Empty file")
            return

        shiftedLines = self.circularShiftFilter.circularShift(lines)
        sortedLines = self.sortFilter.sort(shiftedLines)

        self.outputFilter.display(sortedLines)

filePath = 'input.txt'
outputFilePath = 'output.txt'

kwicSystem = KWICAlgorithm(filePath, outputFilePath)
kwicSystem.execute()


