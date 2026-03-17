import java.util.Arrays;

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
        ++birdsPerDay[birdsPerDay.length - 1];
    }

    public boolean hasDayWithoutBirds() {
        for (int birdCount : birdsPerDay) {
            if (birdCount == 0)
                return true;
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        return Arrays.stream(birdsPerDay).limit(numberOfDays).reduce(0, Integer::sum);
    }

    private int[] getFirstDays(int numberOfDays) {
        return Arrays.copyOfRange(birdsPerDay, 0, numberOfDays);
    }

    public int getBusyDays() {
        return (int) Arrays.stream(birdsPerDay).filter(x -> x >= 5).count();
    }
}
