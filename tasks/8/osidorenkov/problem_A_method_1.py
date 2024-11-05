class LineStorage:
    def __init__(self):
        self.lines = []

    def add_line(self, line):
        self.lines.append(line)

    def get_lines(self):
        return self.lines


class CircularShifter:
    def __init__(self, line_storage):
        self.line_storage = line_storage
        self.shifts = []

    def generate_shifts(self):
        for line in self.line_storage.get_lines():
            words = line.split()
            for i in range(len(words)):
                shifted_line = ' '.join(words[i:] + words[:i])
                self.shifts.append(shifted_line)

    def get_shifts(self):
        return self.shifts


class Alphabetizer:
    def __init__(self, shifter):
        self.shifter = shifter
        self.sorted_shifts = []

    def sort_shifts(self):
        self.sorted_shifts = sorted(self.shifter.get_shifts())

    def get_sorted_shifts(self):
        return self.sorted_shifts


class KWICSystem:
    def __init__(self):
        self.line_storage = LineStorage()
        self.shifter = CircularShifter(self.line_storage)
        self.alphabetizer = Alphabetizer(self.shifter)

    def add_line(self, line):
        self.line_storage.add_line(line)

    def generate_kwic_index(self):
        self.shifter.generate_shifts()
        self.alphabetizer.sort_shifts()
        return self.alphabetizer.get_sorted_shifts()


if __name__ == "__main__":
    kwic = KWICSystem()
    while True:
        line = input()
        if line == '':
            break
        kwic.add_line(line)
    index = kwic.generate_kwic_index()
    print(*index, sep='\n')
