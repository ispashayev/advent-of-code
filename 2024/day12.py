from dataclasses import dataclass

with open("day12.txt") as f:
    plots: dict[complex, str] = {
        i + j * 1j: plant
        for i, line in enumerate(f.readlines())
        for j, plant in enumerate(line.strip())
    }


# Parts 1 and 2


@dataclass
class Region:
    plant_type: str
    plots: set[complex]
    perimeter: int
    sides: int

    @property
    def area(self) -> int:
        return len(self.plots)


transitions = (1, 1j, -1, -1j)


def explore_region(start_position: complex) -> Region:
    region = Region(
        plant_type=plots[start_position],
        plots=set(),
        perimeter=0,
        sides=0,
    )

    def _explore_region_dfs(position: complex):
        if position in region.plots:
            return

        region.plots.add(position)

        next_plots: list[complex] = []

        for transition in transitions:
            next_position = position + transition
            ortho_transition = transition * -1j
            ortho_position = position + ortho_transition
            diagonal_position = next_position + ortho_transition

            if plots.get(next_position) == region.plant_type:
                next_plots.append(next_position)
                if (
                    plots.get(ortho_position) == region.plant_type
                    and plots.get(diagonal_position) != region.plant_type
                ):
                    region.sides += 1
            else:
                if plots.get(ortho_position) != region.plant_type:
                    region.sides += 1

        region.perimeter += 4 - len(next_plots)

        for plot in next_plots:
            _explore_region_dfs(position=plot)

    _explore_region_dfs(position=start_position)

    return region


explored_positions: set[complex] = set()
regions: list[Region] = []

for position in plots:
    if position not in explored_positions:
        region = explore_region(start_position=position)
        explored_positions.update(region.plots)
        regions.append(region)


total_fence_price_by_perimeter = sum(
    region.area * region.perimeter for region in regions
)

print("Part 1:", total_fence_price_by_perimeter)

total_fence_price_by_sides = sum(region.area * region.sides for region in regions)

print("Part 2:", total_fence_price_by_sides)
