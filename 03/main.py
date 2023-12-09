import re

def partOne(filename):
    with open(filename) as f:
        sum = 0
        symbols = {"+", "*", "#", "$", "/", "@", "%", "=", "-", "&", ">", "<", "|", "_", "^", ":", ":"}
        valids = set()
        full_file = list()
        for line in f:
            full_file.append(line.strip())
        for row, line in enumerate(full_file): 
            for col, char in enumerate(line):
                if char not in symbols:
                    continue
                # add top
                valids.add((row-1, col))
                # add bottom
                valids.add((row+1, col))
                # add left
                valids.add((row, col-1))
                # add right
                valids.add((row, col+1))
                # add top left
                valids.add((row-1, col-1))
                # add top right
                valids.add((row-1, col+1))
                # add bottom left
                valids.add((row+1, col-1))
                # add bottom right
                valids.add((row+1, col+1))
        for row in range(0, len(full_file)):
            buffer = ""
            add = False
            for col in range(0, len(full_file[0])):
                val = full_file[row][col]
                try:
                   integer = int(val)
                   buffer += val
                   if (row, col) in valids:
                        add = True
                except ValueError:
                    if len(buffer) > 0 and add:
                        sum += int(buffer)
                    buffer = ""
                    add = False
            if len(buffer) > 0 and add:
                sum += int(buffer)
        return sum

def test(ans, file, fn):
    r = fn(file)
    if r != ans:
        raise ValueError("Want %d, got %d.", ans, r)
    else:
        return "SUCCESS"

print("PartOne Test: " + test(4361, "test.txt", partOne))
print("PartOne Full: " + test(539637, "data.txt", partOne))