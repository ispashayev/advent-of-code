with open("day2.txt") as f:
    reports = []
    for line in f.readlines():
        reports.append([int(e) for e in line.split()])


# Part 1

num_safe = 0


def is_safe(report: list[int]) -> bool:
    """
    Returns a boolean indicating if the report's levels are safe or not.
    """
    is_decreasing = report[1] < report[0]

    for i in range(1, len(report)):
        if is_decreasing:
            diff = report[i - 1] - report[i]
        else:
            diff = report[i] - report[i - 1]

        if not 1 <= diff <= 3:
            return False

    return True


for report in reports:
    if is_safe(report=report):
        num_safe += 1

print("Part 1:", num_safe)


# Part 2

num_safe_dampened = 0


def is_safe_dampened(report: list[int], dampened: bool) -> int | None:
    for i in range(1, len(report)):
        diff = report[i] - report[i - 1]

        if not 1 <= diff <= 3:
            if dampened:
                return False

            cut_forward_report = report[:i] + report[i + 1 :]
            cut_backward_report = report[: i - 1] + report[i:]
            return is_safe_dampened(
                report=cut_forward_report, dampened=True
            ) or is_safe_dampened(report=cut_backward_report, dampened=True)

    return True


for report in reports:
    if is_safe_dampened(report=report, dampened=False) or is_safe_dampened(
        report=report[::-1], dampened=False
    ):
        num_safe_dampened += 1

print("Part 2:", num_safe_dampened)
