import java.util.List;

public class Knapsack {

    public int maximumValue(final int totalWeight, final List<Item> items) {
        final int[] weights = items.stream().mapToInt(Item::getValue).toArray();
        final int[] values = items.stream().mapToInt(Item::getWeight).toArray();
        return recursiveKnapsack(weights, values, items.size(), totalWeight);
    }

    private int recursiveKnapsack(int[] weights, int[] values, int totalItemCount, int weightLimit) {
        if (totalItemCount <= 0) {
            return 0;
        } else if (weights[totalItemCount - 1] > weightLimit) {
            return recursiveKnapsack(weights, values, totalItemCount - 1, weightLimit);
        } else {
            return Math.max(
                    recursiveKnapsack(weights, values, totalItemCount - 1, weightLimit), values[totalItemCount - 1]
                            + recursiveKnapsack(weights, values, totalItemCount - 1,
                                    weightLimit - weights[totalItemCount - 1]));
        }
    }
}

class Item {
    private final int value;
    private final int weight;

    public Item(final int value, final int weight) {
        this.value = value;
        this.weight = weight;
    }

    public int getWeight() {
        return weight;
    }

    public int getValue() {
        return value;
    }
}
