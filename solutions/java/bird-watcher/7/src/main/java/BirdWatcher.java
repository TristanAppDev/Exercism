import java.util.Arrays;
import java.util.stream.IntStream;

class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return birdsPerDay.clone();
    }

    public int getToday() {
        return (birdsPerDay.length == 0) ? 0 : birdsPerDay[birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        birdsPerDay[birdsPerDay.length - 1]++;
    }

    public boolean hasDayWithoutBirds() {
        return IntStream.of(birdsPerDay).anyMatch(x -> x == 0);
    }

    public int getCountForFirstDays(int numberOfDays) {
        return IntStream.of(getFirstDays(numberOfDays)).sum();
    }

    private int[] getFirstDays(int numberOfDays) {
        return Arrays.copyOfRange(birdsPerDay, 0, numberOfDays);
    } 

    public int getBusyDays() {
        return (int) IntStream.of(birdsPerDay).filter(x -> x >= 5).count();
    }
}
