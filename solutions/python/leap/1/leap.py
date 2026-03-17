"""leap_year exercise"""

def leap_year(year):
    """return is year a leap year"""
    return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)
