from util import get_input

def parser(line):
    instructions = line.split(' ')
    return (instructions[0], int(instructions[1]))

instructions = get_input('day2.txt', parser)

def part_a(instructions):
    horizontal = 0
    depth = 0
    for instruction, value  in instructions:
        if instruction == 'forward':
            horizontal += value
        elif instruction == 'down':
            depth += value
        elif instruction == 'up':
            depth -= value
    return horizontal*depth

def part_b(instructions):
    horizontal = 0
    depth = 0
    aim = 0
    for instruction, value  in instructions:
        if instruction == 'forward':
            horizontal += value
            depth += (aim*value)
        elif instruction == 'down':
            aim += value
        elif instruction == 'up':
            aim -= value
    return horizontal*depth

print(part_a(instructions))
print(part_b(instructions))

