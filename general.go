func ShoppingSummaryCounter(str string) map[string]int {

    
summary := make(map[string]int)
currentItem := ""

For each character in the string:
    If the character is a space:
        Increase summary[currentItem] by 1
        Reset currentItem to empty string
    Else:
        Add the character to currentItem

After the loop:
    If currentItem is not empty:
        Increase summary[currentItem] by 1


}


