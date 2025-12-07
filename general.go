FUNCTION Compact(list):

    count = 0

    FOR each item in the list:
        IF item is NOT empty:
            put item at position count
            count = count + 1
        END IF
    END FOR

    cut the list so it keeps only the first 'count' items

    RETURN count
END FUNCTION
