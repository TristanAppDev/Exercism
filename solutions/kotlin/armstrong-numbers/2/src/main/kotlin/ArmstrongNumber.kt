import kotlin.math.log10

object ArmstrongNumber {

    fun check(input: Int): Boolean {
        if (input < 10) {
            return true
        }

        val exp = log10(input.toDouble()).toInt() + 1;
        var result = 0
        var remainder = input

        while (remainder > 0) {
            result += pow(remainder % 10, exp)
            remainder /= 10
        }

        return input == result
    }

    private fun pow(n: Int, exp: Int): Int {
        var exponent = exp
        var result = 1

        while (exponent != 0) {
            result *= n
            --exponent
        }
        return result
    }
}
