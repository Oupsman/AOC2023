
import regex

spelled_out_numbers = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
INPUT = 'input_1.txt'

def replace(match):
#   print("Match:", match)
    if match in ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]:
        return match
    return str(spelled_out_numbers.index(match))


def main():
    sum = 0

    regexp = "(zero|one|two|three|four|five|six|seven|eight|nine|0|1|2|3|4|5|6|7|8|9)"
        
    with open(INPUT, 'r') as f:
        lines = f.readlines()
        for line in lines:
            res_str = ""
            # strip the newline character
            line = line.strip()    
            matches = regex.findall(regexp, line, overlapped=True)
            for match in matches:
                res_str += replace(match)

            hidden_number = int(res_str[0]) * 10 + int(res_str[len(res_str)-1])
            sum += hidden_number
            
    print("Sum:", sum)

main()
