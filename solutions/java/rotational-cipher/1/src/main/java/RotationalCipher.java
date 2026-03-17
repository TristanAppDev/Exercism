class RotationalCipher {

    private int shiftKey;

    RotationalCipher(int shiftKey) {
        this.shiftKey = shiftKey;
    }

    String rotate(String data) {
        StringBuilder result = new StringBuilder();

        for (var ch : data.toCharArray()) {
            if (Character.isLowerCase(ch)) {
                result.append(shift(ch, shiftKey, 'a'));
            } else if (Character.isUpperCase(ch)) {
                result.append(shift(ch, shiftKey, 'A'));
            } else {
                result.append(Character.toString(ch));
            }
        }

        return result.toString();
    }

    private char shift(char ch, int shiftKey, char base) {
        return (char) (base + ((ch - base) + shiftKey) % 26);
    }
}
