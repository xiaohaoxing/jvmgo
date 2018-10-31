

public class ParseIntTest {
    public static void main(String[] args) {
        foo(args);
    }

    public static void foo(String[] args) {
        try{
            bar(args);
        }catch (NumberFormatException e) {
            System.out.println(e.getMessage());
        }
    }

    public static void bar(String[] args) {
        if (args.length == 0) {
            throw new IndexOutOfBoundsException("No args!");
        }
        int x = Integer.parseInt(args[0]);
        System.out.println(x);
    }
}