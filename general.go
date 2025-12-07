Initialize an empty map called summary
Initialize an empty string called currentItem

For each character in the string:
    If the character is a space:
        Increase summary[currentItem] by 1
        Reset currentItem to empty string
    Else:
        Add the character to currentItem

After the loop:
    If currentItem is not empty:
        Increase summary[currentItem] by 1
