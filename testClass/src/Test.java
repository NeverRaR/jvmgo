import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

public class Test implements Cloneable  {

    private double pi = 3.14;
    @Override
    public Test clone() {
        try {
            return (Test) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new RuntimeException(e);
        }
    }
    public static void main(String[] args) {
        Test obj1 = new Test();
        Test obj2 = obj1.clone();
        obj1.pi = 3.1415926;
        System.out.println(obj1.pi);
        System.out.println(obj2.pi);
    }
}
