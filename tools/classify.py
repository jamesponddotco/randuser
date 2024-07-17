"""
This script classifies words from a word list into three categories: nouns,
adjectives, and verbs. The categories are written to separate files.

The script was written for Python 3.11 with Claude Sonnet 3.5 as the actual
writer, as I've no Python knowledge and didn't feel like writing this in Go.

It does the job though, so ¯\_(ツ)_/¯
"""

import sys
import logging
from typing import List, Dict, Callable
from collections import defaultdict
import nltk
from nltk.corpus import wordnet

WORDNET_RESOURCE = 'wordnet'
OUTPUT_FILES = {
    'n': 'nouns.txt',
    'a': 'adjectives.txt',
    'v': 'verbs.txt',
    'other': 'others.txt'
}

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def download_nltk_resource(resource: str) -> None:
    """Download NLTK resource if not already present."""
    try:
        nltk.data.find(f'corpora/{resource}')
    except LookupError:
        logging.info(f"Downloading NLTK resource: {resource}")
        nltk.download(resource, quiet=True)

def get_word_type(word: str) -> str:
    """Determine the type of a word based on its synsets."""
    synsets = wordnet.synsets(word)
    if any(synset.pos() == 'n' for synset in synsets):
        return 'n'
    elif any(synset.pos() == 'a' for synset in synsets):
        return 'a'
    elif any(synset.pos() == 'v' for synset in synsets):
        return 'v'
    return 'other'

def categorize_words(words: List[str]) -> Dict[str, List[str]]:
    """Categorize words into nouns, adjectives, verbs, and others."""
    categories: Dict[str, List[str]] = defaultdict(list)
    for word in words:
        word_type = get_word_type(word)
        categories[word_type].append(word)
    return categories

def write_words_to_file(words: List[str], filename: str) -> None:
    """Write a list of words to a file."""
    try:
        with open(filename, 'w') as file:
            file.write('\n'.join(words))
        logging.info(f"Successfully wrote {len(words)} words to {filename}")
    except IOError as e:
        logging.error(f"Error writing to {filename}: {e}")

def process_wordlist(wordlist_file: str) -> None:
    """Process the wordlist file and categorize words."""
    try:
        with open(wordlist_file, 'r') as file:
            words: List[str] = file.read().splitlines()

        logging.info(f"Read {len(words)} words from {wordlist_file}")

        categories = categorize_words(words)

        for pos, filename in OUTPUT_FILES.items():
            write_words_to_file(categories[pos], filename)

    except IOError as e:
        logging.error(f"Error reading {wordlist_file}: {e}")
        sys.exit(1)

def main() -> None:
    """Main function to process command-line arguments and initiate word processing."""
    if len(sys.argv) != 2:
        logging.error("Usage: python classify.py <wordlist_file>")
        sys.exit(1)

    wordlist_file = sys.argv[1]
    download_nltk_resource(WORDNET_RESOURCE)
    process_wordlist(wordlist_file)

if __name__ == '__main__':
    main()
