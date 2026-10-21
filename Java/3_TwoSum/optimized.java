import java.util.Arrays;
import java.util.HashMap;

/**
 * optimized
 */
public class optimized {
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
    static public int[] twoSum(int[] nums, int target){
        HashMap<Integer,Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[]{map.get(complement),i};
            }
            map.put(nums[i], i);
        }
            return new int[]{};
        }
}