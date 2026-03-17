object Raindrops {

    fun convert(n: Int): String {
        var result = "";
            
    	result += if (n % 3 == 0) "Pling" else ""
    	result += if (n % 5 == 0) "Plang" else ""
    	result += if (n % 7 == 0) "Plong" else ""
    
    	if (result == "") return n.toString()
    
    	return result
    }
}
