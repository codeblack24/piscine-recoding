FUNCTION DealAPackOfCards(deck):

    There are 12 cards in the deck
    There are 4 players
    Each player should receive 3 cards

    FOR player number from 1 to 4:

        Print "Player X: " (where X is the player number)

        Find the first card for this player:
            start = (player number - 1) * 3

        Find the last card for this player:
            end = start + 3

        FOR every card index from start to end - 1:

            IF this is the last card for the player:
                Print the card without a comma
            ELSE:
                Print the card followed by a comma and space

        END FOR

        Print a new line

    END FOR

END FUNCTION
