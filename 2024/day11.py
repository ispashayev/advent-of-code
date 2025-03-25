from collections import defaultdict


with open("day11.txt") as f:
    stones = list(map(int, f.read().split()))


# Part 1


def run_naive_blink_simulation(stones: list[int], num_blinks: int) -> int:
    stones = stones[:]

    for i in range(num_blinks):
        j = 0
        while j < len(stones):
            if stones[j] == 0:
                stones[j] = 1

            elif len(str(stones[j])) % 2 == 0:
                stone_str = str(stones[j])
                n = len(stone_str)
                stones[j] = int(stone_str[: n // 2])
                stones.insert(j + 1, int(stone_str[n // 2 :]))
                j += 1

            else:
                stones[j] *= 2024

            j += 1

    return len(stones)


print("Part 1:", run_naive_blink_simulation(stones=stones, num_blinks=25))


# Part 2


def run_blink_simulation(stones: list[int], num_blinks: int) -> int:
    stone_distribution: dict[int, int] = defaultdict(int)

    for stone in stones:
        stone_distribution[stone] += 1

    for _ in range(num_blinks):
        present_distribution_state = list(stone_distribution.items())
        for stone, count in present_distribution_state:
            if stone == 0:
                stone_distribution[1] += count
            elif len(stone_str := str(stone)) % 2 == 0:
                mid = len(stone_str) // 2
                stone_distribution[int(stone_str[:mid])] += count
                stone_distribution[int(stone_str[mid:])] += count
            else:
                stone_distribution[2024 * stone] += count

            stone_distribution[stone] -= count

    return sum(stone_distribution.values())


print("Part 2:", run_blink_simulation(stones=stones, num_blinks=75))
