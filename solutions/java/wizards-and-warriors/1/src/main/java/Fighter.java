abstract class Fighter {

    boolean isVulnerable() {
        return false;
    }

    abstract int damagePoints(Fighter fighter);

    @Override
    public String toString() {
        return "Fighter is a " + this.getClass().getName();           
    }
}

class Warrior extends Fighter {

    @Override
    public String toString() {
        return super.toString();
    }

    @Override
    int damagePoints(Fighter wizard) {
        return (wizard.isVulnerable()) ? 10 : 6;
    }
}

class Wizard extends Fighter {

    private boolean isPrepared;

    @Override
    boolean isVulnerable() {
        return !isPrepared;
    }

    @Override
    int damagePoints(Fighter warrior) {
        return (isPrepared) ? 12 : 3;
    }

    void prepareSpell() {
        isPrepared = true;
    }

}
