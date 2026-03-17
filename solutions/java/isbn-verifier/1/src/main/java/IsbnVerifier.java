class IsbnVerifier {

    boolean isValid(String stringToVerify) {
        int sum = 0;
        int counter = 10;

        for (int i = 0; i < stringToVerify.length(); i++) {
            char ch = stringToVerify.charAt(i);
            if (ch >= '0' && ch <= '9') {
                sum += Character.getNumericValue(ch) * counter;
            } else if (ch == 'X' && i == stringToVerify.length() - 1) {
                sum += 10 * counter;
            } else if (ch == '-') {
                continue;
            } else {
                return false;
            }
            counter--;
        }

        return counter == 0 && sum % 11 == 0;
    }

}
