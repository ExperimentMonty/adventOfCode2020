

file_input = open("inputs/6.txt")

group = set()
count_sum = 0

for line in file_input:
    line = line.strip()
    if line == "":
        # Process group
        count_sum += len(group)
        group.clear()
    else:
        # Add questions to the group set
        for question in line:
            group.add(question)

# Do final group
count_sum += len(group)

print(count_sum)