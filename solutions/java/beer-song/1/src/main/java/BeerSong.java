class BeerSong {

    String sing(int start, int takeDown) {
        int stop = start - takeDown + 1;
        StringBuilder song = new StringBuilder();

        for (int i = start; i >= stop; i--) {
            song.append(getVerseByBottleCount(i));
        }

        return song.toString();
    }

    String singSong() {
        return sing(99, 100);
    }

    private String getVerseByBottleCount(int count) {
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