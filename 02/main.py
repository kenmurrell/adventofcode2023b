import re

def partOne(filename: str) -> str:
    red = re.compile('[0-9]+(?=\sred)')
    blue = re.compile('[0-9]+(?=\sblue)')
    green = re.compile('[0-9]+(?=\sgreen)')
    game = re.compile('(?<=Game\s)[0-9]+')
    with open(filename) as f:
        id_sum = 0
        for line in f: 
            game_id, info = line.split(": ")
            red_sum = 0
            blue_sum = 0
            green_sum = 0
            for m in red.findall(info):
                if(red_sum < int(m)):
                    red_sum = int(m)
            for m in blue.findall(info):
                if(blue_sum < int(m)):
                    blue_sum = int(m)
            for m in green.findall(info):
                if(green_sum < int(m)):
                    green_sum = int(m)
            if red_sum <= 12 and green_sum <= 13 and blue_sum <= 14:
                id_sum += int(game.findall(game_id)[0])
        return id_sum

def partTwo(filename: str) -> str:
    red = re.compile('[0-9]+(?=\sred)')
    blue = re.compile('[0-9]+(?=\sblue)')
    green = re.compile('[0-9]+(?=\sgreen)')
    game = re.compile('(?<=Game\s)[0-9]+')
    with open(filename) as f:
        id_sum = 0
        for line in f: 
            game_id, info = line.split(": ")
            red_min = 0
            green_min = 0
            blue_min = 0
            for m in red.findall(info):
                if red_min < int(m):
                    red_min = int(m)
            for m in blue.findall(info):
                if blue_min < int(m):
                    blue_min = int(m)
            for m in green.findall(info):
                if green_min < int(m):
                    green_min = int(m)
            id_sum += (red_min*blue_min*green_min)
        return id_sum

def test(ans, file, fn):
    r = fn(file)
    if r != ans:
        raise ValueError("Want %d, got %d.", ans, r)
    else:
        return "SUCCESS"

print("PartOne Test: " + test(8, "test.txt", partOne))
print("PartOne Full: " + test(2771, "data.txt", partOne))
print("PartTwo Test: " + test(2286, "test.txt", partTwo))
print("PartTwo Full: " + test(70924, "data.txt", partTwo))