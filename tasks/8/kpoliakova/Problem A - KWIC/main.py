from typing import List
from collections import namedtuple

WordStructure = namedtuple('WordStructure', ['Word', 'Index'])

class KWIC:
    def __init__(self, text, signs=(' ', ',', '.', '\n')):
        self.signs = signs
        self.old_text = text
        self.left_padding = 35
        self.right_padding = 35
        self.list_of_words = self.divide_words(text)
        self._word_set_list = [WordStructure(word, index) for index, word in enumerate(self.list_of_words)]

    def divide_words(self, text):
        default_sign = self.signs[0]

        for sign in self.signs[1:]:
            text = text.replace(sign, default_sign)

        word_list = [i.strip() for i in text.split(default_sign)]

        return [word for word in word_list if word]

    def find(self, target):
        searched_words = [word_set for word_set in self._word_set_list if word_set.Word == target]
        if not searched_words:
            return None
        return [self.print_ready_line(word_set) for word_set in searched_words]

    def print_ready_line(self, word_set):
        forward_words = self.forward_words(word_set, self.left_padding)
        backward_words = self.back_words(word_set, self.right_padding)
        return f"{forward_words}{word_set.Word}{backward_words}"

    def forward_words(self, word_structure, left_padding):
        char_count = 0
        forward_list = []
        for idx in range(word_structure.Index - 1, -1, -1):
            result, char_count = self.make_necessary_strings(idx, left_padding, forward_list, char_count)
            if result:
                break
        forward_list.reverse()
        return " " * (left_padding - char_count) + " ".join(forward_list) + " "

    def back_words(self, word_structure, right_padding):
        char_count = 0
        backward_list = []
        for idx in range(word_structure.Index + 1, len(self.list_of_words)):
            result, char_count = self.make_necessary_strings(idx, right_padding, backward_list, char_count)
            if result:
                break
        return " " + " ".join(backward_list) + " " * (right_padding - char_count - 1)

    def make_necessary_strings(self, idx, limit, list_with_words, char_count):
        word = self.list_of_words[idx]
        count = char_count + len(word) + 1
        if count > limit:
            return True, char_count
        char_count = count
        list_with_words.append(word)
        return False, char_count

if __name__ == "__main__":
    text = "For example, a search query including all of the words in an example definition (KWIC is an acronym for Key Word In Context," + \
    " the most common format for concordance lines) and the Wikipedia slogan in English (the free encyclopedia), searched against a Wikipedia " + \
    "page, might yield a KWIC index as follows. A KWIC index usually uses a wide layout to allow the display of maximum 'in context' " + \
    "information (not shown in the following example)."

    kwic = KWIC(text)

    results = kwic.find("KWIC")

    if results:
        for result in results:
            print(result)
