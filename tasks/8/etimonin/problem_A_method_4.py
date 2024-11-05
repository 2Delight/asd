"""
Problem A. Key Word in Context (KWIC)
For problem description see https://en.wikipedia.org/wiki/Key_Word_in_Context

Method 4. Implicit invocation (event-driven)
See Garlan and Shaw paper for description of the method.
In general, any asynchronous or event-driven programming framework will help.
A simpler solution could employ the Observer/Listener pattern instead of direct invocations.
"""

import re

# Event class used to contain required information for subscribers to handle
class Event:
    def __init__(self, keyword, found_position, context_line, context_size):
        self.keyword: str = keyword
        self.found_position: int = found_position
        self.context_line: str = context_line
        self.context_size: int = context_size


# Subscriber handle event from indexer
class Subscriber:
    def __init__(self):
        self.index: list[str] = []

    def handle_event(self, event: Event):
        words = event.context_line.split()

        # find positions of context
        start = max(event.found_position - event.context_size, 0)
        end = event.found_position + event.context_size + 1
        context_string = " ".join(words[start:end])
        self.index.append(context_string)

    def get_index(self):
        return sorted(self.index, key=lambda s: s.lower())


def match_keyword(context_word: str):
    # remove special symbols from word to handle words with special symbols like "(KWIC)"
    # make string lower to handle different cases
    return re.sub(r'[^a-zA-Z]', '', context_word).lower()


# Indexer find matches in text and send events to subscribers
class Indexer:
    def __init__(self, text, keyword, context_size):
        self.text: str = text
        self.context_size: int = context_size
        self.keyword: str = keyword.lower()
        self.subscribers: list[Subscriber] = []

    def find_matches(self):
        lines = self.text.split('\n')
        for line in lines:
            words = line.split()
            for i, word in enumerate(words):
                # if match found, send event to subscribers, in line could be several matches
                if match_keyword(word):
                    self.send_event(Event(self.keyword, i, line, context_size))

    def subscribe(self, subscriber: Subscriber):
        self.subscribers.append(subscriber)

    def send_event(self, event: Event):
        for sub in self.subscribers:
            sub.handle_event(event)


text = """
Key Word In Context (KWIC) is the most common format for KWIC concordance lines. 
The term KWIC was coined by Hans Peter Luhn.[1] The system was based on a concept called keyword in titles, 
which was first proposed for Manchester libraries in 1864 by Andrea Crestadoro.[2]

A KWIC index is formed by sorting and aligning the words within an article title 
to allow each word (except the stop words) in titles to be searchable alphabetically in the index.[3] 
It was a useful indexing method for technical manuals before computerized full text search became common.
"""

keyword = "KWIC"
context_size = 3

indexer = Indexer(text, keyword, context_size)
listener = Subscriber()
indexer.subscribe(listener)

indexer.find_matches()

kwic_index = listener.get_index()

for ln in kwic_index:
    print(ln)
