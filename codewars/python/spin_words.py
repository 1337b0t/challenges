"""
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

Examples:

spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw" spinWords( "This is a test") => returns "This is a test" spinWords( "This is another test" )=> returns "This is rehtona test"
"""

def spin_words(sentence):
    
    # split the inputs into a list of words
    words = list(map(str, str(sentence).strip().split(' ')))
    
    # exit if their is no input
    if len(sentence) is 0:
        return None
    
    # reverse the value of words with lengths greater than or equal to 5
    for idx, val in enumerate(words):
        if len(val) >= 5:
            words[idx] =  words[idx][::-1]
    
    # return the new word list with a space between words
    return " ".join(words[:])