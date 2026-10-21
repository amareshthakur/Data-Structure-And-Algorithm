import java.util.Arrays;

/**
 * Better
 */
public class Better {
public static void main(String[] args) {
        int[] arr1 = {1, 2, 3, 3};
        int[] arr2 = {1, 2, 3, 4};
        int[] arr3 = {1,1,1,3,3,4,3,2,4,2};
        System.out.println(containsDuplicate(arr1));
        System.out.println(containsDuplicate(arr2));
        System.out.println(containsDuplicate(arr3));
        // containsDuplicate(arr3);
        // System.out.println("Array1 -> "+ Arrays.toString(arr3));
    }

    static public boolean containsDuplicate(int[] nums){
        Arrays.sort(nums);
        // System.out.println(Arrays.toString(nums));
        for (int i = 1; i < nums.length; i++) {
            if (nums[i-1] == nums[i]) {
                return true;
            }
        }
        return false;
    }
}