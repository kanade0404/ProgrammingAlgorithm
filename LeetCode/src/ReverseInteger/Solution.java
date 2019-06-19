package ReverseInteger;

public class Solution {
    public int reverse(int x) {
        String st = x + "";
        boolean flg = false;
        if (st.indexOf('-') == 0) {
            flg = true;
            st = st.replace("-", "");
        }
        String[] arr = st.split("");
        String result = "";
        for(String s: arr) {
            result = s + result;
        }
        if (flg) {
            result = "-" + result;
        }
        try {
            return Integer.parseInt(result);
        } catch (NumberFormatException ex) {
            return 0;
        }
    }
}
