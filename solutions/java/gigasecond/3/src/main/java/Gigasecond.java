import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {

    private final static int GIGA_SECOND = 1_000_000_000;

    private final LocalDateTime moment;

    public Gigasecond(LocalDate moment) {
        this(moment.atStartOfDay());
    }

    public Gigasecond(LocalDateTime moment) {
        this.moment = moment.plusSeconds(GIGA_SECOND);
    }

    public LocalDateTime getDateTime() {
        return this.moment;
    }
}
