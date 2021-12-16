import java.util.Dictionary;
import java.util.Hashtable;

public class Blackjack {

    private Dictionary<String, String> decision = new Hashtable<>();

    public Blackjack() {
        decision.put("Stand", "S");
        decision.put("Hit", "H");
        decision.put("Split", "P");
        decision.put("Win", "W");
    }

    public int parseCard(String card) {
        switch (card) {
            case "two":
                return 2;
            case "three":
                return 3;
            case "four":
                return 4;
            case "five":
                return 5;
            case "six":
                return 6;
            case "seven":
                return 7;
            case "eight":
                return 8;
            case "nine":
                return 9;
            case "ten":
                return 10;
            case "jack":
                return 10;
            case "queen":
                return 10;
            case "king":
                return 10;
            case "ace":
                return 11;
            default:
                return 0;
        }
    }

    public boolean isBlackjack(String card1, String card2) {
        return (parseCard(card1) == 10 && card2.equals("ace")) || (parseCard(card2) == 10 && card1.equals("ace"));
    }

    public String largeHand(boolean isBlackjack, int dealerScore) {
        return "";
    }

    public String smallHand(int handScore, int dealerScore) {
        throw new UnsupportedOperationException("Please implement the Blackjack.smallHand method");
    }

    // FirstTurn returns the semi-optimal decision for the first turn, given the
    // cards of the player and the dealer.
    // This function is already implemented and does not need to be edited. It pulls
    // the other functions together in a
    // complete decision tree for the first turn.
    public String firstTurn(String card1, String card2, String dealerCard) {
        int handScore = parseCard(card1) + parseCard(card2);
        int dealerScore = parseCard(dealerCard);

        if (20 < handScore)
            return largeHand(isBlackjack(card1, card2), dealerScore);

        else
            return smallHand(handScore, dealerScore);
    }
}
