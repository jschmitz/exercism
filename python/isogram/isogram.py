def is_isogram(string):
    string = string.upper().replace('-', '').replace(' ', '')
    return len(string) == len(set(string))
