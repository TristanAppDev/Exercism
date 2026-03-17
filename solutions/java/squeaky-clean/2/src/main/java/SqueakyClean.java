class SqueakyClean {
    static String clean(String identifier) {

        StringBuilder result = new StringBuilder();
        boolean kebab = false;

        for (char ch : identifier.toCharArray()) {
            if (ch == '-') {
                kebab = true;
            } else if (isPermittedLetter(ch)) {
                result.append(kebab ? Character.toUpperCase(ch) : ch);
                kebab = false;
            } else if (ch == ' ') {
                result.append('_');
            } else if (Character.isISOControl(ch)) {
                result.append("CTRL");
            }
        }

        return result.toString();
    }

    private static boolean isPermittedLetter(char ch) {
        return Character.isAlphabetic(ch) && (ch < 'α' || 'ω' < ch);
    }
}
