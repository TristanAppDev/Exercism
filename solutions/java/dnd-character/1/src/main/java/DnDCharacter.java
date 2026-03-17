import java.util.Arrays;
import java.util.Random;

class DnDCharacter {
    private int strength, dexterity, constitution, intelligence, wisdom, charisma, hitpoints;

    public DnDCharacter() {
        this.strength = ability();
        this.dexterity = ability();
        this.constitution = ability();
        this.intelligence = ability();
        this.wisdom = ability();
        this.charisma = ability();
        this.hitpoints = 10 + modifier(this.constitution);
    }

    int ability() {
        int[] dices = rollDices();
        return Arrays.stream(dices).sorted().skip(1).reduce(Integer::sum).getAsInt();
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
        return this.hitpoints;
    }

    private int[] rollDices() {
        Random rnd = new Random();
        return Arrays.stream(new int[4]).map(dice -> dice = rnd.nextInt(6) + 1).toArray();
    }
}
