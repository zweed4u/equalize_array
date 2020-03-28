#!/usr/bin/python3


def equalizeArray(arr):
    # arr = [3,3,2,1,3]
    numbers = dict()
    for number in arr:
        if numbers.get(number) is None:
            numbers.update({number: 1})
        else:
            numbers[number] += 1
    print(numbers)
    # {1:1, 2:1, 3:3}
    most_occurrences = max(list(numbers.values()))
    kept = False
    amount_to_delete = 0
    for number, amount in numbers.items():
        if amount == most_occurrences and kept is False:
            kept = True
            continue
        amount_to_delete += amount
    print(amount_to_delete)
    return amount_to_delete


equalizeArray([1, 2, 2, 3])  # 2
equalizeArray([1, 2, 2, 3, 4, 5])  # 4
equalizeArray([3, 3, 2, 1, 3])  # 2

equalizeArray([3, 3, 2, 2, 3, 1, 1, 3])  # 4

equalizeArray([1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3])  # 10
