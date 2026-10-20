class Brute{
    public static void main(String[] args) {
        // System.out.println("Hello world");
        int[] ar1 = {1,2,3,1};
        int[] ar2 = {1,2,3,4};
        int[] ar3 = {1,1,1,3,3,4,3,2,4,2};

        System.out.println(containsDuplicate(ar1));
        System.out.println(containsDuplicate(ar2));
        System.out.println(containsDuplicate(ar3));
    }
    static boolean containsDuplicate(int[] nums){
        for (int i = 0; i < nums.length; i++) {
            // int count = 1;
            for (int j = i+1; j < nums.length; j++) {
                if (nums[i] == nums[j]) {
                    return true;
                }
            }
        }
        return false;
    }
}

// Time complexity of this portion is O(N^2).
// Space complexity is O(N) due to the array,
// and the auxiliary space complexity is O(1).