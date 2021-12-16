import java.util.Random;

class DnDCharacter {
    private int strength, dexterity, constitution, intelligence, wisdom, charisma;
    private Random rnd;

    public DnDCharacter() {
        this.rnd = new Random();
        this.strength = ability();
        this.dexterity = ability();
        this.constitution = ability();
        this.intelligence = ability();
        this.wisdom = ability();
        this.charisma = ability();
    }

    int ability() {
        return rnd.ints(4, 1, 7).sorted().skip(1).sum();
    }

    int modifier(int constitution) {
        return Math.floorDiv(constitution - 10, 2);
    }

    int getStrength() {
        return this.strength;
    }

    int getDexterity() {
        return this.dexterity;
    }

    int getConstitution() {
        return this.constitution;
    }

    int getIntelligence() {
        return this.intelligence;
    }

    int getWisdom() {
        return this.wisdom;
    }

    int getCharisma() {
        return this.charisma;
    }

    int getHitpoints() {
        return 10 + modifier(this.constitution);
    }
}
