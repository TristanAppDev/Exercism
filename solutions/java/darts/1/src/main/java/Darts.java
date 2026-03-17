class Darts {
    int score(double xOfDart, double yOfDart) {

        double magnitude = getMagnitude(xOfDart, yOfDart);

        if (magnitude > 10.0)
            return 0;

        if (magnitude > 5.0)
            return 1;

        if (magnitude > 1.0)
            return 5;

        return 10;
    }

    private double getMagnitude(double xOfDart, double yOfDart) {
        return Math.sqrt((xOfDart * xOfDart) + (yOfDart * yOfDart));
    }
}
