from collections import defaultdict
from itertools import combinations


antenna_map: list[list[str]] = []


with open("day8.txt") as f:
    for line in f.readlines():
        antenna_map.append(list(line.strip()))


# Part 1

n, m = len(antenna_map), len(antenna_map[0])

antenna_location_groups: dict[str, list[tuple[int, int]]] = defaultdict(list)

for i in range(n):
    for j in range(m):
        if antenna_map[i][j] != ".":
            antenna_location_groups[antenna_map[i][j]].append((i, j))


def compute_antinodes(
    c1: tuple[int, int], c2: tuple[int, int]
) -> list[tuple[int, int]]:
    i1, j1 = c1
    i2, j2 = c2

    # Components of a transition vector from c2 to c1
    col_distance, row_distance = j1 - j2, i1 - i2

    # To get c2's antinode, apply the computed transition vector to c1.
    # To get c1's antinode, apply the negated transition vector to c2.
    antinode_candidates: list[tuple[int, int]] = [
        (i1 + row_distance, j1 + col_distance),
        (i2 - row_distance, j2 - col_distance),
    ]

    return [(x, y) for x, y in antinode_candidates if 0 <= x < n and 0 <= y < m]


antinode_locations: set[tuple[int, int]] = set()

for antenna, locations in antenna_location_groups.items():
    for c1, c2 in combinations(iterable=locations, r=2):
        antinodes = compute_antinodes(c1=c1, c2=c2)
        antinode_locations.update(antinodes)

print("Part 1:", len(antinode_locations))


# Part 2


def compute_resonant_antinodes(
    c1: tuple[int, int], c2: tuple[int, int]
) -> list[tuple[int, int]]:
    i1, j1 = c1
    i2, j2 = c2

    col_distance, row_distance = j1 - j2, i1 - i2

    resonant_antinodes: list[tuple[int, int]] = []

    # Compute resonant antinodes for c1
    x, y = c2
    while 0 <= x < n and 0 <= y < m:
        resonant_antinodes.append((x, y))
        x += row_distance
        y += col_distance

    # Compute resonant antinodes for c2
    x, y = c1
    while 0 <= x < n and 0 <= y < m:
        resonant_antinodes.append((x, y))
        x -= row_distance
        y -= col_distance

    return resonant_antinodes


resonating_antinode_locations: set[tuple[int, int]] = set()

for antenna, locations in antenna_location_groups.items():
    for c1, c2 in combinations(iterable=locations, r=2):
        resonating_antinodes = compute_resonant_antinodes(c1=c1, c2=c2)
        resonating_antinode_locations.update(resonating_antinodes)


print("Part 2:", len(resonating_antinode_locations))
