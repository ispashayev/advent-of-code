from typing import Callable


calibration_equations: dict[int, list[int]] = {}

with open("day7.txt") as f:
    for line in f.readlines():
        test_value, input_values = line.split(": ")
        calibration_equations[int(test_value)] = [int(x) for x in input_values.split()]


# Part 1


def operator_search(
    test_value: int, input_values: list[int], operators: list[Callable[[int, int], int]]
) -> bool:
    x, y = input_values[:2]

    if len(input_values) == 2:
        return any(operator(x, y) == test_value for operator in operators)

    remainder = input_values[2:]

    return any(
        operator_search(
            test_value=test_value,
            input_values=[operator(x, y), *remainder],
            operators=operators,
        )
        for operator in operators
    )


calibration_result: int = 0

operators: list[Callable[[int, int], int]] = [
    lambda x, y: x + y,
    lambda x, y: x * y,
]

for test_value, input_values in calibration_equations.items():
    if operator_search(
        test_value=test_value, input_values=input_values, operators=operators
    ):
        calibration_result += test_value

print("Part 1:", calibration_result)


# Part 2

calibrated_calibration_result: int = 0

operators.append(lambda x, y: int(f"{x}{y}"))

for test_value, input_values in calibration_equations.items():
    if operator_search(
        test_value=test_value, input_values=input_values, operators=operators
    ):
        calibrated_calibration_result += test_value

print("Part 2:", calibrated_calibration_result)
