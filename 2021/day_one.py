
from util import get_input

readings = get_input('day1.txt', int)

def part_a(readings):
    increments = 0
    for i in range(1, len(readings)):
        if (readings[i] > readings[i-1]):
            increments += 1
    return increments

def part_b(readings):
    increments = 0
    for i in range (3,len(readings)):
        if sum(readings[i-3:i]) < sum(readings[i-2:i+1]):
            increments += 1
    return increments

print(part_a(readings))
print(part_b(readings))