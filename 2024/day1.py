with open('day1.txt') as f:
    l1, l2 = [], []
    for line in f.readlines():
        l1_i, l2_i = line.split()
        l1.append(int(l1_i))
        l2.append(int(l2_i))

l1.sort()
l2.sort()

total_distance = sum(abs(l1_i - l2_i) for l1_i, l2_i in zip(l1, l2))

print(total_distance)
