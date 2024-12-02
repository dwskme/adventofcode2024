def day1_2():
    left = []
    right = []
    result = []

    with open("day1/input.txt", "r") as file:
        for line in file:
            values = line.split()
            if len(values) == 2:
                left.append(int(values[0]))
                right.append(int(values[1]))

    for l in left:
        count = 0
        for r in right:
            if l == r:
                count += 1
        result.append(l * count)

    print(sum(result))
