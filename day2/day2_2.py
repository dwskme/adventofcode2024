def day2_2():
    count = 0
    with open("day2/input.txt", "r") as file:
        for line in file:
            values = line.split()
            if with_dampener(values):
                count += 1

    print("day2.2", count)


def is_safe(values):
    incr, decr = False, False
    flag = True
    for i in range(1, len(values)):
        left = int(values[i - 1])
        right = int(values[i])
        if left < right:
            incr = True
        elif left > right:
            decr = True
        if incr and decr:
            flag = False
            break
        if abs(left - right) < 1 or abs(left - right) > 3:
            flag = False
            break
    return flag


def with_dampener(values):
    if is_safe(values):
        return True

    for i in range(len(values)):
        new_values = values[:i] + values[i + 1 :]
        if is_safe(new_values):
            return True

    return False
