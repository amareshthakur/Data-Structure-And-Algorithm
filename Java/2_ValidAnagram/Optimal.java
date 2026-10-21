import java.util.Arrays;

public class Optimal {
    public static void main(String[] args) {
        String sa = "ram";
        String ta = "amr";
        // isAnagram(s, t);
        isAnagram(sa, ta);
        // for(int i=0; i<sa.length();i++){
        //     System.out.println(sa.charAt(i));
        // }
        // System.out.println('z' - 'a');
        
    }
    static public boolean isAnagram(String s, String t) {
        if (s.length() != s.length()) {
            return false;
        }
        int[] charCounts = new int[26];
        for (int i = 0; i < s.length(); i++) {
            charCounts[s.charAt(i) - 'a']++;
            charCounts[t.charAt(i) - 'a']--;
            // System.out.println("charCounts[t.charAt(i) - 'a']++ <-> " + charCounts[t.charAt(i) - 'a']++);
        }
        // System.out.println(Arrays.toString(charCounts));
        for (int count : charCounts) {
            if (count != 0) {
                return false;
            }
        }
        return true;
    }
}
