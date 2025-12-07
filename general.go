FUNCTION JumpOver(str):

    IF length of str < 3 THEN
        RETURN newline ("\n")
    END IF

    result ← empty string

    FOR i from 2 to length of str - 1, step 3 DO
        result ← result + str[i]
    END FOR

    RETURN result + newline ("\n")

END FUNCTION
