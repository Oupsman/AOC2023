import math
INPUT = 'input_4.txt'


def part1():
    sum_part1 = 0
    cards = []
    with open(INPUT, 'r') as f:
        lines = f.readlines()
        for line in lines:
            counter = -1
            items = line.strip().split()[2:]
            i = items.index("|")
            winnings = items[:i]
            numbers = items[i+1:]
            for num in numbers:
                if num in winnings:
                    counter += 1
            if counter >= 0:
                result = int(math.pow(2, counter))
                sum_part1 += result
                cards.append(counter+1)
            else:
                cards.append(0)

        num_cards = [1 for _ in cards]
        for i, matches in enumerate(cards):
            for _ in range(num_cards[i]):
                for k in range(matches):
                    num_cards[i + 1 + k] += 1

    print(sum_part1)
    print(sum(num_cards))


part1()
