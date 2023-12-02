import regex

INPUT = 'input_2.txt'
GAMEID = "^Game (\d+):.*$"
MAXRED=12
MAXGREEN=13
MAXBLUE=14

# check if the number of reds, greens, and blues is valid
def check_valid(cubes, max):
    for cube in cubes:
        if int(cube) > max:
            return False
    return True

def get_max(cubes):
    max = 0
    for cube in cubes:
        if int(cube) > max:
            max = int(cube)
    return max

def main():
    sum_step1 = 0
    sum_step2 = 0        
    with open(INPUT, 'r') as f:
        lines = f.readlines()
        for line in lines:
            line = line.strip()
            match = regex.match(GAMEID, line)
            if match:
                gameid = match.group(1)
            else:
                gameid = 0

            reds = regex.findall("(\d+) red", line)                 
            greens = regex.findall("(\d+) green", line)
            blues = regex.findall("(\d+) blue", line)
            # STEP1
            if check_valid(reds, MAXRED) and check_valid(greens, MAXGREEN) and check_valid(blues, MAXBLUE):
                sum_step1 += int(gameid)
            # STEP2
            min_red = get_max(reds)
            min_green = get_max(greens)
            min_blue = get_max(blues)
            power = min_red * min_green * min_blue
            sum_step2 += power

    print("Sum Step1:", sum_step1)
    print("Sum Step2:", sum_step2)

main()