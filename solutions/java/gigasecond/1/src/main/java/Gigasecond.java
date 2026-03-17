import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {
    private final LocalDateTime moment;

    public Gigasecond(LocalDate moment) {
        this(moment.atStartOfDay());
    }

    public Gigasecond(LocalDateTime moment) {
        long GIGA_SECOND = 1_000_000_000;
        this.moment = moment.plusSeconds(GIGA_SECOND);
    }

    public LocalDateTime getDateTime() {
        return this.moment;
    }
}
