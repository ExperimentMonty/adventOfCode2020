from collections import defaultdict

file_input = open("inputs/6.txt")

group = set()
count_sum = 0

group_all = defaultdict(int)
group_member_count = 0
group_all_count = 0

for line in file_input:
    line = line.strip()
    if line == "":
        # Process group
        count_sum += len(group)
        group.clear()

        # For part 2
        group_all_count += len([k for k in group_all if group_all[k] == group_member_count])
        group_member_count = 0
        group_all.clear()
    else:
        # Count group member
        group_member_count += 1
        # Add questions to the group set
        for question in line:
            group.add(question)
            group_all[question] += 1


# Do final group
count_sum += len(group)
group_all_count += len([k for k in group_all if group_all[k] == group_member_count])

print(count_sum)
print(group_all_count)