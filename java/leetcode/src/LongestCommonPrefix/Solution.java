package LongestCommonPrefix;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

public class Solution {
    public String longestCommonPrefix (String[] strs) {
        List<String> list = Arrays.stream(strs).sorted(Comparator.comparingInt(String::length)).collect(Collectors.toList());
        String result = "";
        for (int i = 1; i < list.get(0).length() + 1; i++) {
            int count = 0;
            for (int j = 0; j < list.size(); j++) {
                String s = list.get(j);
                if (s.substring(0, i).equals(list.get(0).substring(0, i))) {
                    count++;
                }
            }
            if (count == list.size()) {
                result = list.get(0).substring(0, i);
            }
        }
        return result;
    }
}
