MAX_HEIGHT = 9

topographic_map: list[list[int]] = []

with open("day10.txt") as f:
    for line in f.readlines():
        topographic_map.append(list(map(int, line.strip())))

n, m = len(topographic_map), len(topographic_map[0])

trailheads = [(i, j) for i in range(n) for j in range(m) if topographic_map[i][j] == 0]


def climb(position: tuple[int, int], height: int) -> list[tuple[int, int]]:
    if height == MAX_HEIGHT:
        return []

    x, y = position
    return [
        (i, j)
        for i, j in (
            (x - 1, y),
            (x + 1, y),
            (x, y - 1),
            (x, y + 1),
        )
        if 0 <= i < n and 0 <= j < m and topographic_map[i][j] == height + 1
    ]


# Part 1


def get_trail_score_by_distinct_peaks(trailhead: tuple[int, int]) -> int:
    peaks: set[tuple[int, int]] = set()

    positions: list[tuple[int, int]] = [trailhead]
    while positions:
        position = positions.pop(0)

        x, y = position
        height = topographic_map[x][y]

        if height == MAX_HEIGHT:
            peaks.add(position)

        positions.extend(climb(position=position, height=height))

    return len(peaks)


trailhead_score_sum = sum(
    get_trail_score_by_distinct_peaks(trailhead=trailhead) for trailhead in trailheads
)

print("Part 1:", trailhead_score_sum)


# Part 2


def get_trail_score_by_distinct_paths(trailhead: tuple[int, int]) -> int:
    trail_score = 0

    def _explore_trail(position: tuple[int, int]):
        nonlocal trail_score

        x, y = position
        height = topographic_map[x][y]

        if height == MAX_HEIGHT:
            trail_score += 1

        for new_position in climb(position=position, height=height):
            _explore_trail(position=new_position)

    _explore_trail(position=trailhead)

    return trail_score


trailhead_score_sum = sum(
    get_trail_score_by_distinct_paths(trailhead=trailhead) for trailhead in trailheads
)

print("Part 2:", trailhead_score_sum)
