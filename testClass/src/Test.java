import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

public class Test {

    public static void main(String[] args) {
        System.out.println(void.class.getName()); // void
        System.out.println(boolean.class.getName()); // boolean
        System.out.println(byte.class.getName()); // byte
        System.out.println(char.class.getName()); // char
        System.out.println(short.class.getName()); // short
        System.out.println(int.class.getName()); // int
        System.out.println(long.class.getName()); // long
        System.out.println(float.class.getName()); // float
        System.out.println(double.class.getName()); // double
        System.out.println(Object.class.getName()); // java.lang.Object
        System.out.println(int[].class.getName()); // [I
        System.out.println(int[][].class.getName()); // [[I
        System.out.println(Object[].class.getName()); // [Ljava.lang.Object;
        System.out.println(Object[][].class.getName()); // [[Ljava.lang.Object;
        System.out.println(Runnable.class.getName()); // java.lang.Runnable
        System.out.println("abc".getClass().getName()); // java.lang.String
        System.out.println(new double[0].getClass().getName()); // [D
        System.out.println(new String[0].getClass().getName()); //[Ljava.lang.S
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
