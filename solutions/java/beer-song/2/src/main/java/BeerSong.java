import java.util.stream.Collectors;
import java.util.stream.IntStream;

class BeerSong {

    String sing(int start, int takeDown) {
        return IntStream
                .range(0, takeDown)
                .mapToObj(index -> verse(start - index))
                .collect(Collectors.joining());
    }

    String singSong() {
        return sing(99, 100);
    }

    private String verse(int count) {
        switch (count) {
            case 0:
                return "No more bottles of beer on the wall, no more bottles of beer.\n"
                        + "Go to the store and buy some more, 99 bottles of beer on the wall.\n\n";
            case 1:
                return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, "
                        + "no more bottles of beer on the wall.\n\n";
            case 2:
                return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, "
                        + "1 bottle of beer on the wall.\n\n";
            default:
                return String.format("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and "
                        + "pass it around, %d bottles of beer on the wall.\n\n", count, count, count - 1);
        }
    }
}