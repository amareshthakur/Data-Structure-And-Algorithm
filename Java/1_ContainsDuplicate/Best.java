import java.util.HashSet;

/**
 * Best
 */
public class Best {
    public static void main(String[] args) {
        int[] arr1 = {1, 2, 3, 3};
        int[] arr2 = {1, 2, 3, 4};
        int[] arr3 = {1,1,1,3,3,4,3,2,4,2};
        // System.out.println(containsDuplicate(arr1));
        // System.out.println(containsDuplicate(arr2));
        // System.out.println(containsDuplicate(arr3));
        containsDuplicate(arr3);
    }

     static public void containsDuplicate(int[] nums){
        HashSet<Integer> seenNumbers = new HashSet<>();
        for (Integer num : nums) {
            // if (seenNumbers.contains(num)) {
                // return true;
            // }
            seenNumbers.add(num);
        }
        // return false;
        System.out.println(seenNumbers);
    }    
}