// https://leetcode.com/problems/group-anagrams/ -> Problem 49

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Brute
 */
public class optimized {
    public static void main(String[] args) {
        String[] stars = { "eat", "tea", "tan", "ate", "nat", "bat" };
        List<List<String>> result= groupAnagrams(stars);
        System.out.println(result);

    }

    // this perform better in leet code
    // Time complexity: O(m∗nlogn)
    // Space complexity: O(m∗n)

    static public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> res = new HashMap<>();
        for (String s : strs) {
            char[] charArray = s.toCharArray();
            Arrays.sort(charArray);
            String sortedS = new String(charArray);
            res.putIfAbsent(sortedS, new ArrayList<>());
            res.get(sortedS).add(s);
        }
        return new ArrayList<>(res.values());
    }

}

