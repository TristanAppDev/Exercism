public class SalaryCalculator {
    
    private final double BASE_SALARY = 1000.0;
    private final double MAXIMUM_SALARY = 2000.0;

    private final int DAYS_WITHOUT_PENALTY = 5;
    private final double MULTIPLIER_WITHOUT_PENALTY = 1;
    private final double MULTIPLIER_WITH_PENALTY = 0.85;

    private final int LIMIT_SOLD_WITHOUT_IMPROVEMENT = 20;
    private final int MULTIPLIER_PER_PRODUCT_NORMAL = 10;
    private final int MULTIPLIER_PER_PRODUCT_IMPROVED = 13;

    public double multiplierPerDaysSkipped(int daysSkipped) {
        return (daysSkipped <= DAYS_WITHOUT_PENALTY) ? MULTIPLIER_WITHOUT_PENALTY : MULTIPLIER_WITH_PENALTY;
    }

    public int multiplierPerProductsSold(int productsSold) {
        return (productsSold <= LIMIT_SOLD_WITHOUT_IMPROVEMENT) ? MULTIPLIER_PER_PRODUCT_NORMAL : MULTIPLIER_PER_PRODUCT_IMPROVED;
    }

    public double bonusForProductSold(int productsSold) {
        return productsSold * multiplierPerProductsSold(productsSold);
    }

    public double finalSalary(int daysSkipped, int productsSold) {
        double finalSalary = BASE_SALARY * multiplierPerDaysSkipped(daysSkipped) + bonusForProductSold(productsSold);
        return (finalSalary > MAXIMUM_SALARY) ? MAXIMUM_SALARY : finalSalary;
    } 
}
