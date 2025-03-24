with open("day9.txt") as f:
    disk_map = f.read().strip()

# A base map that will be copied for parts 1 and 2
_block_map: list[int | None] = [
    None for contiguous_blocks in disk_map for _ in range(int(contiguous_blocks))
]

# Helper data structures for part 2
file_index: list[
    tuple[int, int, int]
] = []  # Each element represents file_id, file_start_block_index, file_size
free_space_index: list[
    tuple[int, int]
] = []  # Each element represents free_space_start, free_space_end


# Initialize data structures
disk_head = 0
for i in range(0, len(disk_map), 2):
    file_id, file_size = i // 2, int(disk_map[i])
    file_index.append((file_id, disk_head, file_size))
    _block_map[disk_head : disk_head + file_size] = (file_id,) * file_size

    disk_head += file_size

    if i + 1 < len(disk_map):
        free_space = int(disk_map[i + 1])
        free_space_index.append((disk_head, disk_head + free_space))

        disk_head += free_space


# Part 1

part1_block_map = _block_map[:]  # makes a copy

l, r = 0, len(part1_block_map) - 1

while l < r:
    if part1_block_map[l] is not None:
        l += 1
    elif part1_block_map[r] is None:
        r -= 1
    else:
        part1_block_map[l] = part1_block_map[r]
        part1_block_map[r] = None
        l += 1
        r -= 1

part1_checksum = sum(
    i * file_id for i, file_id in enumerate(part1_block_map) if file_id is not None
)
print("Part 1:", part1_checksum)


# Part 2

part2_block_map = _block_map[:]

for file_id, file_start_block_index, file_size in reversed(file_index):
    i, file_moved = 0, False
    free_space_start, free_space_end = free_space_index[i]

    while (
        i < len(free_space_index)
        and not file_moved
        and free_space_start < file_start_block_index
    ):
        free_space = free_space_end - free_space_start

        if free_space >= file_size:
            # Whole file can fit in the current free space range
            for offset in range(file_size):
                part2_block_map[free_space_start + offset] = file_id
                part2_block_map[file_start_block_index + offset] = None

            # Update the free space index
            if file_size == free_space:
                del free_space_index[i]
            else:
                free_space_index[i] = (free_space_start + file_size, free_space_end)

            file_moved = True

        # Try the next available free space range
        i += 1
        free_space_start, free_space_end = free_space_index[i]

part2_checksum = sum(
    i * file_id for i, file_id in enumerate(part2_block_map) if file_id is not None
)
print("Part 2:", part2_checksum)
