def get_input(filepath, parse):
    readings = []

    with open(filepath, 'r') as f:
        for line in f:
            readings.append(parse(line))

    return readings