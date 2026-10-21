package HashIng;

import java.util.HashSet;
import java.util.Iterator;

public class HashhSet {
    public static void main(String[] args) {
        //
        HashSet<Integer> set = new HashSet<>();
        set.add(5);
        set.add(1);
        set.add(3);
        set.add(1);
        System.out.println(set);

        // Search contains
        if (set.contains(1)) {
            System.out.println("set contains 1");
        }
        if (!set.contains(100)) {
            System.out.println("set does not contains 100");
        }

        set.remove(1);
        System.out.println(set);
        System.out.println(set.size());

        // /Iterator
        Iterator it = set.iterator();
        while (it.hasNext()) {
            System.out.print(it.next() + "<->");
        }
    }
}

// `HashSet` is an unordered collection that stores only unique elements.
//  It does not guarantee insertion order or sorted order.
//  If elements appear sorted, it is just a coincidence, not a guarantee.