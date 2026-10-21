import java.util.Arrays;

/**
 * Brute
 */
public class Brute {
    public static void main(String[] args) {
        int[] arr1 = { 2, 7, 11, 15 };
        int target = 9;
        int[] index = twoSum(arr1, target);
        // System.out.println("Numbers are at index " + Arrays.toString(index));
        System.out.println("Numbers are at index " + Arrays.toString(index) + " To get sum " + target +" In Array "+Arrays.toString(arr1));
        int[] arrw = { 3, 2, 4 };
        target = 6;
        index = twoSum(arrw, target);
        System.out.println("Numbers are at index " + Arrays.toString(index) + " To get sum " + target+" In Array "+Arrays.toString(arrw));
    }

    static public int[] twoSum(int[] nums, int target) {
        //
        for (int i = 0; i < nums.length; i++) {
            int requiredNum = target - nums[i];
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] == requiredNum) {
                    return new int[] { i, j };
                }
            }
        }
        return new int[] {};
    }
}