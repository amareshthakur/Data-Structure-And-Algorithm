import java.util.Arrays;

/**
 * optimal
 */
public class optimal {
    public static void main(String[] args) {
        int[] arrw = { 3, 2, 4 };
        twoSum(arrw, 6);
    }
    static public int[] twoSum(int[] nums, int target){
        Arrays.sort(nums);
        System.out.println(Arrays.toString(nums));
        return new int[]{};
    }
}

// Note: we do NOT sort + binary search here, because sorting
// destroys the original indices (i, j) that we must return.
// Sorting would only work if the method returned a boolean
// (true = a pair summing to target exists, false = it doesn't).