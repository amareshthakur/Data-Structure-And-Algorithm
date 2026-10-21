class Brute{
    public static void main(String[] args) {
        int[] arr1 = {1, 2, 3, 3};
        int[] arr2 = {1, 2, 3, 4};
        int[] arr3 = {1,1,1,3,3,4,3,2,4,2};
        System.out.println(containsDuplicate(arr1));
        System.out.println(containsDuplicate(arr2));
        System.out.println(containsDuplicate(arr3));
    }

    static public boolean containsDuplicate(int[] nums){
        for (int i = 0; i < nums.length; i++) {
            for (int j = i+1; j < nums.length; j++) {
                if (nums[i] == nums[j]){
                    return true;
                }
            }
        }
        return false;
    }
}