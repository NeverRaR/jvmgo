import java.io.BufferedOutputStream;
import java.io.FileDescriptor;
import java.io.FileOutputStream;
import java.io.PrintStream;


public class Test implements Cloneable  {

    private double pi = 3.14;
    public static void main(String[] args) {
       System.out.println("hello world!");
        new PrintStream(new BufferedOutputStream(new FileOutputStream(FileDescriptor.out), 128), true);
    }
}
