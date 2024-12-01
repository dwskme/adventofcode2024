left = []
right = []
sum = 0


with open("input.txt", "r") as file:
    for line in file:
        values = line.split()
        if len(values) == 2:
            left.append(int(values[0]))
            right.append(int(values[1]))


left.sort()
right.sort()

# print(len(left) == len(right))


for l, r in zip(left, right):
    sum += abs(l - r)

print(sum)
