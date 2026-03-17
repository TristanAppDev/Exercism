public class Twofer {
    public String twofer(String name) {
        if(name == null) name = "you";
        return "One for " + name + ", one for me.";
    }
    public String twofer() {
        return twofer("you");
    }
}
