with open('input_3.txt') as f:
    lines = [line.rstrip() for line in f]


def get_adjacents(number, location):
    nx, ny = location
    num_len = len(str(number))

    adjacents = [
        (nx - 1, ny - 1),
        (nx - 1, ny),
        (nx - 1, ny + 1),
        (nx + num_len, ny - 1),
        (nx + num_len, ny),
        (nx + num_len, ny + 1)
    ]

    for m in range(num_len):
        # Add the adjacents "above" and "below"
        adjacents.append((nx + m, ny - 1))
        adjacents.append((nx + m, ny + 1))

    return adjacents


part_numbers = dict()
symbols = dict()
part_number_sum = 0
gears = dict()  # For part 2
gear_ratio_sum = 0

y_coord = 0

for this_line in lines:
    # We will build the numbers digit by digit as a string
    current_number = ''
    x_coord = 0
    for this_char in this_line:
        if this_char.isdigit():
            current_number += this_char
        else:
            if current_number.isdigit():
                part_numbers[(x_coord - len(current_number), y_coord)] = int(current_number)
                current_number = ''  # Clear the current number
            if this_char != '.':
                symbols[(x_coord, y_coord)] = this_char
                if this_char == '*':
                    gears[(x_coord, y_coord)] = {'n': 0, 'g_ratio': 1}
        x_coord += 1
    if current_number.isdigit():
        part_numbers[(x_coord - len(current_number), y_coord)] = int(current_number)
        current_number = ''  # Clear the current number
    y_coord += 1

for part in part_numbers:
    # Get all the coordinates that are adjacent to this number
    part_adj = get_adjacents(part_numbers[part], part)

    # Check each adjacent location to see if a symbol is there
    for adj_loc in part_adj:
        if adj_loc in symbols:
            # Add the part number to the part number sum
            part_number_sum += part_numbers[part]

            # If the symbol is a gear, increment the number of adjacent parts for that gear
            # and increase its gear ratio
            if symbols[adj_loc] == '*':
                gears[adj_loc]['n'] += 1
                gears[adj_loc]['g_ratio'] *= part_numbers[part]
            break   # Break the for loop because we don't want to double-count any symbols

for gear in gears:
    # We are only concerned about gears with only 2 adjacent numbers
    if gears[gear]['n'] == 2:
        gear_ratio_sum += gears[gear]['g_ratio']

print(f'Part 1: {part_number_sum}')
print(f'Part 2: {gear_ratio_sum}')
