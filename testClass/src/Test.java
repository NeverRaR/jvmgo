import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

public class Test {

    public static void main(String[] args) {
        String s1 = "abc1";
        String s2 = "abc1";
        System.out.println(s1 == s2); // true
        int x = 1;
        String s3 = "abc" + x;
        System.out.println(s1 == s3); // false
        s3 = s3.intern();
        System.out.println(s1 == s3); // true
    }
    private static void bubbleSort(int[] arr) {
        boolean swapped = true;
        int j = 0; int tmp;
        while (swapped) {
            swapped = false;
            j++; for (int i = 0; i < arr.length - j; i++) {
                if (arr[i] > arr[i + 1]) {
                    tmp = arr[i];
                    arr[i] = arr[i + 1];
                    arr[i + 1] = tmp;
                    swapped = true;
                }
            }
        }

    }
    private static void printArray(int[] arr) {
        for (int i : arr) {
            System.out.printf("%d,",i);
        }
    }
}
