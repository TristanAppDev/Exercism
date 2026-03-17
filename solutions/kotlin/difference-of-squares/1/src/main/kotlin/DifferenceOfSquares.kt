class Squares constructor(private val n: Int) {

    fun sumOfSquares(): Int {
        return (n * (n + 1) * (2 * n + 1)) / 6
    }

    fun squareOfSum(): Int {
        return pow((n * (n + 1)) / 2)
    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }

    private fun pow(n: Int): Int {
        return n * n
    }
}
