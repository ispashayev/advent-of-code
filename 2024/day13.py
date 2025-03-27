import numpy as np
import re


button_a = re.compile(r"Button A: X\+(?P<x>\d+), Y\+(?P<y>\d+)")
button_b = re.compile(r"Button B: X\+(?P<x>\d+), Y\+(?P<y>\d+)")
prize = re.compile(r"Prize: X=(?P<x>\d+), Y=(?P<y>\d+)")

per_token_cost = np.array([3, 1])

button_systems = []
prizes = []

with open("day13.txt") as f:
    puzzle_input = f.read()
    claw_machine_blobs = puzzle_input.split("\n\n")
    for claw_machine_blob in claw_machine_blobs:
        button_a_def = button_a.search(claw_machine_blob)
        button_b_def = button_b.search(claw_machine_blob)
        prize_def = prize.search(claw_machine_blob)

        x_a, y_a = int(button_a_def["x"]), int(button_a_def["y"])
        x_b, y_b = int(button_b_def["x"]), int(button_b_def["y"])
        prize_x, prize_y = int(prize_def["x"]), int(prize_def["y"])

        button_systems.append(np.array([[x_a, x_b], [y_a, y_b]]))
        prizes.append(np.array([prize_x, prize_y]))


# Parts 1 and 2


def get_token_cost(button_system, prize_vector) -> int:
    if abs(np.linalg.det(button_system)) < 1e-6:
        return 0

    inv_button_system = np.linalg.inv(button_system)
    button_presses = np.round(np.dot(inv_button_system, prize_vector)).astype(int)

    if np.any(np.dot(button_system, button_presses) != prize_vector):
        return 0

    return np.dot(button_presses, per_token_cost)


part1_token_cost = part2_token_cost = 0

for button_system, prize in zip(button_systems, prizes, strict=True):
    part1_token_cost += get_token_cost(button_system=button_system, prize_vector=prize)
    part2_token_cost += get_token_cost(
        button_system=button_system, prize_vector=prize + 10000000000000
    )


print("Part 1:", part1_token_cost)
print("Part 2:", part2_token_cost)
