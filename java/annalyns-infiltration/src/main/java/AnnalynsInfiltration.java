class AnnalynsInfiltration {
    public static boolean canFastAttack(boolean knightIsAwake) {
        return !knightIsAwake;
    }

    public static boolean canSpy(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake) {
        return knightIsAwake || archerIsAwake || prisonerIsAwake;
    }

    public static boolean canSignalPrisoner(boolean archerIsAwake, boolean prisonerIsAwake) {
        return prisonerIsAwake && !archerIsAwake;
    }

    public static boolean canFreePrisoner(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake, boolean petDogIsPresent) {
        return onlyPrisonerAwake(knightIsAwake, archerIsAwake, prisonerIsAwake) || canUseDog(archerIsAwake, petDogIsPresent);
    }

    private static boolean onlyPrisonerAwake(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake) {
        return prisonerIsAwake && !knightIsAwake && !archerIsAwake;
    }

    private static boolean canUseDog(boolean archerIsAwake, boolean petDogIsPresent) {
        return !archerIsAwake && petDogIsPresent;
    }

}
