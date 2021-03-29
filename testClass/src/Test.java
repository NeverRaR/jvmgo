import java.io.BufferedOutputStream;
import java.io.FileDescriptor;
import java.io.FileOutputStream;
import java.io.PrintStream;
import java.util.Properties;


public class Test implements Cloneable  {

    public static void main(String[] args) {
        PrintStream  _out = new PrintStream(
                new BufferedOutputStream(new FileOutputStream(FileDescriptor.out), 128), true);
        _out.println("hello world");
    }
}
