package TwoSum;

import java.util.Arrays;

public class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] result = new int[2];
        for (int k = 0; k < nums.length; k++) {
            for(int v = 1; v < nums.length; v++) {
                if (k != v && nums[k] + nums[v] == target){
                    result[0] = k;
                    result[1] = v;
                    return result;
                }
            }
        }
        return result;
    }
}
