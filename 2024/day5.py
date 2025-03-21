from collections import defaultdict

page_ordering_rules: dict[int, set[int]] = defaultdict(set)
page_updates: list[list[int]] = []

with open("day5.txt") as f:
    # Read first part of puzzle input
    while line := f.readline().strip():
        p, q = line.split("|")
        page_ordering_rules[int(p)].add(int(q))

    # Read second part of puzzle input
    while line := f.readline().strip():
        page_updates.append(list(map(int, line.split(","))))


# Parts 1 and 2


def is_valid_update(page_update: list[int]) -> bool:
    n = len(page_update)

    return all(
        page_update[j] in page_ordering_rules[page_update[i]]
        for i in range(n - 1)
        for j in range(i + 1, n)
    )


def fix_page_update(page_update: list[int]) -> None:
    if is_valid_update(page_update=page_update):
        return

    n = len(page_update)

    for i in range(n - 1):
        for j in range(i + 1, n):
            if page_update[i] in page_ordering_rules[page_update[j]]:
                page_update[i], page_update[j] = page_update[j], page_update[i]
                fix_page_update(page_update=page_update)
                return


part1_result, part2_result = 0, 0

for page_update in page_updates:
    if is_valid_update(page_update=page_update):
        part1_result += page_update[len(page_update) // 2]
    else:
        fix_page_update(page_update=page_update)
        part2_result += page_update[len(page_update) // 2]

print("Part 1:", part1_result)
print("Part 2:", part2_result)
