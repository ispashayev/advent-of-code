word_search = []

with open("day4.txt") as f:
    for line in f.readlines():
        word_search.append(list(line))


# Part 1

n, m = len(word_search), len(word_search[0])

xmas_count, stack = 0, ""


def search_xmas(i: int, j: int):
    global xmas_count

    for v1 in range(-1, 2):
        for v2 in range(-1, 2):
            if v1 == 0 and v2 == 0:
                continue

            if not (0 <= i + 3 * v1 < n and 0 <= j + 3 * v2 < m):
                continue

            word = "".join(word_search[i + k * v1][j + k * v2] for k in range(4))
            if word == "XMAS":
                xmas_count += 1


for i in range(n):
    for j in range(m):
        search_xmas(i, j)

print("Part 1:", xmas_count)


# Part 2

x_mas_count = 0


def search_x_mas(i: int, j: int):
    global x_mas_count

    word1 = "".join(word_search[i + offset][j + offset] for offset in range(-1, 2))
    word2 = "".join(word_search[i + offset][j - offset] for offset in range(-1, 2))

    if (word1 == "MAS" or word1 == "SAM") and (word2 == "MAS" or word2 == "SAM"):
        x_mas_count += 1


for i in range(1, n - 1):
    for j in range(1, m - 1):
        search_x_mas(i, j)

print("Part 2:", x_mas_count)
