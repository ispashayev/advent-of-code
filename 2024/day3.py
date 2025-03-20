import re


with open('day3.txt') as f:
    memory = f.read()


# Part 1

def get_multiplicative_sum(memory: str) -> int:
    mul_pattern = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)', re.DOTALL)
    all_matches = mul_pattern.findall(memory)
    return sum(int(x)*int(y) for x, y in all_matches)

print('Part 1:', get_multiplicative_sum(memory=memory))


# Part 2

multiplicative_sum = 0

enabled_pattern = re.compile(r'(.*?)(don\'t\(\)|$)', re.DOTALL)
disabled_pattern = re.compile(r'(.*?)(do\(\)|$)', re.DOTALL)

i, enabled = 0, True
while i < len(memory):
    remaining_memory = memory[i:]
    
    if enabled:
        memory_match = enabled_pattern.match(remaining_memory)
        multiplicative_sum += get_multiplicative_sum(memory=memory_match.group())
    else:
        memory_match = disabled_pattern.match(remaining_memory)
    
    i += max(1, memory_match.end())

    enabled = not enabled

print('Part 2:', multiplicative_sum)
