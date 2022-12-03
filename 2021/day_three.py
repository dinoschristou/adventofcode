from util import get_input

def parser(line):
    bits =[]
    for c in line:
        if c != '\n':
            bits.append(int(c))
    return bits

readings = get_input('day3.txt', parser)

def part_a(readings):
    number_of_bits = len(readings[1])
    barrier = len(readings) / 2
    counts = [0]*number_of_bits
    for r in readings:
        for i in range(len(r)):
            counts[i] += r[i]
    gamma_string = ''
    epsilon_string= ''
    for c in counts:
        if c > barrier:
            gamma_string += '1'
            epsilon_string += '0'
        else:
            gamma_string += '0'
            epsilon_string += '1'
    print(int(gamma_string,2)* int(epsilon_string,2))

def subset(readings, pos, oxygen):
    if (pos > len(readings[0])):
        raise ValueError('position out of bounds')
    if len(readings) == 1:
        return readings[0]
    else:
        # just use a comprehension
        sum_first = sum(map(lambda r : r[pos], readings) )

        filter_func = lambda r : r[pos] == 0
        cut_off = len(readings)/2
        if oxygen and sum_first >= (len(readings)/2):
            filter_func = lambda x : x[pos] == 1
        elif oxygen and sum_first 
        return subset(list(filter(filter_func, readings)), pos+1, oxygen)


print(subset(readings, 0, True))


part_a(readings)