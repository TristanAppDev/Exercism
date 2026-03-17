public class LogLevels {
    
    public static String message(String logLine) {
        int startOfMessage = logLine.indexOf(":");
        return logLine.substring(startOfMessage + 2).trim();
    }

    public static String logLevel(String logLine) {
        return logLine.split(" ")[0].replaceAll("[^A-Za-z]", "").toLowerCase(); 
    }

    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
    }
}
