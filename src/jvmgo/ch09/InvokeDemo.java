public class InvokeDemo implements Runnable {
    public static void main(String[] args) {
        new InvokeDemo().test();
    }

    public void test() {
        InvokeDemo.staticMethod();
        InvokeDemo demo = new InvokeDemo();
        demo.instanceMethod();
        super.equals(null);
        this.run();
        ((Runnable)demo).run();
    }
    public InvokeDemo() {}
    public static void staticMethod() {}
    public void instanceMethod() {}
    @Override
    public void run() {}
}