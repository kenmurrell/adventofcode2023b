
def partOne():
    with open("data.txt") as f:
        sum = 0
        for line in f:
            first = 0
            last = 0
            for char in line:
                try:
                    num = int(char)
                    if first == 0:
                        first = num
                    last = num
                except ValueError:
                    pass
            sum += (first*10) + last
        print(sum) 

def partTwo():
    numdict = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "zero": 0
    }
    with open("data.txt") as f:
        sum = 0
        for line in f:
            first = 0
            last = 0
            buffer = ""
            for char in line:
                try:
                    num = int(char)
                    if first == 0:
                        first = num
                    last = num
                except ValueError:
                    buffer+=char
                    for k,v in numdict.items():
                        if buffer.rfind(k) >= 0 and buffer.rfind(k) + len(k) == len(buffer):
                            if first == 0:
                                first = v
                            last = v
                                    
            sum += (first*10) + last
        print(sum) 

partOne()
partTwo()