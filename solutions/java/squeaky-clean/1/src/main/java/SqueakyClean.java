class SqueakyClean {
    static String clean(String identifier) {

        StringBuilder result = new StringBuilder();
        boolean kebab = false;

        for (char character : identifier.toCharArray()) {
            if (character == '-') {
                kebab = true;
            } else if (isPermittedLetter(character)) {
                result.append(kebab ? Character.toUpperCase(character) : character);
                kebab = false;
            } else if (character == ' ') {
                result.append('_');
            } else if (Character.isISOControl(character)) {
                result.append("CTRL");
            }
        }

        return result.toString();
    }

    private static boolean isPermittedLetter(char ch) {
        return Character.isAlphabetic(ch) && (ch < 'α' || 'ω' < ch);
    }
}
