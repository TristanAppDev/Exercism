class IsbnVerifier {

    boolean isValid(String stringToVerify) {
        var sum = 0;
        var counter = 10;

        for (var i = 0; i < stringToVerify.length(); i++) {
            var ch = stringToVerify.charAt(i);
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
