import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.Iterator;
// import java.util.Collection;
import java.util.List;

/**
 * AryList
 */
public class AryList {
    public static void main(String[] args) {
               // List or Collection -> Interface
        // List<Integer> list1 = new ArrayList<>();
        

        // ArrayList -> concrete class
        ArrayList<Integer> list = new ArrayList<>();
        // add
        list.add(10);
        list.add(20);
        list.add(30);
        // System.out.println(list);
        list.add(720);
        // remove
        list.remove(1);
        // System.out.println(list);
        // addAll
        List<Integer> list2 = new ArrayList<>();
        list2.add(101);
        list2.add(102);
        list.addAll(list2);
        // System.out.println("List "+list);
        // System.out.println("List2 "+list2);
        list.removeAll(list2);
        // System.out.println("List "+list);
        // System.out.println("List2 "+list2);

        // System.out.println(list.size());
        // System.out.println("List2 "+list2);
        list2.clear();
        // System.out.println("List2 "+list2);
        // System.out.println("list2 size -> " + list2.size());
        
        // I want to traverse list using iterator
        // Iterator<Integer> iterator = list.iterator();
        // while(iterator.hasNext()){
            // System.out.println("Element: "+iterator.next());
        // }

        List<Integer> list3 = new ArrayList<>();
        list3.add(110);
        list3.add(430);
        list3.add(520);
        // System.out.println(list3.get(0));
        // System.out.println(list3.get(1));
        // System.out.println(list3.get(2));
        // System.out.println("Before Set list3 -> " + list3);
        list3.set(0, 123);
        // System.out.println(list3);

        // ArrayList,vector,LinedList,stack -> toArray[]
        Object[] arr= list3.toArray();
        // for (Object elem : arr) {
        //     System.out.println("Element -> "+elem);
        // }
        // contains
        // System.out.println(list3.contains(430));
        // System.out.println(list3.contains(100));

        list.add(60);
        list.add(6);
        System.out.println("printing entire list ->"+list);
        Collections.sort(list);
        System.out.println("printing entire list ->"+list);

    }
}