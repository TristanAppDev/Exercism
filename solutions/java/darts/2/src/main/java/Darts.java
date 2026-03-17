class Darts {

    private static final int RADIUS_CIRCLE_OUTER = 10;
    private static final int RADIUS_CIRCLE_MIDDLE = 5;
    private static final int RADIUS_CIRCLE_INNER = 1;

    int score(double xOfDart, double yOfDart) {

        double magnitude = getMagnitude(xOfDart, yOfDart);

        if (magnitude > RADIUS_CIRCLE_OUTER)
            return 0;
        if (magnitude > RADIUS_CIRCLE_MIDDLE)
            return 1;
        if (magnitude > RADIUS_CIRCLE_INNER)
            return 5;

        return 10;
    }

    private double getMagnitude(double xOfDart, double yOfDart) {
        return Math.sqrt((xOfDart * xOfDart) + (yOfDart * yOfDart));
    }
}
