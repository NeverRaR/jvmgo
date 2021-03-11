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

    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) throws RuntimeException {
        int x = 32768; // ldc
        Test myObj = new Test(); // new
        Test.staticVar = x; // putstatic
        x = Test.staticVar; // getstatic
        myObj.instanceVar = x; // putfield
        x = myObj.instanceVar; // getfield
        Object obj = myObj;
        if (obj instanceof Test) { // instanceof
            myObj = (Test) obj; // checkcast
            System.out.println(myObj.instanceVar);
        }
    }
    public void hello(){
        int i=100;
    }
}
