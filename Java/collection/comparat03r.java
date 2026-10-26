// package collection;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public class comparat03r {
    public static void main(String[] args) {
        // System.out.println("hello");
        List<Integer> list = new ArrayList<>();
        list.add(2);
        list.add(1);
        list.add(3);
        list.sort(new IntegerComperator());
        // System.out.println(list);
        List<String> words = Arrays.asList("banana", "apple", "date");
        words.sort(new StringLengthComperator());
        System.out.println(words);
    }
}

class StringLengthComperator implements Comparator<String> {
    @Override
    public int compare(String s1, String s2) {
        // [date, apple, banana]
        return s1.length() - s2.length();
    }
}

class IntegerComperator implements Comparator<Integer> {
    @Override
    public int compare(Integer i1, Integer i2) {
        // For descending order
        // return i2 - i1;
        // For ascending order
        return i1 - i2;
    }
}