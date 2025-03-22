from collections import defaultdict
from dataclasses import dataclass


patrol_map: list[list[str]] = []

with open("day6.txt") as f:
    for line in f.readlines():
        patrol_map.append(list(line.strip()))


# Part 1

unit_vectors = {
    "^": (-1, 0),
    ">": (0, 1),
    "<": (0, -1),
    "v": (1, 0),
}

unit_vector_transition_map = {
    "^": ">",
    ">": "v",
    "v": "<",
    "<": "^",
}


n, m = len(patrol_map), len(patrol_map[0])

start_coordinates, start_heading = next(
    ((i, j), patrol_map[i][j])
    for i in range(n)
    for j in range(m)
    if patrol_map[i][j] in unit_vectors
)


def apply_unit_vector(i: int, j: int, heading: str) -> tuple[int, int]:
    unit_vector = unit_vectors[heading]
    return i + unit_vector[0], j + unit_vector[1]


@dataclass
class WalkResult:
    positions_visited: list[tuple[int, int]]
    is_loop: bool


def walk(start_coordinates: tuple[int, int], heading: str) -> WalkResult:
    # Map from positions to all of the observed headings at that position
    states_visited: dict[tuple[int, int], set[str]] = defaultdict(set)

    i, j = start_coordinates
    is_loop = False

    while 0 <= i < n and 0 <= j < m:
        if (i, j) in states_visited and heading in states_visited[(i, j)]:
            is_loop = True
            break

        states_visited[(i, j)].add(heading)

        next_i, next_j = apply_unit_vector(i=i, j=j, heading=heading)

        if not (0 <= next_i < n and 0 <= next_j < m):
            # Guard has walked off the map
            break

        while patrol_map[next_i][next_j] == "#":
            # Next space is an obstacle, so guard turns right instead. It's possible
            # that there is an obstacle to the right as well, in which case the guard
            # ends up turning backwards.
            heading = unit_vector_transition_map[heading]
            next_i, next_j = apply_unit_vector(i=i, j=j, heading=heading)

        i, j = next_i, next_j

    return WalkResult(
        positions_visited=list(states_visited.keys()),
        is_loop=is_loop,
    )


walk_result = walk(start_coordinates=start_coordinates, heading=start_heading)

print("Part 1:", len(walk_result.positions_visited))


# Part 2

i, j = start_coordinates
heading = start_heading

obstacle_position_results: dict[tuple[int, int], bool] = {}

for position in walk_result.positions_visited:
    if position == start_coordinates or position in obstacle_position_results:
        continue

    i, j = position
    patrol_map[i][j] = "#"

    obstructed_walk_result = walk(
        start_coordinates=start_coordinates, heading=start_heading
    )
    obstacle_position_results[position] = obstructed_walk_result.is_loop

    patrol_map[i][j] = "."


print("Part 2:", sum(obstacle_position_results.values()))
