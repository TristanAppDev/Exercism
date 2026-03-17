class ReverseString {
    // Solution without build in function
    String reverse(String inputString) {
        StringBuilder result = new StringBuilder();
        for (int i = inputString.length() - 1; i >= 0; i--) {
            result.append(inputString.charAt(i));
        }
        return result.toString();
    }
}
