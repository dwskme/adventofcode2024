def day2_1():
    count = 0
    with open("day2/input.txt", "r") as file:
        for line in file:
            values = line.split()
            flag = True
            incr, decr = False, False

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

            if flag:
                count += 1
        print("day2.1", count)
