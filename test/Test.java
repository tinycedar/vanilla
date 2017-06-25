public class Test {

    public static void main(String[] args){
        int a = 1;
        int b = a + 2;
        System.out.println("Hello world !");
        System.out.println(b);

        staticMethod();
    }

    public static void staticMethod() {
        System.out.println("========= static method invoked ! ==========");
    }
}
