package PalindromeNumber;

public class Solution {
    public boolean isPalindrome(int x) {
        String str = x + "";
        if (str.indexOf('-') == 0) {
            return false;
        }
        String[] arr = str.split("");
        String result = "";
        for(String s: arr) {
            result = s + result;
        }
        return result.equals(str);
    }
}
