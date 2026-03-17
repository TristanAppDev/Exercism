using System;
using System.Linq;

public static class PhoneNumber
{
    private const string NEW_YORK_DIALING_CODE = "212";
    private const string FAKE_PREFIX_CODE = "555";

    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber) => 
        (
            phoneNumber.Substring(0, 3).Equals(NEW_YORK_DIALING_CODE), 
            phoneNumber.Substring(4, 3).Equals(FAKE_PREFIX_CODE),
            phoneNumber.Substring(8, 4)
        );

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) => 
        phoneNumberInfo.IsFake;
}
