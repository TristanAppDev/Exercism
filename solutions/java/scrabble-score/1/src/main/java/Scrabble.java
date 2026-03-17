class Scrabble {

    private String word;

    Scrabble(String word) {
        this.word = word;
    }

    int getScore() {
        return word.toLowerCase().chars().reduce(0, (score, letter) -> score + getLetterScore(letter));
    }

    private int getLetterScore(int letter) {
        switch (letter) {
            case 'q', 'z':
                return 10;
            case 'j', 'x':
                return 8;
            case 'k':
                return 5;
            case 'f', 'h', 'v', 'w', 'y':
                return 4;
            case 'b', 'c', 'm', 'p':
                return 3;
            case 'd', 'g':
                return 2;
            default:
                return 1;
        }
    }
}
