FUNCTION Compact(pointer to list):

    Let s be the list we got from the pointer

    Start a counter at 0

    FOR each position i in the list:
        IF the item at position i is NOT an empty string:
            
            Move this item to the front of the list
            Specifically, put it at position equal to the counter

            Increase the counter by 1

        END IF
    END FOR

    Now remove everything after the counter position
    (this cuts off all empty strings at the end)

    Update the pointer so it points to this shorter list

    RETURN the counter (the number of non-empty items)
END FUNCTION
