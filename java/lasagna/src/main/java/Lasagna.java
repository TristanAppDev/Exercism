public class Lasagna {

    private final int expectedOvenTime = 40; 
    
    public int expectedMinutesInOven() {
        return expectedOvenTime;
    }

    public int remainingMinutesInOven(int ovenTime) {
        return expectedOvenTime - ovenTime;
    }

    public int preparationTimeInMinutes(int layersCount) {
        return layersCount * 2;
    }

    public int totalTimeInMinutes(int layersCount, int ovenTime) {
        return preparationTimeInMinutes(layersCount) + ovenTime;
    }
}
