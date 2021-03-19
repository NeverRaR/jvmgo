import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

public class Test {

    public static final boolean FLAG = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public static final int INT = 123456789;
    public static final long LONG = 1123123122345678901L;
    public static final float PI = 3.14f;

    public static int staticVar=100;
    public int instanceVar;

    public static void main(String[] args) {
        long x = fibonacci(30);
        System.out.println(x);
    }
    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        }
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}
