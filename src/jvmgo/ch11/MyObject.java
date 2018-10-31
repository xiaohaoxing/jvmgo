package test;
/**
 *  a test class file
 */
public class MyObject {
    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) {
        int x = 32798;
        MyObject obj = new MyObject();
        MyObject.staticVar = x;
        x = MyObject.staticVar;
        obj.instanceVar = x;
        x = obj.instanceVar;
        Object o = obj;
        if (o instanceof MyObject) {
            obj = (MyObject)o;
            System.out.println(obj.instanceVar);
        }
    }
}