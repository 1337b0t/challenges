import sys

def check_string():
    S = '3'

    try:
        return int(S)
    except:
        return 'Bad String'

if __name__ == "__main__":
    print(check_string())
    