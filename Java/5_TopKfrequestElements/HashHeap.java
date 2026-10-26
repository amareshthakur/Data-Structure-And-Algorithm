
// Optimal solution approach 1(HashMap + Heap (priority Queue))
// Time -> O(n*logk)
// Space -> O(n)
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Queue;

public class HashHeap {
    public static void main(String[] args) {
        int[] nums = { 1, 3, 4, 3, 4, 2, 3, 4, 2, 5, 4, 5, 5 };
        // 1 -> 1
        // 2 -> 2
        // 3 -> 3
        // 4 -> 4
        // 5 -> 3
        Map<Integer, Integer> count = new HashMap<>();
        for (int n : nums) {
            count.put(n, count.getOrDefault(n, 0) + 1);
            System.out.println("n -> " + n + " [count.getOrDefault(n, 0)+1] -> " + count.getOrDefault(n, 0) + 1);
        }
    }

    static public int[] topKFrequent(int[] nums, int k) {
        if (k == nums.length) {
            return nums;
        }
        Map<Integer, Integer> count = new HashMap<>();
        for (int n : nums) {
            count.put(n, count.getOrDefault(n, 0) + 1);
        }
        Queue<Integer> heap = new PriorityQueue<>(
                (a, b) -> count.get(a) - count.get(b));
        for (int n : count.keySet()) {
            heap.add(n);
            if (heap.size() > k) {
                heap.poll();
            }
        }
        int[] ans = new int[k];
        for (int i = 0; i < k; i++) {
            ans[i] = heap.poll();
        }
        return ans;
    }
}
