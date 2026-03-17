import java.util.HashMap;

public class Blackjack {

    private final String STRATEGY_STAND = "S";
    private final String STRATEGY_HIT = "H";
    private final String STRATEGY_SPLIT = "P";
    private final String STRATEGY_WIN = "W";

    private final HashMap<String, Integer> cards = new HashMap<>() {
        {
            put("two", 2);
            put("three", 3);
            put("four", 4);
            put("five", 5);
            put("six", 6);
            put("seven", 7);
            put("eight", 8);
            put("nine", 9);
            put("ten", 10);
            put("jack", 10);
            put("queen", 10);
            put("king", 10);
            put("ace", 11);
        }
    };

    public int parseCard(String card) {
        return cards.getOrDefault(card, 0);
    }

    public boolean isBlackjack(String card1, String card2) {
        return parseCard(card1) + parseCard(card2) == 21;
    }

    public String largeHand(boolean isBlackjack, int dealerScore) {
        if (!isBlackjack) {
            return STRATEGY_SPLIT;
        }

        if (dealerScore < 10) {
            return STRATEGY_WIN;
        }

        return STRATEGY_STAND;
    }

    public String smallHand(int handScore, int dealerScore) {
        if (handScore >= 17 || (handScore >= 12 && dealerScore < 7)) {
            return STRATEGY_STAND;
        }
        return STRATEGY_HIT;
    }

    // FirstTurn returns the semi-optimal decision for the first turn, given the
    // cards of the player and the dealer.
    // This function is already implemented and does not need to be edited. It pulls
    // the other functions together in a
    // complete decision tree for the first turn.
    public String firstTurn(String card1, String card2, String dealerCard) {
        int handScore = parseCard(card1) + parseCard(card2);
        int dealerScore = parseCard(dealerCard);

        if (20 < handScore) {
            return largeHand(isBlackjack(card1, card2), dealerScore);
        } else {
            return smallHand(handScore, dealerScore);
        }
    }
}
